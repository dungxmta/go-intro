package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func task01() {
	fmt.Println("task 01")
	// nếu ko dùng Done sẽ raise deadlock: wg ko biết lúc nào kết thúc
	wg.Done()
}

func task02() {
	defer wg.Done()
	fmt.Println("task 02")
}

// synchronized goroutines
//  thứ tự chạy các goroutines do goruntime quản lý
//  logic wg: chờ cho các goroutines hoàn thành xong thì mới kết thúc main
func demo03() {
	fmt.Println("begin")
	defer fmt.Println("end")

	wg.Add(2)
	// neu ko dung wait group -> 2 task co the chua tra ve ket qua nhung main() da ket thuc
	go task01()

	fmt.Println("before call task02")
	go task02()

	wg.Wait()
}

/*
begin
before call task02
task 02
task 01
end
*/
