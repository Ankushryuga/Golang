# Golang practice codes.


# Go memory management:
    =>
    ✅ It depends on whether the allocated value “escapes” the current function.
    
    Example 1: Stack allocation (does not escape)
    go
    Copy
    Edit
    func foo() {
        p := new(int)
        *p = 42
        fmt.Println(*p)
    }
    In this case, p does not escape the function — Go compiler may allocate it on the stack.


# Heap Allocation (escapes):
    =>
    func bar() *int {
    p := new(int)
    *p = 100
    return p // p escapes the function
    }
    Here, p is returned from the function, so Go must allocate it on the heap, or it will be garbage-collected too early.



# 📘 How does Go decide?
    Go uses escape analysis to decide:
    
    If a variable stays within the function, it stays on the stack.
    
    If it escapes (e.g., returned, stored in heap struct, used in a goroutine), it’s allocated on the heap.
