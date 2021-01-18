package main

import (
	"fmt"
	"math/rand"
	"time"
)

func send07(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("[%s] %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			fmt.Printf(`... after put "%v" to channel: %v`, i, msg)
		}
	}()

	return c
}

// go run ./07_receive_only.go

func main() {
	defer fmt.Println("end")

	c := send07("1st")

	// c <- "PANIC"

	for i := 0; i < 5; i++ {
		fmt.Println("\nMain: ", <-c)
	}
}

/*
Main:  [1st] 0
... after put "0" to channel: 1st
Main:  [1st] 1
... after put "1" to channel: 1st
Main:  [1st] 2
... after put "2" to channel: 1st
Main:  [1st] 3
... after put "3" to channel: 1st
Main:  [1st] 4
end
*/
