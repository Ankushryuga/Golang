/****
panic is Go's built-in mechanism for handling unexpected runtime errors. when a panic occurs, it immediately stops execution of the current function and begins 
unwinding the stack, running any defer statements along the way. If the panic reaches the top of goroutine and isn't recovered, the program crashes.
*/
