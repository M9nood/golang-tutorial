package main

import "fmt"

func main() {
	var array1 []int = []int{1, 2, 3, 4}
	var array2 = []int{5, 6, 7, 8}
	var array3 [3]string

	fmt.Println(array1[0])

	for _, item := range array1 {
		fmt.Print(item, ",")
	}
	fmt.Print("\n")
	for _, item := range array2 {
		fmt.Print(item, ",")
	}
	fmt.Print("\n")

	array3[0], array3[1], array3[2] = "android", "ios", "vue"
	for _, item := range array3 {
		fmt.Print(item, ",")
	}
	fmt.Print("\n")
}
