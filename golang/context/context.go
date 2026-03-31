package main

import(
	"context"
	"time"
	"fmt"
	"os"
	"math/rand"
)

type Value int
func Process(ctx context.Context) (Value, error){
	time.Sleep(6*time.Second)
	return 22, nil
}

func StreamWithDone(ctx context.Context, ch chan Value) error{
	for{
		value, err := Process(ctx)
		if err != nil{
			return err
		}
		select{
		case <- ctx.Done():
			return ctx.Err()
		case ch <- value:
			fmt.Println(<-ch)
			return nil
		}
	}

}

func example_context(){
	output := make(chan Value)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if err := StreamWithDone(ctx , output); err != nil{
		fmt.Println("something error")
		os.Exit(1)
	}
	cancel()
}

func sleepRandom(fromFunction string, ch chan int) {
    defer func() { fmt.Println(fromFunction, "sleepRandom complete") }()

    seed := time.Now().UnixNano()
    r := rand.New(rand.NewSource(seed))
    randomNumber := r.Intn(100)
    sleeptime := randomNumber + 100

    fmt.Println(fromFunction, "Starting sleep for", sleeptime, "ms")
    time.Sleep(time.Duration(sleeptime) * time.Millisecond)
    fmt.Println(fromFunction, "Waking up, slept for ", sleeptime, "ms")

    if ch != nil {
        ch <- sleeptime
    }
}

func sleepRandomContext(ctx context.Context, ch chan bool) {
    defer func() {
        fmt.Println("sleepRandomContext complete")
        ch <- true
    }()

    sleeptimeChan := make(chan int)

    go sleepRandom("sleepRandomContext", sleeptimeChan)

    select {
        case <-ctx.Done():
            // Если контекст отменен, выбирается этот случай
            // Это случается, если заканчивается таймаут doWorkContext или
            // doWorkContext или main вызывает cancelFunction
            // Высвобождаем ресурсы, которые больше не нужны из-за прерывания работы
            // Посылаем сигнал всем горутинам, которые должны завершиться (используя каналы)
            // Обычно вы посылаете что-нибудь в канал,
            // ждете выхода из горутины, затем возвращаетесь
            // Или используете группы ожидания вместо каналов для синхронизации
            fmt.Println("sleepRandomContext: Time to return")

        case sleeptime := <-sleeptimeChan:
            // Этот вариант выбирается, когда работа завершается до отмены контекста
            fmt.Println("Slept for ", sleeptime, "ms")
    }
}

func doWorkContext(ctx context.Context) {
    ctxWithTimeout, cancelFunction := context.WithTimeout(ctx, time.Duration(150)*time.Millisecond)

    defer func() {
        fmt.Println("doWorkContext complete")
        cancelFunction()
    }()

    ch := make(chan bool)
    go sleepRandomContext(ctxWithTimeout, ch)

    select {
        case <-ctx.Done():
            // Этот случай выбирается, когда переданный в качестве аргумента контекст уведомляет о завершении работы
            // В данном примере это произойдёт, когда в main будет вызвана cancelFunction
            fmt.Println("doWorkContext: Time to return")

        case <-ch:
            // Этот вариант выбирается, когда работа завершается до отмены контекста
            fmt.Println("sleepRandomContext returned")
    }
}

func example_context_2(){
	ctx := context.Background()
	ctxWithCancel, cancelFunction := context.WithCancel(ctx)

	defer func(){
		fmt.Println("Main Defer: canceling context")
    	cancelFunction()
	}()

	go func(){
		sleepRandom("Main", nil)
    	cancelFunction()
    	fmt.Println("Main Sleep complete. canceling context")
	}()

	doWorkContext(ctxWithCancel)
}

func main(){
	// example_context()	
	example_context_2()
}