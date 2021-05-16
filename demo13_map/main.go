package main

import "fmt"

func main() {
	var numbers = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("", numbers["one"])

	var colors = make(map[string]string)
	colors["red"] = "#f00"
	colors["green"] = "#0f0"
	colors["blue"] = "#00f"
	fmt.Println("", colors)
	fmt.Println("", colors["green"])

	var courses = make(map[string]map[string]int)
	courses["android"] = make(map[string]int)
	courses["android"]["price"] = 200
	courses["android"]["code"] = 101

	courses["go"] = map[string]int{"price": 300, "code": 102}

	fmt.Println("", courses)
}
