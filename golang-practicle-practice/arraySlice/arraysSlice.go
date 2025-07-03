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



	var num []int		//a slice of unspecified size.. similar to num == []int {0,0,0,0,0}
	
	num1 :=make([]int, 4, 7)
	//if a slice is declared with no inputs, then by default, it is initialized as nil, its length and capacity are zero.

	if num == nil{
		fmt.Println("slice is nil")
	}


	//1. Slice declaration and initialization with values:
	s1 := []int{1,2,3,4,5}

	//2. Declare without initializing (nil slice)
	var s2 []int

	//using make
	s3 := make([]int, 3, 4)		//len 3 and cap 4
	s4 := make([]int 3)		//len 3 and cap 3


	//apend
	var s5 []int
	s5=append(s5, 10,20, 30)


	//slice from an array:
	arr := [5]int{1,2,3,4,5}
	s6 := arr[1:4]	//2,3,4


	s7 := arr[1:]	//2,3,4,5

	
}
