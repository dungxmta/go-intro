package main

import (
	"log"
	"sync"
	"time"
)

var wg03 sync.WaitGroup

func send03(ch chan int) {
	defer func() {
		log.Println("done send")
		wg03.Done()
	}()

	for num := 0; num < 5; num++ {
		log.Println("[+] start sending: ", num)
		ch <- num
		log.Println("[+] after sending: ", num)
	}

	close(ch)
}

func receive03(ch chan int) {
	defer func() {
		log.Println("done receive")
		wg03.Done()
	}()

	for num := range ch {
		log.Printf("[-] receive: %v, slepping 2s...\n", num)
		time.Sleep(time.Second * 2)
	}
}

func main() {
	defer log.Println("end")

	// Buffered Channels
	bufferLen := 2
	ch := make(chan int, bufferLen)

	wg03.Add(2)

	go send03(ch)
	go receive03(ch)

	wg03.Wait()
}

// go run ./03_unblocking.go

/*
2021/01/17 10:58:46 [+] start sending:  0
2021/01/17 10:58:46 [+] after sending:  0
2021/01/17 10:58:46 [+] start sending:  1
2021/01/17 10:58:46 [+] after sending:  1
2021/01/17 10:58:46 [+] start sending:  2
2021/01/17 10:58:46 [+] after sending:  2
2021/01/17 10:58:46 [+] start sending:  3
2021/01/17 10:58:46 [-] receive: 0, slepping 2s...
2021/01/17 10:58:48 [-] receive: 1, slepping 2s...
2021/01/17 10:58:48 [+] after sending:  3
2021/01/17 10:58:48 [+] start sending:  4
2021/01/17 10:58:50 [-] receive: 2, slepping 2s...
2021/01/17 10:58:50 [+] after sending:  4
2021/01/17 10:58:50 done send
2021/01/17 10:58:52 [-] receive: 3, slepping 2s...
2021/01/17 10:58:54 [-] receive: 4, slepping 2s...
2021/01/17 10:58:56 done receive
2021/01/17 10:58:56 end
*/
