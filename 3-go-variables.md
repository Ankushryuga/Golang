# variable :
    =>
    are a storage area that the programs can manipulate, each variable in Go has a specific type, which determine the size
    and layout of the variable's memory, the range of values that can be stored within that memory, and the set of operations that can
    be applied to the variable.

# variable definition in Go:
    =>
    syntax:  
    var variable_list optional_data_type;
    Example:
    var i, j, k int;
    var c, ch byte;
    var f, salary float32;
    d = 32;
    e = 5, g=4; 


# Static Type Declaration in Go
    =>
    A static type variable declaration provides assurance to the compiler that there is one variable available with the given type and name,
    so that the compiler can proceed for further compilation without requiring the complete detail of the variable.
    A variable declaration has its meaning at the time of compilation only, the compiler needs the actual variable declaration at the time of linking of the program.


# Dynamic Type Declaration / Type Inference in Go:
    =>
    A dynamic type variable declaration requires the compiler to interpret the type of the variable based on the value passed to it.
    The compiler doesn't require a variable to have type statically as a necessary requirement.


# Mixed variable declaration in Go,
    => Variable of different types can be declared in one go using type inference.
    Example:
    var a, b, c = 3, 4, "foo"
    fmt.Printf("b is of type %T\n", b)   fmt.Printf("c is of type %T\n", c)


# Ivalues and rvalues in Go:
    => There are 2 kinds of expressions in Go - 
      1. Ivalue -> expressions that refer to a memory location called "Ivalue" expression. An Ivalue may appear as either the left-hand or right-hand side of an assignment.
      2. rvalue -> it refers to a data value that is stored at some address in memory, an rvalue is an expression that cannot have a value assigned to it which means an 
                    rvalue may appear on the right-but not left-hand side of an assignment.

                
    
