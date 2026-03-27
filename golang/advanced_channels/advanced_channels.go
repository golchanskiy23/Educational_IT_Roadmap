package main

import (
	"fmt"
	"time"
	"sync"
)

func sync_channel(){
	// синхронный канал
    ch := make(chan bool)
    go func() {
        ch <- true
    }()
	// допустим сначала чтение, тогда main-горутина блокируется
	// в recvq и ждёт записи:
	// удаляет из очереди -> записывает данные в стек -> снимает блокировку
    <-ch
}

func sync_buffered_channel(){
	ch := make(chan bool, 1)
	// проверка, можем ли записать данные в канал, после следующей попытки
	// записи блокируем записывающую горутину, пока не прочитаем
    ch <- true
    go func() {
        <-ch
    }()
	// читаем из буфера -> кладём значение из стека горутины из sendq в буфер -> разблокируем в main ->
	// удаляем main из sendq (ибо main ждёт, когда в буфере освободится место, иначе вечная блокировка)
    ch <- true
}

func deadlock_example(){
	c := make(chan string)
	c <- "deadlock"
}

func squares(c chan int){
	/*defer close(c)
	for i := 0; i <= 9; i++{
		c <- i*i
	}*/
	num := <-c
	c <- num*num
}

// for полезен, когда не знаем сколько данных нужно считать
func main_block_unblock_for(){
	c := make(chan int)

	go squares(c)

	for{
		val, ok := <- c
		if !ok{
			fmt.Println(val, ok, "channel closed")
			break
		}
		fmt.Println(val, ok)
	}

	// автоматическая остановка цикла при закрытом канале
	for val := range c {
        fmt.Println(val)
    }
}

func read_squares(c chan int) {
    for i := 0; i < 3; i++ {
        num := <-c
        fmt.Println(num * num)
    }
}

func buffered_channel(){
    c := make(chan int, 3)

	wg := &sync.WaitGroup{}
    wg.Add(1)

    go func() {
        defer wg.Done()
        read_squares(c)
    }()
    
    c <- 1
    c <- 2
    c <- 3
	c <- 4

	// 1,4,9	16 (не прочитается, но будет в буфере)

	close(c)
	wg.Wait()
}

func cubes(c chan int){
	num := <-c
	c <- num*num*num
}

func two_chans_work(){
	squareChan := make(chan int)
    cubeChan := make(chan int)
	
    go squares(squareChan)
    go cubes(cubeChan)

    testNum := 3

    squareChan <- testNum
    cubeChan <- testNum

    squareVal, cubeVal := <-squareChan, <-cubeChan
    sum := squareVal + cubeVal

    fmt.Println("[main] sum of square and cube of", testNum, " is", sum)
}

// улучшение безопасноти типов в канале
func unidirectional_channel(){
	c := make(chan string)
	go func(channel <-chan string){
		fmt.Printf("Value is %s", <-c)
	}(c)
	c <- "testing"
}

func greeter(c chan chan string){
	cv := make(chan string)
	c <- cv
}

func greet(c chan string){
	c <- "tmp"
}

func chan_chan(){
	cc := make(chan chan string)

	go greeter(cc)
	c := <- cc

	go greet(c)
	fmt.Print(<-c)
}

var start time.Time

func init() {
    start = time.Now()
}

func timer(){
	chan1 := make(chan string)
    chan2 := make(chan string)

    go func(c chan string){
		time.Sleep(3 * time.Second)
    	c <- "Hello from service 1"
	}(chan1)

    go func(c chan string){
		time.Sleep(5 * time.Second)
    	c <- "Hello from service 2"
	}(chan2)

    select {
    case res := <-chan1:
        fmt.Println("Response from service 1", res, time.Since(start))
    case res := <-chan2:
        fmt.Println("Response from service 2", res, time.Since(start))
    case <-time.After(2 * time.Second):
        fmt.Println("No response received", time.Since(start))
    }

    fmt.Println("main() stopped", time.Since(start))
}

func service(wg *sync.WaitGroup, instance int){
	time.Sleep(2 * time.Second)
    fmt.Println("Service called on instance", instance)
    wg.Done() // decrement counter
}

func wg(){
	var wg sync.WaitGroup // create waitgroup (empty struct)

    for i := 1; i <= 3; i++ {
        wg.Add(1) // increment counter
        go service(&wg, i)
    }

    wg.Wait() // blocks here
}

func sqrWorker(wg *sync.WaitGroup, tasks <-chan int, results chan <-int, id int){
	for num := range tasks {
        time.Sleep(time.Millisecond) // simulating blocking task
        fmt.Printf("[worker %v] Sending result by worker %v\n", id, id)
        results <- num * num
    }
	wg.Done()
}


// создаём определённый набор горутин для выполнения задачи
func worker_pool(){
	tasks := make(chan int, 10)
    results := make(chan int, 10)

	var wg sync.WaitGroup
    // launching 3 worker goroutines
    for i := 0; i < 3; i++ {
		wg.Add(1)
        go sqrWorker(&wg, tasks, results, i)
    }

    // passing 5 tasks
    for i := 0; i < 5; i++ {
        tasks <- i * 2 // non-blocking as buffer capacity is 10
    }

    // closing tasks
    close(tasks)

	// wait until all workers done their job
    wg.Wait()

    // receving results from all workers
    for i := 0; i < 5; i++ {
        result := <-results // blocking because buffer is empty
        fmt.Println("[main] Result", i, ":", result)
    }
}

func fib(n int) <-chan int{
	c := make(chan int, n)
	go func(){
		for i,j := 0,1; i < n; i,j = i+j,i {
			c <- i
		}
		close(c)
	}()

	return c
}

func getInputChannel() <- chan int{
	input := make(chan int, 100)
	numbers := []int{0,1,2,3,4,5,6,7,8,9}

	go func(){
		for _, val := range numbers{
			input <- val
		}
		close(input)
	}()

	return input
}

func getSquareChan(input <-chan int) <-chan int{
	output := make(chan int, 100)

	go func(){
		for num := range input{
			output <- num*num
		}
		close(output)
	}()

	return output
}

func merge(channels ...<-chan int) <-chan int{
	var wg sync.WaitGroup

	merged := make(chan int, 100)
	wg.Add(len(channels))

	for _, channel := range channels{
		go func(sc <-chan int){
			defer wg.Done()
			for sq := range sc{
				merged <- sq
			}
		}(channel)
	}

	go func(){
		wg.Wait()
		close(merged)
	}()

	return merged
}

func fan_in_out(){
	chanInputNums := getInputChannel()
	
	// fan-out
	chanOptSqr1 := getSquareChan(chanInputNums)
    chanOptSqr2 := getSquareChan(chanInputNums)

	// fan-in
	chanMergedSqr := merge(chanOptSqr1, chanOptSqr2)

	// sum all squares
	var sqrSum int = 0
	for num := range chanMergedSqr {
        sqrSum += num
    }

	fmt.Printf("Sum is: %d", sqrSum)
}

func generator(){
	for fb := range fib(10){
		fmt.Println("Current fibonacci number is ", fb)
	}
}

func main() {
	// sync_channel()
	// sync_buffered_channel()
	// deadlock_example()
	// main_block_unblock_for()
	// buffered_channel()
	// two_chans_work()	
	// unidirectional_channel()
	// chan_chan()
	// timer()
	// wg()
	// worker_pool()
	
	// generator()
	// fan_in_out()
}