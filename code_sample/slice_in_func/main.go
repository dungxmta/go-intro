package main

import "fmt"

// truyền vào array sẽ giúp
// nội dung của biến x không bị thay đổi
func once(x [3]int) {
	for i := range x {
		x[i] *= 2
	}
}

// truyền vào con trỏ ngầm định (slice)
// khiến nội dung của biến x bị thay đổi
func twice(x []int) {
	for i := range x {
		x[i] *= 2
	}
}

func main() {
	data := [3]int{8, 9, 0}

	once(data)
	fmt.Println(data)

	twice(data[0:])
	fmt.Println(data)

	// kết quả:
	// [8 9 0]
	// [16 18 0]
}
