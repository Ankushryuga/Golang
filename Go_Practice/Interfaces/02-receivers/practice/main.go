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

	// Task 3: Temperature converter
	temp := Celsius(50)
	fmt.Println("Fahernheit value", temp.ToFahrenheit())
	fmt.Println("Kelvin Value", temp.ToKelvin())

	// Task 4: Todo
	todo := Todo{Title: "Go receiver"}
	todo.MarkCompleted()
	todo.Print()
	todo.MarkIncomplete()
	todo.Print()

	// Task 5: Employee Salary
	employee := Employee{Name: "Kush", Salary: 2000}
	employee.IncreaseSalary(40)
	employee.PrintSalary()
}
