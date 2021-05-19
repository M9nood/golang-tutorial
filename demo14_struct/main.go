package main

import "fmt"

func main() {
	var p1 product
	p1 = product{name: "Chair", price: 100, stock: 10}
	p1.name = "desk"
	show(p1)
	update(&p1)
	show(p1)

	// p1 = p1.clear()
	p1 = p1.setDiscount(10)
	show(p1)
}

func (p product) clear() product {
	p.price = 0
	p.stock = 0
	return p
}

func (p product) setDiscount(d int) product {
	p.price = p.price - d
	return p
}

type product struct {
	name  string
	price int
	stock int
}

func show(p product) {
	fmt.Println(p)
}

func update(p *product) {
	p.price = p.price + 50
	p.stock = 5
}
