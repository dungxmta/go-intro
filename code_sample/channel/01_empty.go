package main

import (
	"log"
	"time"
)

func emptyChannel() {
	defer log.Println("i'm done!!!")
	var ch chan int

	// if ch == nil {
	// 	ch = make(chan int)
	// }

	go func() {
		defer log.Println("end goroutines")
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second * 1)
			log.Println("...")
		}
		ch <- 1
	}()

	// <- ch
	time.Sleep(time.Second * 6)
}

func main() {
	emptyChannel()
}

/*
2021/01/17 10:06:20 ...
2021/01/17 10:06:21 ...
2021/01/17 10:06:22 ...
2021/01/17 10:06:24 end goroutines
2021/01/17 10:06:24 i'm done!!!
*/
