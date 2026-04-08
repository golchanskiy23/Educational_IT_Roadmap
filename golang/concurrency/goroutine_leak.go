// намеренно плохой код - не запускать, только если не хочешь словить утечку горутин
package main

import(
	"time"
	_ "net/http/pprof"
	"sync"
)

func blocked_receive_channel(){
	ch := make(chan int)
	go func(){
		ch <- 1
	}()

	time.Sleep(100*time.Millisecond)
}

func blocked_writing_channel(){
	ch := make(chan int)
	go func(){
		<-ch
	}()

	time.Sleep(100*time.Millisecond)
}

func infinite_cycle(){
	go func(){
		for {}
	}()
}

func wg_misuse(){
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
    // забыли Done()
	}()

	wg.Wait() // вечная блокировка
}

func work() chan int {
    ch := make(chan int)
	go func(){
		ch <- 42
    	ch <- 13 // (!) утечка
	}()
    return ch
}

func main(){
	// нет читателя - блокировка
	blocked_receive_channel()
	// нет писателя - блокировка
	blocked_writing_channel()
	// в этом случае планировщик скорее будет вытеснять горутину в глобальную очередь
	// но при отсутствии других горутин main завершится - утечка горутины
	infinite_cycle()
	wg_misuse()


	// диагностика - через pprof - потёкшие горутины видны
	go func() {
    	log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	// go tool pprof http://localhost:6060/debug/pprof/goroutine
	// http://localhost:6060/debug/pprof/goroutine?debug=2
	
	// профиль goroutineleak из профайлера в go 1.26
	prof := pprof.Lookup("goroutineleak")

    defer func() {
        time.Sleep(50 * time.Millisecond)
        prof.WriteTo(os.Stdout, 2)
    }()

    <-work()
}