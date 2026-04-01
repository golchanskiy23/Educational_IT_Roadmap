/*
# Основной запуск
go run gc.go

# Увидеть решения компилятора по escape analysis
go run -gcflags="-m" gc_demo.go 2>&1 | grep -E "escape|does not"

# Подробный лог каждого GC-цикла в stderr
GODEBUG=gctrace=1 go run gc_demo.go 2>&1

# Трейс для визуализации в браузере
go run gc.go 2>&1 | go tool trace
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func staysOnStack() int {
	x := 42
	return x
}

func escapesToHeap() *int {
	x := 42
	return &x
}

func interfaceBoxing() interface{} {
	return 42
}

func allocPressure(n int) {
	for i := 0; i < n; i++ {
		_ = make([]byte, 1024) // 1 KB, сразу становится мусором
	}
}

var bufPool = sync.Pool{
	New: func() any {
		return make([]byte, 1024)
	},
}

func allocWithPool(n int) {
	for i := 0; i < n; i++ {
		b := bufPool.Get().([]byte)
		_ = b
		bufPool.Put(b) // возвращаем — GC не видит как мусор
	}
}

func readStats(label string) {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)

	fmt.Printf("\n── %s ──\n", label)
	fmt.Printf("  Alloc (живых байт):     %6d KB\n", ms.Alloc/1024)
	fmt.Printf("  TotalAlloc (всего):     %6d KB\n", ms.TotalAlloc/1024)
	fmt.Printf("  HeapObjects (объектов): %6d\n", ms.HeapObjects)
	fmt.Printf("  NumGC (циклов GC):      %6d\n", ms.NumGC)

	if ms.NumGC > 0 {
		lastPause := ms.PauseNs[(ms.NumGC+255)%256]
		fmt.Printf("  Последняя STW-пауза:    %6d µs\n", lastPause/1000)
	}

	// GCCPUFraction — доля CPU, потраченная на GC (0.0–1.0)
	fmt.Printf("  GCCPUFraction:          %.4f\n", ms.GCCPUFraction)
}

type Tracked struct {
	name string
}

func newTracked(name string) *Tracked {
	t := &Tracked{name: name}
	// Финализатор = хук на момент, когда GC решил собрать объект
	runtime.SetFinalizer(t, func(t *Tracked) {
		fmt.Printf("  [GC sweep] объект '%s' собран\n", t.name)
	})
	return t
}

func demonstrateCollection() {
	fmt.Println("\n── Три объекта: два станут мусором, один останется жить ──")
	a := newTracked("A (останется)")
	_ = newTracked("B (мусор)")       // нет ссылки → недостижим
	_ = newTracked("C (мусор)")       // нет ссылки → недостижим
	_ = a                             // a достижим из стека → не мусор

	runtime.GC() // принудительно запускаем цикл
	// Финализаторы вызываются асинхронно, даём время
	time.Sleep(10 * time.Millisecond)
	runtime.GC()
	time.Sleep(10 * time.Millisecond)

	fmt.Printf("  a='%s' всё ещё жив (есть ссылка)\n", a.name)
}

var sink *Tracked 

func demonstrateWriteBarrier() {
	fmt.Println("\n── Write barrier: объект, созданный во время GC ──")

	go func() {
		runtime.GC()
	}()

	// Новые объекты во время _GCmark фазы создаются чёрными —
	// они не попадут в sweep этого цикла.
	sink = newTracked("D (создан во время GC)")
	time.Sleep(20 * time.Millisecond)

	fmt.Printf("  sink='%s' жив после GC (создан чёрным)\n", sink.name)
	sink = nil // теперь отпускаем — следующий GC соберёт
}

func escape_analysis(){
	fmt.Println("\n[1] Escape analysis")
	fmt.Println("  staysOnStack():", staysOnStack(), "← значение, не указатель")
	p := escapesToHeap()
	fmt.Println("  escapesToHeap():", *p, "← объект в куче, живёт после return")
	fmt.Println("  interfaceBoxing():", interfaceBoxing(), "← boxing = аллокация")
}

func baseline_stats(){
	runtime.GC()
	readStats("Baseline (после старта)")
}

func allocation(){
	fmt.Println("\n[2] Аллокации без Pool (10 000 объектов по 1 KB)")
	allocPressure(10_000)
	readStats("После 10k аллокаций без Pool")
}

func allocation_pool(){
	fmt.Println("\n[3] Аллокации с sync.Pool (10 000 объектов)")
	gcBefore := getNumGC()
	allocWithPool(10_000)
	gcAfter := getNumGC()
	readStats("После 10k аллокаций с Pool")
	fmt.Printf("  GC-циклов во время Pool-теста: %d\n", gcAfter-gcBefore)
}

func tricolor(){
	fmt.Println("\n[4] Трёхцветный алгоритм — финализаторы как наблюдатели")
	demonstrateCollection()
}

func writeBarrier(){
	fmt.Println("\n[5] Write barrier — объект созданный во время GC")
	demonstrateWriteBarrier()
}

func stats(){
	runtime.GC()
	readStats("Итог")

	fmt.Println("  GODEBUG=gctrace=1 go run gc.go   	  — лог каждого GC-цикла")
	fmt.Println("  go tool pprof -alloc_space             — профиль аллокаций")
	fmt.Println("  go tool trace                          — временна́я шкала GC-фаз")
	fmt.Println("  GOEXPERIMENT=greenteagc go build       — Green Tea GC (Go 1.25+)")
}

func main() {
	escape_analysis()
	baseline_stats()
	allocation()
	allocation_pool()
	tricolor()
	writeBarrier()
	stats()
}

func getNumGC() uint32 {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	return ms.NumGC
}