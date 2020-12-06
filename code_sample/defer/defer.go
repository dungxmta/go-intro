package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		panic("cannot create file")
	}

	// chắc chắn file sẽ được close dù hàm có bị panic hay return
	defer f.Close()
	fmt.Fprintf(f, "hello")
}
