package main

import (
	"log"
	"sync"
	"time"
)

var wg02 sync.WaitGroup

func send02(ch chan int) {
	defer func() {
		log.Println("done send")
		wg02.Done()
	}()

	for num := 0; num < 3; num++ {
		log.Println("[+] start sending: ", num)
		ch <- num
		log.Println("[+] after sending: ", num)
	}

	close(ch)
}

func receive02(ch chan int) {
	defer func() {
		log.Println("done receive")
		wg02.Done()
	}()

	for num := range ch {
		log.Printf("[-] receive: %v, slepping 2s...\n", num)
		time.Sleep(time.Second * 2)
	}
}

func main() {
	defer log.Println("end")

	// UnBuffered Channels
	ch := make(chan int)

	wg02.Add(2)

	go send02(ch)
	go receive02(ch)

	wg02.Wait()
}

// go run ./02_blocking.go

/*
2021/01/17 10:53:44 [+] start sending:  0
2021/01/17 10:53:44 [+] after sending:  0
2021/01/17 10:53:44 [+] start sending:  1
2021/01/17 10:53:44 [-] receive: 0, slepping 2s...
2021/01/17 10:53:46 [-] receive: 1, slepping 2s...
2021/01/17 10:53:46 [+] after sending:  1
2021/01/17 10:53:46 [+] start sending:  2
2021/01/17 10:53:48 [+] after sending:  2
2021/01/17 10:53:48 done send
2021/01/17 10:53:48 [-] receive: 2, slepping 2s...
2021/01/17 10:53:50 done receive
2021/01/17 10:53:50 end
*/

/*
comment: close(ch)

2021/01/17 10:53:44 [+] start sending:  0
2021/01/17 10:53:44 [+] after sending:  0
2021/01/17 10:53:44 [+] start sending:  1
2021/01/17 10:53:44 [-] receive: 0, slepping 2s...
2021/01/17 10:53:46 [-] receive: 1, slepping 2s...
2021/01/17 10:53:46 [+] after sending:  1
2021/01/17 10:53:46 [+] start sending:  2
2021/01/17 10:53:48 [+] after sending:  2
2021/01/17 10:53:48 done send
2021/01/17 10:53:48 [-] receive: 2, slepping 2s...
fatal error: all goroutines are asleep - deadlock!
*/
