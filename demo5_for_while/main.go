package main

import "fmt"

func main() {
	fnFor()
	fnWhile()
	fnWhileUsingBreak()
}

func fnFor() {
	for index := 0; index < 10; index++ {
		fmt.Printf("Index %d\n", index)
	}
}

func fnWhile() {
	i := 0
	for i < 5 {
		fmt.Printf("Index %d\n", i)
		i++
	}
}

func fnWhileUsingBreak() {
	i := 0
	for true {
		if i > 5 {
			break
		}
		fmt.Printf("Index %d\n", i)
		i++
	}
}
