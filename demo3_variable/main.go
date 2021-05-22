package main

import (
	"fmt"
)

var count int = 0

func main() {
	fmt.Println("Begin")

	// Explicit declaration
	var tmp1 int = 0
	// var tmp2 string = "hello"
	// var tmp3 bool = true

	fmt.Println(tmp1)

	// Implicit declaration
	tmp5 := 0
	fmt.Println(tmp5)
	count++
	fmt.Println(count)
	run()

}

func run() {
	count++
	fmt.Println(count)
}
