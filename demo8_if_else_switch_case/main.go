package main

import "fmt"

func main() {
	someValue := 10
	if someValue == 10 {
		fmt.Println("== 10")
	} else {
		fmt.Println("!= 10")
	}

	if someValue > 10 || someValue < 2 {
		fmt.Println("someValue > 10 || someValue < 2")
	} else {
		fmt.Println("NOT someValue > 10 || someValue < 2")
	}

	// if result := doSomething(); result == "ok"
	if doSomething() == "ok" {
		fmt.Println("ok")
	} else {
		fmt.Println("No ok")
	}

	fnSwitchCase()
}

func doSomething() string {
	return "oks"
}

func fnSwitchCase() {
	index := 3
	switch index {
	case 0:
		fmt.Println("switch = 0")
		break
	case 1:
		fmt.Println("switch = 1")
		break
	case 2:
		fmt.Println("switch = 2")
	default:
		fmt.Println("something else")
		break

	}
}
