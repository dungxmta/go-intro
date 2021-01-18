package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg10 sync.WaitGroup

func worker(id string, stopCh chan struct{}) {
	defer wg10.Done()
	defer fmt.Printf("[%v] i'm done !!!\n", id)

	for {
		select {
		case <-stopCh:
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
	stopCh := make(chan struct{})

	time.AfterFunc(time.Second*10, func() {
		close(stopCh)
	})

	wg10.Add(5)

	for i := 0; i < 5; i++ {
		go func(id string) {
			worker(id, stopCh)
		}(fmt.Sprintf("%v", i))
	}

	wg10.Wait()
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
