# A pointer :
    =>
    is a powerful feature in the Golang programming language that allows you to work with the memory addresses.

In Golang, Dereferencing a pointer means accessing the value stored at the memory address pointer is pointing to. This concept is important in Golang as it enables you to optimize performance, share data between functions without copying, and manipulate data.


To declare a pointer, we use the * operator, and to get the memory address of a variable, we can use &(AND) operator.


# Example:
    =>
      package main
      import "fmt"
      
      func main() {
         var x int = 64   
         var p *int        
         p = &x          
      
         fmt.Println("Value of x=", x)
         fmt.Println("Address of x=", &x)
         fmt.Println("Value of p=", p)
         fmt.Println("Dereferenced value of p=", *p)
         *p = 21 
         fmt.Println("New value of x=", x)
      }


        package main
        import "fmt"
        
        type Student struct {
           name string
           age  int
        }
        
        func main() {
           p := &Student{"Indu", 18}
           fmt.Println("Name=", p.name)
           fmt.Println("Age=", (*p).age)
        }

## Passing and Dereferencing Pointers:
    => 
        package main
        import "fmt"
        
        func updateValue(val *int) {
           *val = 100
        }
        
        func main() {
           num := 25
           updateValue(&num)
           fmt.Println("Updated Value=", num)
        }


# Structs:
        => Structure is another user-defined data type available in Go programming, which allows you to combine data items of different kinds.

# Example:
        =>
            package main 
            import "fmt"
            
            type Books struct {
               title string
               author string
               subject string
               book_id int
            }
            func main() {
               var Book1 Books    /* Declare Book1 of type Book */
               var Book2 Books    /* Declare Book2 of type Book */
             
               /* book 1 specification */
               Book1.title = "Go Programming"
               Book1.author = "Mahesh Kumar"
               Book1.subject = "Go Programming Tutorial"
               Book1.book_id = 6495407
            
               /* book 2 specification */
               Book2.title = "Telecom Billing"
               Book2.author = "Zara Ali"
               Book2.subject = "Telecom Billing Tutorial"
               Book2.book_id = 6495700
             
               /* print Book1 info */
               fmt.Printf( "Book 1 title : %s\n", Book1.title)
               fmt.Printf( "Book 1 author : %s\n", Book1.author)
               fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
               fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)
            
               /* print Book2 info */
               fmt.Printf( "Book 2 title : %s\n", Book2.title)
               fmt.Printf( "Book 2 author : %s\n", Book2.author)
               fmt.Printf( "Book 2 subject : %s\n", Book2.subject)
               fmt.Printf( "Book 2 book_id : %d\n", Book2.book_id)
            }

# Structures as Function Arguments:
        =>
            package main
            import "fmt"
            
            type Books struct {
               title string
               author string
               subject string
               book_id int
            }
            func main() {
               var Book1 Books    /* Declare Book1 of type Book */
               var Book2 Books    /* Declare Book2 of type Book */
             
               /* book 1 specification */
               Book1.title = "Go Programming"
               Book1.author = "Mahesh Kumar"
               Book1.subject = "Go Programming Tutorial"
               Book1.book_id = 6495407
            
               /* book 2 specification */
               Book2.title = "Telecom Billing"
               Book2.author = "Zara Ali"
               Book2.subject = "Telecom Billing Tutorial"
               Book2.book_id = 6495700
             
               /* print Book1 info */
               printBook(Book1)
            
               /* print Book2 info */
               printBook(Book2)
            }
            func printBook( book Books ) {
               fmt.Printf( "Book title : %s\n", book.title);
               fmt.Printf( "Book author : %s\n", book.author);
               fmt.Printf( "Book subject : %s\n", book.subject);
               fmt.Printf( "Book book_id : %d\n", book.book_id);
            }


                
                package main
                
                import "fmt"
                
                type Books struct {
                   title string
                   author string
                   subject string
                   book_id int
                }
                func main() {
                   var Book1 Books   /* Declare Book1 of type Book */
                   var Book2 Books   /* Declare Book2 of type Book */
                 
                   /* book 1 specification */
                   Book1.title = "Go Programming"
                   Book1.author = "Mahesh Kumar"
                   Book1.subject = "Go Programming Tutorial"
                   Book1.book_id = 6495407
                
                   /* book 2 specification */
                   Book2.title = "Telecom Billing"
                   Book2.author = "Zara Ali"
                   Book2.subject = "Telecom Billing Tutorial"
                   Book2.book_id = 6495700
                 
                   /* print Book1 info */
                   printBook(&Book1)
                
                   /* print Book2 info */
                   printBook(&Book2)
                }
                func printBook( book *Books ) {
                   fmt.Printf( "Book title : %s\n", book.title);
                   fmt.Printf( "Book author : %s\n", book.author);
                   fmt.Printf( "Book subject : %s\n", book.subject);
                   fmt.Printf( "Book book_id : %d\n", book.book_id);
                }
