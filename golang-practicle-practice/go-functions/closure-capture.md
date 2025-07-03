# Closure Capture:
    => Closure capture refers to the way an anonymous function (or closure) "remembers" the variables from its surrounding scope, even after that
    scope has exited.
    This is powerful feature but can be tricky - especially with goroutines or loops.

        package main
        import "fmt"
        
        func main() {
        	msg := "Hello"
        	
        	closure := func() {
        		fmt.Println(msg)
        	}
        	
        	msg = "Hi"
        	closure() // Output: Hi
        }


# closure capture summary:
    =>
    | Case                | Behavior                                |
    | ------------------- | --------------------------------------- |
    | Captures variable   | Holds reference, not value              |
    | Loop variable       | Must be passed explicitly to avoid bugs |
    | goroutine + closure | Most common cause of weird bugs         |


# ✅ Use Cases for Closures
    1. Creating function factories (return functions from functions)
    2. Capturing shared state across calls
    3. Implementing callbacks or filters
    4. Simplifying concurrency patterns
