package main

import (
	"fmt"
	"time"
)

func print(s string) {
	defer fmt.Println(s, "done")
	for i := 0; i < 3; i++ {
		fmt.Printf("[%10v] %d\n", s, i)
	}
}

func demo02() {

	print("direct")

	go print("goroutine")

	go func() {
		fmt.Println("anonymous function ")
	}()

	go func(s string) {
		fmt.Println(s)
	}("anonymous function with param")

	time.Sleep(time.Second * 2)
}

/*
[    direct] 0
[    direct] 1
[    direct] 2
direct done
anonymous function
anonymous function with param
[ goroutine] 0
[ goroutine] 1
[ goroutine] 2
goroutine done
*/
