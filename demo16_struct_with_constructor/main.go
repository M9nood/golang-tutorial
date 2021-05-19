package main

import (
	"demo16/employee"
)

func main() {
	e := employee.New("M9nood", "toei", 30, 20)
	e.LeavesRemaining()
}
