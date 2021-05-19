package main

import (
	"demo17/employee"
)

func main() {
	e := employee.Init("M9nood", "toei", 30, 20)
	e.LeavesRemaining()
	e = employee.Init("Dev", "toei", 30, 20)
	e.LeavesRemaining()
}
