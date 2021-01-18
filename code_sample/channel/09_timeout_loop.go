package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

/**
Each msg timeout after 1s - Main will end after 5s
*/

func send09(id string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			log.Printf("[+][%v] before: %v\n", id, i)
			c <- fmt.Sprintf("[%s] %d", id, i)
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			log.Printf("[+][%v] after: %v\n", id, i)
		}
	}()

	return c
}

/*
2021/01/17 10:38:08 [+][1st] before: 0
2021/01/17 10:38:08 [-][1st] 0
2021/01/17 10:38:08 [+][1st] after: 0
2021/01/17 10:38:08 [+][1st] before: 1
2021/01/17 10:38:08 [-][1st] 1
2021/01/17 10:38:08 [+][1st] after: 1
2021/01/17 10:38:08 [+][1st] before: 2
2021/01/17 10:38:08 [-][1st] 2
2021/01/17 10:38:09 [+][1st] after: 2
2021/01/17 10:38:09 [+][1st] before: 3
2021/01/17 10:38:09 [-][1st] 3
2021/01/17 10:38:09 [+][1st] after: 3
2021/01/17 10:38:09 [+][1st] before: 4
2021/01/17 10:38:09 [-][1st] 4
2021/01/17 10:38:10 > 1s...
2021/01/17 10:38:10 end
*/
func timeoutEachMsg(c <-chan string) {
	quitLoop := false

	for {
		select {
		case msg := <-c:
			log.Printf("[-]%v\n", msg)
		case <-time.After(1 * time.Second):
			log.Println("> 1s...")
			quitLoop = true
		}

		if quitLoop {
			break
		}
	}
}

/*
2021/01/17 10:38:29 [+][1st] before: 0
2021/01/17 10:38:29 [-][1st] 0
2021/01/17 10:38:29 [+][1st] after: 0
2021/01/17 10:38:29 [+][1st] before: 1
2021/01/17 10:38:29 [-][1st] 1
2021/01/17 10:38:30 [+][1st] after: 1
2021/01/17 10:38:30 [+][1st] before: 2
2021/01/17 10:38:30 [-][1st] 2
2021/01/17 10:38:30 [+][1st] after: 2
2021/01/17 10:38:30 [+][1st] before: 3
2021/01/17 10:38:30 [-][1st] 3
2021/01/17 10:38:30 [+][1st] after: 3
2021/01/17 10:38:30 [+][1st] before: 4
2021/01/17 10:38:30 [-][1st] 4
2021/01/17 10:38:31 [+][1st] after: 4
2021/01/17 10:38:31 [+][1st] before: 5
2021/01/17 10:38:31 [-][1st] 5
2021/01/17 10:38:32 [+][1st] after: 5
2021/01/17 10:38:32 [+][1st] before: 6
2021/01/17 10:38:32 [-][1st] 6
2021/01/17 10:38:32 [+][1st] after: 6
2021/01/17 10:38:32 [+][1st] before: 7
2021/01/17 10:38:32 [-][1st] 7
2021/01/17 10:38:33 [+][1st] after: 7
2021/01/17 10:38:33 [+][1st] before: 8
2021/01/17 10:38:33 [-][1st] 8
2021/01/17 10:38:34 Timeout 5s...
2021/01/17 10:38:34 end
*/
func timeoutMainAfter(c <-chan string) {
	breakLoop := false
	timeout := time.After(5 * time.Second)

	for {
		select {
		case msg := <-c:
			log.Printf("[-]%v\n", msg)
		case <-timeout:
			log.Println("Timeout 5s...")
			breakLoop = true
		}

		if breakLoop {
			break
		}
	}
}

func main() {
	defer log.Println("end")
	c := send09("1st")

	// timeoutEachMsg(c)
	timeoutMainAfter(c)
}
