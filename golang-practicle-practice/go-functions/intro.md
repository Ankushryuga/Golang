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

    


## Methods: When to use and when not to use:
    =>
    ## When to use::
    1. You want behavior tied to a type (struct or  named type)
        - If the function logically operates on or modifies a value (e.g, struct), it should be a method.
        Example: type Circle struct{
            Radius float64    
        }
        func (c Circle) Area() float64{
            return 3.14 * c.Radius * c.Radius
        }
    
    2. When you want to implement an interface.
        - Interfaces in Go are satisfied implicitly by implementing required mmethods.    

        type Stringer interface{
            String() string
        }

        type Person struct{
            Name string
        }

        func (p Person) String() string{
            return "Person: "+ p.Name
        }
        
     3. When you need to mutate the receiver (use a pointer receiver).
         - If the method modifies the state of receiver, use a pointer.
         func (p *Person) SetName(name String){
             p.Name=name
         }

     4. You want to group behavior with data for clarity:
         - Methods helps encapsulates functionality within a type, improving code orgnaization and reaability.

    ## When to Not to Use::
     1. The function does not operate on a specific type.
         - Use standalone functions if the behavior is general-purpose or not tied to a type.

     2.  The function operates on multiple unrelated types.
