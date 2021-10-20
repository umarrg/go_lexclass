package main

import (
	"fmt"
)

func main() {
	AddInput()
	SubInput()
	MulInput()
	DivInput()
}

func AddInput() {
	fmt.Println("Enter number ")
	var x int
	fmt.Scanln(&x)
	fmt.Println("Enter number")
	var y int
	fmt.Scanln(&y)
	fmt.Printf(" the total sum of %v and %v is %v \n", x, y, Add(x, y))

}

func SubInput() {
	fmt.Println("Enter number")
	var x int
	fmt.Scanln(&x)
	fmt.Println("Enter number ")
	var y int
	fmt.Scanln(&y)
	fmt.Printf(" the difference of %v and %v is %v \n", x, y, Sub(x, y))

}

func MulInput() {
	fmt.Println("Enter number ")
	var x int
	fmt.Scanln(&x)
	fmt.Println("Enter number ")
	var y int
	fmt.Scanln(&y)
	fmt.Printf(" the product of %v and %v is %v \n", x, y, Mul(x, y))

}

func DivInput() {
	fmt.Println("Enter number ")
	var x float64
	fmt.Scanln(&x)
	fmt.Println("Enter number")
	var y float64
	fmt.Scanln(&y)
	fmt.Printf(" the division of  %v and %v is %v \n", x, y, Div(x, y))
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
