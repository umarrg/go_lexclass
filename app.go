package main

import (
	"fmt"
)

func main() {
	fmt.Println("x =: ")
	var x int
	fmt.Scanln(&x)
	fmt.Println("y =: ")
	var y int
	fmt.Scanln(&y)
	fmt.Printf(" the total sum of %v and %v is %v \n", x, y, Add(x, y))

}

func Add(x int, y int) int {
	return x + y
}

func Sub(x int, y int) int {
	return x - y
}

func Mul(x int, y int) int {
	return x * y
}

func Div(x float64, y float64) float64 {
	return x / y
}
