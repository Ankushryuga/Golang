# Functions:
    => also known as method, sub-routine, or procedure.

# syntax:
    =>
    func function_name( [parameter list] ) [return -types] {
        //body code..
    }
    ///single return type..
    func sum(x, y int) int{}
    
    //multiple return type..
    func swap(x, y int)(int, int){}


## call by value:
    =>
    Call by value method of passing arguments to a function copies the actual value of an argument into the formal parameter of the function.
        func callbyvalue(x, y int) int{}


# Call by reference:
    =>
    call by reference method of passing arguments to a function copies the address of an argument into the formal parameter.
    func swap(x *int, y *int){}

    



    
