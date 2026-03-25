package main

import (
    "fmt"
    "runtime"
    "time"
    "strconv"
)

func printNumbers(prefix string) {
    for i := 0; i < 5; i++ {
        fmt.Printf("%s: %d\n", prefix, i)
        time.Sleep(1 * time.Millisecond) // Имитация длительной работы
    }
}

func someLongComputation() int{
    time.Sleep(100 * time.Millisecond)
    return 22
}

func generator(nums ...int) <-chan int{
    out := make(chan int)
    go func(){
        for _, num := range nums{
            out <- num
        }
        close(out)
    }()
    return out
}

func square(in <-chan int) <- chan int{
    out := make(chan int)
    go func(){
        for n := range in{
            out <- n*n
        }
        close(out)
    }()
    return out
}

func toString(in <-chan int) <-chan string{
    out := make(chan string)
    go func(){
        for val := range in{
            out <- strconv.Itoa(val)
        }
        close(out)
    }()
    return out
}

func operation(in chan string){
    time.Sleep(5*time.Second)
}

func timeout_control(){
    ch := make(chan string)
    go operation(ch)

    select{
    case res := <-ch:
        fmt.Println(res)
    case <-time.After(1*time.Second):
        fmt.Println("Timeout")
    }
}

func worker(id int, ch chan string) {
    for {
        time.Sleep(time.Second)
        ch <- fmt.Sprintf("worker %d: завершил задачу", id)
    }
}

func select_orchestrator(){
    ch1 := make(chan string)
    ch2 := make(chan string)

    go worker(1, ch1)
    go worker(2, ch2)

    for i := 0; i < 5; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println(msg1)
        case msg2 := <-ch2:
            fmt.Println(msg2)
        }
    }
}

// Планировщик Go будет мультиплексировать эти Goroutines на одном потоке ОС, 
// попеременно выделяя им время CPU для выполнения.
func main() {
    runtime.GOMAXPROCS(1) // Ограничение использования одним процессорным ядром
    go printNumbers("Goroutine1")
    go printNumbers("Goroutine2")
    time.Sleep(100 * time.Millisecond) // Дать время для завершения Goroutines

    ch := make(chan int)
    go func() {
       result := someLongComputation()
       ch <- result
    }()
    // проверка закрытия канала
    for {
       result, ok := <-ch
       if !ok {
           break
       }
       fmt.Println("Received:", result)
    }

    gen := generator(2,3,4)
    sq := square(gen)
    str := toString(sq)

    for s := range str{
        fmt.Println(s)
    }

    // таймаут для одного канала
    timeout_control()
    // select для нескольких каналов
    select_orchestrator()
}