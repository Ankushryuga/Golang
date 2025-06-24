package main

import "fmt"

func main() {
	var array [5]int
	fmt.Println("Please enter the values")
	for i := 0; i < len(array); i++ {
		fmt.Scanln(&array[i])
	}
	fmt.Println("Your output")
	for i := 0; i < len(array); i++ {
		fmt.Println("value of i: ", array[i])
	}
	fmt.Println("Range example: ")
	for i, x := range array {
		fmt.Println("value at index ", i, "is :", x)
	}
}
