package main

import "fmt"

type game struct {
	title string
	price float64
}

func (g game) print() {
	fmt.Printf("%-15s: $%.2f\n", g.title, g.price)
}

/**
You can use same method name among different types.
you don't need to type `printGame`, it's just: `print`.

// previous old code.
func (g game) printGame() {
fmt.Printf("%-15s: $%.2f\n", g.title, g.price)
}

NOTE: you cannot use the same function name within the same package.

func printGame(g game) {
	fmt.Printf("%-15s: $%.2f\n", g.title, g.price)
}
*/
