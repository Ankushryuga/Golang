package main

import "fmt"

func main() {
	// Task 1: Product Design
	product := Product{Name: "Motorolla", Price: 4000}
	product.ApplyDiscount(1000)
	fmt.Println("Final Price", product.FinalPrice())

	// Task 2: Shopping Cart
	cart := Cart{}
	cart.AddItem("Laptop Mouse")
	cart.AddItem("Keyboard")
	cart.PrintItems()
}
