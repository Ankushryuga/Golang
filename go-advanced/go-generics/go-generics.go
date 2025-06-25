package main

import "fmt"

func printSliceGenerics[T any](value []T) {
	for _, v := range value {
		fmt.Println(v)
	}
}
func main() {
	intSlice := []int{1, 2, 3, 4}
	stringSlice := []string{"hello", "boy"}

	printSliceGenerics(intSlice)
	printSliceGenerics(stringSlice)
}
