package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg11 sync.WaitGroup

func worker11(ctx context.Context, id string) {
	defer wg11.Done()
	defer fmt.Printf("[%v] i'm done !!!\n", id)

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("[%v] stopped !!!\n", id)
			return
		default:
		}

		// sleeping ...
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		fmt.Printf("[%v] i'm working...\n", id)
	}
}

func main() {
	parentCtx := context.TODO()
	ctx, cancelFn := context.WithCancel(parentCtx)

	time.AfterFunc(time.Second*10, func() {
		cancelFn()
	})

	wg11.Add(5)

	for i := 0; i < 5; i++ {
		go func(ctx context.Context, id string) {
			worker11(ctx, id)
		}(ctx, fmt.Sprintf("%v", i))
	}

	wg11.Wait()
}

// go run ./10_stop_signal.go

/*
[3] i'm working...
[3] i'm working...
[0] i'm working...
[4] i'm working...
[2] i'm working...
[1] i'm working...
[1] i'm working...
[2] i'm working...
[2] i'm working...
[2] i'm working...
[3] i'm working...
[3] stopped !!!
[3] i'm done !!!
[0] i'm working...
[0] stopped !!!
[0] i'm done !!!
[4] i'm working...
[4] stopped !!!
[4] i'm done !!!
[1] i'm working...
[1] stopped !!!
[1] i'm done !!!
[2] i'm working...
[2] stopped !!!
[2] i'm done !!!
*/
