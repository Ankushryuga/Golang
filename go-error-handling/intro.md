# Error Handling:
    => Errors: these are exception-based mechanisms, when error occurs, the execution of a program stops completely with a built-in error message, 

## we can handle errors using:
    =>
    1. New() function
    2. Errof() function

# The error Interface in Go
    => In Go, errors are represented by the error interface, which is a built-in interface with a single method, Error() string. Any type that implements this method can be used as an error.

# definition of the error interface::
    =>
    type error interface{
        Error() string
    }
    => you can create custom error types by implementing the error interface, or you can use the standard errors package to create simple error values.(errors.New() function is used to create new error).


# Common Error Handling Patterns:
    =>
    1. Returning Errors: it's a common pattern to return errors as the last value in the function signature. If a function can return an error, its signature should include an error as the last return value.
            => Example:
            func divide (a, b float64) (float64, error){
                if b==0{
                    return 0, errors.New("division by zero is not allowed")
                }
                return a/b, nill
            }


    2. Handling Errors: When calling a function that returns an error, it's essential to check the error and handle it appropriately. This Idiomatic way to do this in Go is to use the 
        ** if err != nill pattern: **
        # Example:

        result, err := divide(19, 0)
        if err != nil{
            fmt.Println("Error: ", err)
        }else{
            fmt.Println("Result: ", result)
        }


    3. Wrapping Errors: when you want to add additional context to an error before returning it to the caller, you can do this using the fmt.Errorf() function, which allows you to create a formatted error message and wrap the original error:
        => Example:
        func processFile(filename string) error{
            file, err := os.Open(filename)
            if err != nil {
                return fmt.Errorf("failed to open file: %w", err)
            }
            defer file.Close()
            return nil
        }


# How do you handle Panic and Recover from them in Go?
    => 
    Go provides, built-in panic() and recover() functions can be used. When an error occurs,panic() is called and the program execution stops, you can use the defer statement to call recover(), which stops the panic and resumes execution from the point of the nearest enclosing function call.
        
    It's important to note that you should only use panic and recover for exceptional cases such as when a program encounters unexpected errors. It's not meant to handle normal control flow or act as a replacement for checking errors.
    
    Using these functions properly can help you ensure programs remain robust and reliable.



# defer statement in Golang, give an example of a deferred function?
    =>
    The defer statement in Golang is used to postpone the execution of a function until the surrounding function completes. It is often used when you want to make sure some cleanup tasks are performed before exiting a function, regardless of errors or other conditions.
    Here's an example of a deferred function's call:
