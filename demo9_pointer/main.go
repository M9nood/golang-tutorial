package main

import "fmt"

func main() {
	msg := "some message"
	var msgPointer *string = &msg
	fmt.Println(msg)
	fmt.Println(*msgPointer)

	changeMessage(&msg, "New massage")
	fmt.Println(msg)

	changeMessage(msgPointer, "New message from pointer")
	fmt.Println(msg)

	changeMessage(msgPointer, "New message from pointer 2")
	fmt.Println(*msgPointer)
}

func changeMessage(aPointer *string, newMessage string) {
	*aPointer = newMessage
}
