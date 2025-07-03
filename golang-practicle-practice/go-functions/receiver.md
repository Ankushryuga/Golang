# Receiver:
    =>
    Methods are like functions but with a key difference: they have a receiver argument, which allows access to the receiver's properties.
    The receiver can be a struct or non struct type, but both must be in the same package. 
    Methods cannot be created for types defined in other pacakges, including built-in types like int or string;
    otherwise compiler will raise an error.



# Table of Content:
    =>
    1. Methods with struct Type Receiver
    2. Methods with non-struct type receiver
    3. Methods with Pointer receiver
    4. Methods Accepting Both pointer and value



# Methods with struct type receiver:
    =>
    you define a method where the receiver is of a struct type, the receiver is accessible inside the method.

# Methods with Non-Struct Type Receiver:
    =>
    you cannot define a method with a receiver type from another pacakge.


    package main
    import "fmt"
    //creating a custom type based on int
    type number int
    func (n number) square() number{
        return n*n
    }

    func main(){
    a := number(4)
    b := a.square()

    fmt.Println("Square of ",a, "is", b) 
    }


# Methods with pointer receiver:
    =>
    Methods can have pointer receiver, this allows changes made in the method to reflect in the caller, which is not possible with value receiver.

    ## Syntax:
    func (g *Type) method_name(...Type)Type{ //code}

    package main
    import "fmt"

    type person struct{
        name string
    }

    //method with pointer receiver to modify data::
    func (p *person) changeName(name string){
        p.name=name    
    }

    func main(){
        a := person{name:"ankush"}

        fmt.Println("Before", a.name)

        a.changeName("banti")
        fmt.Println("After", a.name)
    }


# Method accepting both pointer and value:
    => 
    Go methods can accept both value and pointer receivers, you can pass either a pointer or a value, and the method will handle it accordingly.
        
        package main
        import "fmt"
        type person struct{
            name string
        }
        func (p * person) updateName(name string){
            p.name=name
        }

        func (p person) showName(){
            fmt.Println("Name: ", p.name)
        }

        func main(){
            a := person{name:"ankush"}
            a.updateName("banti")
            fmt.Println("After pointer method :", a.name)
            (&a).showName()
        }
