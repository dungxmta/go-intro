package main

import "fmt"

const primitiveTypes = `int, uint, int8, uint8, ...
                        bool, string
                        float32, float64
                        complex64, complex128`

func main() {
	fmt.Printf("Value: %6v, Type: %T\n", "Baiju", "Baiju")
	fmt.Printf("Value: %6v, Type: %T\n", 7, 7)
	fmt.Printf("Value: %6v, Type: %T\n", uint(7), uint(7))
	fmt.Printf("Value: %6v, Type: %T\n", int8(7), int8(7))
	fmt.Printf("Value: %6v, Type: %T\n", true, true)
	fmt.Printf("Value: %6v, Type: %T\n", 7.0, 7.0)
	fmt.Printf("Value: %v, Type: %T\n", (1 + 6i), (1 + 6i))
}

/*
Value:  Baiju, Type: string
Value:      7, Type: int
Value:      7, Type: uint
Value:      7, Type: int8
Value:   true, Type: bool
Value:      7, Type: float64
Value: (1+6i), Type: complex128
*/
