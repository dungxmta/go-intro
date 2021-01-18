package main

import (
	"fmt"
	"math/rand"
	"time"
)

//
// Merge 2 channels into 1 channel
//

func send08(id string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			fmt.Printf("[+][%v] before: %v\n", id, i)
			c <- fmt.Sprintf("[%s] %d", id, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			fmt.Printf("[+][%v] after: %v\n", id, i)
		}
	}()

	return c
}

/*
[+][1st] before: 0
[-][1st] 0
[+][2nd] before: 0
[-][2nd] 0
[+][1st] after: 0
[+][1st] before: 1
[-][1st] 1
[+][2nd] after: 0
[+][2nd] before: 1
[-][2nd] 1
[+][1st] after: 1
[+][1st] before: 2
[-][1st] 2
[+][2nd] after: 1
[+][2nd] before: 2
[-][2nd] 2
[+][1st] after: 2
[+][1st] before: 3
[-][1st] 3
[+][2nd] after: 2
[+][2nd] before: 3
[-][2nd] 3
end
*/
func merge2ChannelWithGoroutines(ch01, ch02 <-chan string) <-chan string {
	merged := make(chan string)

	go func() {
		for {
			merged <- <-ch01
		}
	}()

	go func() {
		for {
			merged <- <-ch02
		}
	}()

	return merged
}

func merge2ChannelWithSelect(ch01, ch02 <-chan string) <-chan string {
	merged := make(chan string)

	go func() {
		for true {
			select {
			case v := <-ch01:
				merged <- v
			case v := <-ch02:
				merged <- v
			case <-time.After(time.Second * 1):
				fmt.Println("...")
			}
		}
	}()

	return merged
}

func test() <-chan int {
	ch03 := make(chan int)
	time.AfterFunc(time.Second*3, func() {
		ch03 <- 1
	})
	return ch03
}

func main() {
	defer fmt.Println("end")

	ch01 := send08("1st")
	ch02 := send08("2nd")

	// merged := merge2ChannelWithGoroutines(ch01, ch02)
	merged := merge2ChannelWithSelect(ch01, ch02)

	for i := 0; i < 8; i++ {
		fmt.Printf("[-]%v\n", <-merged)
	}
}
