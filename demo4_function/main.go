package main

import (
	"fmt"
)

func main() {
	display()
	display2("Good morning")
	const a, b = 2, 3
	fmt.Println("sum = ", sum(a, b), " baht")

	x, y := swap(a, b)
	fmt.Printf("%d,%d => %d,%d\n", a, b, x, y)

	x, y = swap2(10, 20)
	fmt.Printf("%d,%d => %d,%d\n", 10, 20, x, y)
}

func display() {
	fmt.Println("M9nnod")
}

func display2(msg string) {
	fmt.Println(msg)
}

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func swap(num1 int, num2 int) (int, int) {
	return num2, num1
}

func swap2(num1, num2 int) (x, y int) {
	y = num1
	x = num2
	return
}
