package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	ch <- 1 // send 1 to channel
	fmt.Println("Step 1")
	fmt.Println(<-ch)

	ch <- 2
	fmt.Println("Step 2")
	fmt.Println(<-ch)

	time.Sleep(1 * time.Second)
}
