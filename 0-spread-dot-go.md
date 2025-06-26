## ... (ellipsis) in Go is a special syntax used in two different but related ways:
        =>
        1. Variadic Parameters (Function definition):=> ... to define a function that accepts any number of arguments of a specified type known as a variadic function:
        # Example:
            func  sum(nums ...int) int {    //internally nums is a slice.
                total := 0 
                for _, num := range nums{
                    total += num
                }
                return total
            }

        // call function:
        sum(1,2,3)
        sum(1,2,3,4)


## slice expansion: (Function call)
        => in function calls, ... is used to unpack or expand a slice into individual arguments - especially useful when passing a slice to a varidic function.

        ## Example:
        nums := []int{1,2,3}
        sum(nums...)    // expand to sum(1,2,3)


# | Usage             | Meaning                                      | Example                 |
| ----------------- | -------------------------------------------- | ----------------------- |
| `func f(x ...T)`  | Variadic parameter (accepts many args)       | `func log(v ...string)` |
| `f(x...)`         | Expands slice into arguments (function call) | `sum([]int{1,2}...)`    |
| `append(a, b...)` | Expands slice `b` into elements for append   | `append(s1, s2...)`     |



# use ...:
    1. To accept unlimited arguments (...int)
    2. To pass a slice to a variadic function (slice... )
    3. To expand and merge slices (append(a, b...)) 




