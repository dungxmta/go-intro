package main

import (
	"fmt"
	"time"
)

func demo01() {
	// sử dụng từ khoá go để tạo goroutine
	go fmt.Println("Hello from another goroutine")
	fmt.Println("Hello from main goroutine")

	// chờ 1 giây để có thể chạy được goroutine của hàm fmt.Println trước khi hàm main kết thúc
	time.Sleep(time.Second)
}
