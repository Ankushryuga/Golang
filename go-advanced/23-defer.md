# defer:
      =>
      In Golang, the defer keyword is used to delay the execution of a function until the surrounding function completes. The deferred function calls are executed in Last-In-First-Out (LIFO) order. That means the most recently deferred function is executed first, followed by the second most recently deferred function, and so on.

      The defer keyword is useful when we want to ensure that certain operations are performed before a function returns, regardless of whether an error occurs or not. This can help simplify error handling and make code more readable.

# Syntax:  
      => defer functionName(arguments)



# Deferred Functions and Panics
      =>
      The defer keyword is particularly useful in cases where we want to recover from a panic. A panic is an error that occurs at runtime when a program cannot continue executing. When a panic occurs, the program is terminated and an error message is printed to the console.
      By using the defer keyword, we can ensure that certain cleanup functions are executed even in the event of a panic.
      
#   Example:
      =>
      package main
      func main() {
         defer cleanup()
         // Perform some operations that may cause a panic
      }
      func cleanup() {
         if r := recover(); r != nil {
            // Log the error
         }
         // Clean up resources
      }
