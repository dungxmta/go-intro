package main

import (
	"log"
	"sync"
	"time"
)

// =====================================
// select with 1 channel
// =====================================

var wg04 sync.WaitGroup

func send04(ch chan int) {
	defer func() {
		log.Println("done send")
		wg04.Done()
	}()

	for num := 0; num < 3; num++ {
		log.Println("[+] start sending: ", num)
		ch <- num
		log.Println("[+] after sending: ", num)
	}

	close(ch)
}

func receive04OneCh(ch chan int) {
	defer func() {
		log.Println("done receive")
		wg04.Done()
	}()

	// for num := range ch {
	// 	log.Printf("[-] receive: %v, slepping 2s...\n", num)
	// 	time.Sleep(time.Second * 2)
	// }

	// value_from_channel, is_channel_closed := <-channel
	for {
		select {
		case num, ok := <-ch:
			if !ok {
				return
			}
			log.Printf("[-] receive: %v, slepping 2s...\n", num)
			time.Sleep(time.Second * 2)
		}
	}

	// for {
	// 	num, ok := <-ch
	// 	if !ok {
	// 		break
	// 	}
	// 	log.Printf("[-] receive: %v, slepping 2s...\n", num)
	// 	time.Sleep(time.Second * 2)
	// }
}

/*
2021/01/17 10:28:45 [+] start sending:  0
2021/01/17 10:28:45 [+] after sending:  0
2021/01/17 10:28:45 [+] start sending:  1
2021/01/17 10:28:45 [-] receive: 0, slepping 2s...
2021/01/17 10:28:47 [-] receive: 1, slepping 2s...
2021/01/17 10:28:47 [+] after sending:  1
2021/01/17 10:28:47 [+] start sending:  2
2021/01/17 10:28:49 [-] receive: 2, slepping 2s...
2021/01/17 10:28:49 [+] after sending:  2
2021/01/17 10:28:49 done send
2021/01/17 10:28:51 done receive
2021/01/17 10:28:51 end
*/
func oneChanWithSelect() {
	ch := make(chan int)

	wg04.Add(2)

	go send04(ch)
	go receive04OneCh(ch)

	wg04.Wait()
}

// =====================================
// select with default value
// =====================================

/*
2021/01/17 10:31:36 start
2021/01/17 10:31:38 999 doa hong
2021/01/17 10:31:38 done
2021/01/17 10:31:38 end
*/
func blockingWithSelect() {
	log.Println("start")
	defer log.Println("done")
	girl := make(chan int)

	time.AfterFunc(time.Second*2, func() { // send flowers after 2s
		girl <- 999
	})

	select {
	case flowers := <-girl: // blocking here
		log.Printf("%d doa hong", flowers)
	}
}

/*
2021/01/17 10:34:39 start
2021/01/17 10:34:39 999 doa hong cua tau dau???
2021/01/17 10:34:39 done
2021/01/17 10:34:39 end
*/
func nonblockingWithSelect() {
	log.Println("start")
	defer log.Println("done")
	girl := make(chan int)

	time.AfterFunc(time.Second*2, func() { // no changes to send:v
		girl <- 999
	})

	select {
	case flowers := <-girl: // non-blocking
		log.Println(flowers)
	default:
		log.Println("999 doa hong cua tau dau???")
	}
}

// go run ./04_select_01.go
func main() {
	defer log.Println("end")

	// oneChanWithSelect()
	blockingWithSelect()
	nonblockingWithSelect()
}
