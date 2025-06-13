# Slice is an abstraction over Go Array, 
    =>  It provides many utility functions required on Array and is widely used in Go programming.

# Defining a slice
    => To define a slice, you can declare it as an array without specifying its size. Alternatively, you can use make function to create a slice.

# Example:
    =>
    var numbers []int /* a slice of unspecified size */
    /* numbers == []int{0,0,0,0,0}*/
    numbers = make([]int,5,5) /* a slice of length 5 and capacity 5*/

# Example:
    =>
    package main
    import "fmt"
    func main() {
       var numbers []int
       printSlice(numbers)
       
       if(numbers == nil){
          fmt.Printf("slice is nil")
       }
    }
    func printSlice(x []int){
       fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x),x)
    }

# Subslicing:
    package main
    import "fmt"
    func main() {
       /* create a slice */
       numbers := []int{0,1,2,3,4,5,6,7,8}   
       printSlice(numbers)
       
       /* print the original slice */
       fmt.Println("numbers ==", numbers)
       
       /* print the sub slice starting from index 1(included) to index 4(excluded)*/
       fmt.Println("numbers[1:4] ==", numbers[1:4])
       
       /* missing lower bound implies 0*/
       fmt.Println("numbers[:3] ==", numbers[:3])
       
       /* missing upper bound implies len(s)*/
       fmt.Println("numbers[4:] ==", numbers[4:])
       
       numbers1 := make([]int,0,5)
       printSlice(numbers1)
       
       /* print the sub slice starting from index 0(included) to index 2(excluded) */
       number2 := numbers[:2]
       printSlice(number2)
       
       /* print the sub slice starting from index 2(included) to index 5(excluded) */
       number3 := numbers[2:5]
       printSlice(number3)
       
    }
    func printSlice(x []int){
       fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x),x)
    }


## Slice of Slices:
        =>var matrix [][]int
        // Method 1: Literal initialization
            grid := [][]int{
               {4, 5, 6},       // First row (slice)
               {7, 8},          // Second row (can be different length)
            }
            // Method 2: Using make()
            rows := 3
            matrix := make([][]int, rows)      // Outer slice
            for i := range matrix {
               matrix[i] = make([]int, 2)     // Each row has 2 columns (can vary)
            }
# Initialize a slice of slices:

            matrix := make([][]int, 3) // Create an outer slice with 3 inner slices
            for i := range matrix {
                matrix[i] = make([]int, 4) // Each inner slice has 4 elements
            }
