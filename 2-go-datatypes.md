# Types of Go Data Types:
    =>
    1. Boolean Types: true, false.
    2. Numeric Types: integer(uint8, uint16, uint32, uint64:=> (u-> unsigned). int8, int16, int32, int64:=> signed) floating (float32, float64), complex64 (complex number with float32 real and imaginary parts, complex128 (float64 real and imaginary parts)),
    Other: (byte (same as uint8)), rune(same as int32), uint(32 or 64 bits), int(same size as uint), uintptr(an unsigned integer to store the uninterpreted bits of a pointer value).
    3. String Types:
    4. Derived Types: 
      4.1. Pointer types: it stores the memory address of another variable.
      4.2. Array Types: fixed-size sequeunce of elements of the same type.
      4.3. Structure Types: Its a collection of fields, each with a name and a type. The structure type allow you to store different types of value.
      4.4. Union Types: Go don't provide the support of unions, but unions can be used as interface {} to hold any type of value.
      4.5. Function Types: Used to organize and structure the code, allowing grouping of the related code together in purpose to make it reusable and easy to maintain.
      4.6. Slice Types: its dynamically size and flexible view of an array.,
      4.7. Map Types: its a unordered collection of key-value pairs.
      4.8. Channel Types: channel are useful when you are working with goroutines; these types are used for communication b/w goroutines.



# Example:
    1. Pointer types:
      var x int = 42
      var ptr *int=&x;

    2. Array type:
      var arr [3]int = [3]int{10, 20, 40}

    3. Structure Types:
      type Student struct{
        Name string
        Age int
      }

    4. Union type: 
       var u interface {}="Ankush"

    5. Functional Type:
        var AddTwoNums=func(a, b int) int {
            return a+b 
        }

    6. Slice Type: 
        var arr []int=[]int{1,3,4,5}

    7. Map Types:
        var countryCodes=map[string]string{"USA": "+1","India": "+91"}

    8. Channel Types:
        ch := make(chan int)
          
