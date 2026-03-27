package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
	_ "time"

	"golang.org/x/sys/unix"
)

func run_processor(){
	fmt.Println(runtime.GOMAXPROCS(2)) 
	// LRQ size = 256
	wg := &sync.WaitGroup{}

	// основной процессор складывает горутины в LRQ, 
	// а оттуда их ворует второй процессор
	for i := 0; i < 256; i++{
		wg.Add(1)
		go func(id int){
			defer wg.Done()
			for{
				for i:=0; i < 100000000; i++{
					_ = i*i
				}
			}
		}(i)
	}

	wg.Wait()
}

func print_core_nums(){
	fmt.Println(runtime.GOMAXPROCS(2)) 
	// число логических ядер
	fmt.Println(runtime.NumCPU())
	
	// число процессоров
	// впервые - общее число процессоров, затем OK
	fmt.Println(runtime.GOMAXPROCS(2)) // 4
	fmt.Println(runtime.GOMAXPROCS(2)) // 2
	// 1 системный тред для sysmon и 4 треда для каждого процессора
	
	/*
	[bloom@archlinux scheduler]$ go build
	/ / период вывода информации
	[bloom@archlinux scheduler]$ GODEBUG=schedtrace=1000 ./pohui
	SCHED 0ms: gomaxprocs=4 idleprocs=1 threads=5 spinningthreads=1 needspinning=0 idlethreads=0 runqueue=0 [ 1 0 0 0 ] schedticks=[ 0 0 0 0 ]
	4
	4
	2
	*/

	// syscall при выводе на экран, короткий - нет handoff
	fmt.Println("Num Gorutine", runtime.NumGoroutine())

	/*
	[bloom@archlinux scheduler]$ go build
	[bloom@archlinux scheduler]$ GODEBUG=schedtrace=1000 ./pohui
	SCHED 0ms: gomaxprocs=4 idleprocs=2 threads=5 spinningthreads=1 needspinning=0 idlethreads=2 runqueue=0 [ 0 0 0 0 ] schedticks=[ 1 0 0 2 ]
	4
	4
	2
	SCHED 1000ms: gomaxprocs=2 idleprocs=2 threads=5 spinningthreads=0 needspinning=0 idlethreads=3 runqueue=0 [ 0 0 ] schedticks=[ 1 0 ]
	SCHED 2006ms: gomaxprocs=2 idleprocs=2 threads=5 spinningthreads=0 needspinning=0 idlethreads=3 runqueue=0 [ 0 0 ] schedticks=[ 1 0 ]
	SCHED 3011ms: gomaxprocs=2 idleprocs=2 threads=5 spinningthreads=0 needspinning=0 idlethreads=3 runqueue=0 [ 0 0 ] schedticks=[ 1 0 ]
	SCHED 4014ms: gomaxprocs=2 idleprocs=2 threads=5 spinningthreads=0 needspinning=0 idlethreads=3 runqueue=0 [ 0 0 ] schedticks=[ 1 0 ]
	SCHED 5019ms: gomaxprocs=2 idleprocs=2 threads=5 spinningthreads=0 needspinning=0 idlethreads=3 runqueue=0 [ 0 0 ] schedticks=[ 1 0 ]
	*/
}

func thread_handoff(){
	runtime.GOMAXPROCS(3)

	// syscall горутины
	// ожидается 12 тредов : 10 - воркеры, 1 - sysmon, 1 - main
	wg := sync.WaitGroup{}
	for i := 1; i < 10; i++{
		wg.Add((1))
		go func(){
			defer wg.Done()
			var buf [1]byte
			// бесконечное ожидание чтения
			unix.Read(unix.Stdin, buf[:])
		}()
	}

	wg.Wait()
}

func heavyComputation(n int) int{
	start := time.Now()
	time.Sleep(time.Duration(n)*time.Second)
	return int(time.Since(start).Seconds())
}

func mutex_loading(){
	runtime.GOMAXPROCS(4)

	var mu sync.Mutex
	var wg sync.WaitGroup
	counter, numWorkers := 0,4

	for i := 0; i < numWorkers; i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			for{
				mu.Lock()
				// времени на ожидание тратится больше, чем на саму работу
				// из-за этого процессоры простаивают
				work := heavyComputation(100)
				counter += work
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
}

// процесс инициализации работы планировщика
func main(){
	// print_core_nums()
	// run_processor()
	// thread_handoff()
	// mutex_loading()
}