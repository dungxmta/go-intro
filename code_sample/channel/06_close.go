package main

import (
	"log"
	"time"
)

/*
2021/01/17 10:40:47 ...
2021/01/17 10:40:48 ...
2021/01/17 10:40:49 value={} | closed=false
2021/01/17 10:40:49 end
*/
func closeCh01() {
	ch := make(chan struct{})

	time.AfterFunc(time.Second*2, func() {
		close(ch)
	})

	for {
		select {
		case value, closed := <-ch:
			log.Printf("value=%v | closed=%v", value, closed)
			return
		default:
			log.Println("...")
			time.Sleep(time.Second * 1)
		}
	}
}

/*
2021/01/17 10:43:28 i'm done
2021/01/17 10:43:28 end
*/
func closeCh02() {
	ch := make(chan int)
	close(ch)

	// ch <- 1 // PANIC !!!

	for {
		select {
		case <-ch:
			log.Println("i'm done")
			return
		default:
			log.Println("...")
			time.Sleep(time.Second * 1)
		}
	}
}

// go run ./06_close.go
func main() {
	defer log.Println("end")

	// closeCh01()
	closeCh02()
}
