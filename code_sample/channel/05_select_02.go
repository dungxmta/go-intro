package main

import (
	"log"
	"time"
)

// =====================================
// select with multi channel
// =====================================

// go run ./05_select_02.go
func main() {
	defer log.Println("end")
	outOfMoney := make(chan struct{})

	girl01 := make(chan string)
	girl02 := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second * 2)
			log.Println("[1] sending: 99")
			girl01 <- ">> 99  doa hong"
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second * 5)
			log.Println("[2] sending: 100")
			girl02 <- ">> 100 doa hong"
		}
	}()

	time.AfterFunc(time.Second*15, func() {
		outOfMoney <- struct{}{}
	})

	for {
		select {
		case flowers := <-girl01:
			log.Println(flowers, "[1]")
		case flowers := <-girl02:
			log.Println(flowers, "[2]")
		case <-outOfMoney:
			log.Println("i'm out of money !!!")
			return
		default:
			log.Println("...buy flowers...")
			time.Sleep(time.Second * 1)
		}
	}
}

/*
2021/01/17 10:23:20 ...buy flowers...
2021/01/17 10:23:21 ...buy flowers...
2021/01/17 10:23:22 [1] sending: 99
2021/01/17 10:23:22 >> 99  doa hong [1]
2021/01/17 10:23:22 ...buy flowers...
2021/01/17 10:23:23 ...buy flowers...
2021/01/17 10:23:24 [1] sending: 99
2021/01/17 10:23:24 >> 99  doa hong [1]
2021/01/17 10:23:24 ...buy flowers...
2021/01/17 10:23:25 [2] sending: 100
2021/01/17 10:23:25 >> 100 doa hong [2]
2021/01/17 10:23:25 ...buy flowers...
2021/01/17 10:23:26 [1] sending: 99
2021/01/17 10:23:26 >> 99  doa hong [1]
2021/01/17 10:23:26 ...buy flowers...
2021/01/17 10:23:27 ...buy flowers...
2021/01/17 10:23:28 [1] sending: 99
2021/01/17 10:23:28 >> 99  doa hong [1]
2021/01/17 10:23:28 ...buy flowers...
2021/01/17 10:23:29 ...buy flowers...
2021/01/17 10:23:30 [2] sending: 100
2021/01/17 10:23:30 [1] sending: 99
2021/01/17 10:23:30 >> 100 doa hong [2]
2021/01/17 10:23:30 >> 99  doa hong [1]
2021/01/17 10:23:30 ...buy flowers...
2021/01/17 10:23:31 ...buy flowers...
2021/01/17 10:23:32 [1] sending: 99
2021/01/17 10:23:32 >> 99  doa hong [1]
2021/01/17 10:23:32 ...buy flowers...
2021/01/17 10:23:33 ...buy flowers...
2021/01/17 10:23:34 [1] sending: 99
2021/01/17 10:23:34 >> 99  doa hong [1]
2021/01/17 10:23:34 ...buy flowers...
2021/01/17 10:23:35 [2] sending: 100
2021/01/17 10:23:35 >> 100 doa hong [2]
2021/01/17 10:23:35 i'm out of money !!!
2021/01/17 10:23:35 end
*/
