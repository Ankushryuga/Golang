# Slices

Slices are one of the most central Go DS.

A slice is a lightweight view over an underlying array.

- Slice Internals first
  A Go slice is a "**three-field struct**" under the hood:

```bash
┌─────────┬──────┬─────┐
│  ptr    │  len │ cap │
└─────────┴──────┴─────┘
    │
    ▼
[ 1 | 2 | 3 | 4 ]   ← backing array
```

    - **ptr** - pointer to the first element in the backing array
    - **len** - how many elements are visible
    - **cap** - how many elements the array can hold before reallocation

Multiple slices can share the same backing array. This is the root of both the power and traps in. these patterns.

## Basic slice

```go
nums := []int{1,2,3}
fmt.Println(nums)
```

### Length and Capacity

```go
s := []int{10, 20, 30}
fmt.Println(len(s))
fmt.Println(cap(s))
```

A slice has:

- pointer to underlying array
- length
- capacity

This is critical for performance and bugs.

### Appending

```go
s := []int{1,2}
s = append(s, 3)
```

Always assign result of `append`, because it may allocate a new array.

- Bad:

```go
append(s, 3)
```

- Good:

```go
s = append(s, 3)
```

### Slicing

```go
a := []int{1, 2, 3, 4, 5}
b := a[1:4] // [2 3 4]
```

This shares underlying memory with `a`.

That means modifying `b` may affect `a`.

```go
b[0] = 99
fmt.Println(a)      // [1 99 3 4 5]
```

> [!NOTE]
> Shared backing arrays cause subtle bugs

## Copying a slice

To get an independent copy:

```go
src := []int{1,2,3}
dst := make([]int, len(src))
copy(dst, src)
```

Now changes to `dst` do not affect src

### Nil slice vs empty slice

```go
var a []int             // nil slice
b := []int{}            // empty slice
c := make([]int, 0)     // empty slice
```

Both can often be used similarly:

```go
fmt.Println(len(a))     // 0
fmt.Println(len(b))     // 0
```

But:

- `a == nil` is true
- `b == nil` is false

For JSON:

- nil slice may marshal as `null`
- empty slice may marshal as `[]`

> [!CAUTION]
> Choose intentionally depending on API semantics

### Preallocation

When size is known or estimated, preallocate.

```go
items := make([]string, 0, 100)
```

This reduces allocation during append.

> [!NOTE]
> preallocate in hot paths or known-size workloads, but don't obsess everywhere.

### Range over slices

```go
nums := []int{10, 20, 30}

for i, v := range nums{
    fmt.Println(i, v)
}
```

- If you only need value:

```go
for _, v := range nums{
    fmt.Println(v)
}
```

### Common range pitfall

The range variable is reused each iteration.

This matters especially with pointers or closures.

- Bad pattern:

```go
users := []User{{Name: "A"}, {Name: "B"}}
var ptrs []*User

for _, u := range users{
    ptrs = append(ptrs, &u)
}
```

#### WHat actually happens in a range loop?

when you write:

```go
for _, u := range users{}
```

**Go creates a single variable** `u` before the loop starts, and **copies each element into it** on every iteration. its roughly equivalent to:

```go
var u User              // created ONCE, at a fixed memory address

for i := 0; i < len(users); i++ {
    u = users[i]        // value is overwritten each iteration
}
```

The key insight: `u` lives at **one fixed address** for the entire loop duration. Its value changes, but its address does not.

> [!IMPORTANT]
> why `&u` is dangerous

```go
for _, u := range users{
    ptrs = append(ptrs, &u)         // always appending the SAME address
}

...
```

Every iteration appends the address of the same variable `u`. After the loop. `u` holds the last element (`{Name: "B"}`), so **all pointers in `ptrs` points to `"B"`**.

```bash
Memory:
    u (addr: 0xc0001)   ->  {Name: "B"}     <- Final value after loop ends.

    ptrs[0] -> 0xc0001  <- points to "B", not "A"!
    ptrs[1] -> 0xc0001  <- also points to "B"
```

This is silent - no compiler error, no panic. The bug only surfaces at runtime.

All pointers may point to same loop variable copy.

- Correct:

```go
for i := range users {
    ptrs = append(ptrs, &users[i])
}
```

`&users[i]` takes the address of the **actual element inside the backing array**, not the loop variable. Each element lives at a distinct address.

```bash

Memory:
users[0] (addr: 0xc0010) → {Name: "A"}
users[1] (addr: 0xc0020) → {Name: "B"}

ptrs[0] → 0xc0010 ✓ points to "A"
ptrs[1] → 0xc0020 ✓ points to "B"

```

**The closure variant - equally dangerous**

The same bug appears with goroutines and closures:

- Bad:

```go
for _, u := range users{
    go func() {
        fmt.Println(u.Name)         // captures the variable u, not its value
    }()
}

// likely prints "B B"  - both goroutines see the final value of u
```

- Correct:

```go
// Fix 1: Pass as argument (copies the value at that moment)
for _, u := range users {
    go func(u User) {
        fmt.Println(u.Name)
    }(u)
}

// Fix 2: Shadow the variable inside the loop.
for _, u := range users{
    u := u      // creates a NEW variable scoped to this iteration
    go func() {
        fmt.Println(u.Name)
    }()
}
```

The `u := u` trick looks odd but is idiomatic Go - the inner `u` shadows the outer one and gets its own address.

## Delete from slice

common pattern:

```go
s := []int{1, 2, 3, 4}
i := 1
s = append(s[:i], s[i+1:]...)
```

`s[:i]` is `[1]` and `s[i+1:]` is `[3, 4]`. `append` writes `3, 4` into the backing array starting at index 1, then returns a slice with `len` reduced by one.

```bash
Before: [ 1 | 2 | 3 | 4 ]
               ↑ overwrite from here
After:; [1 | 3 | 4 | 4 ]
                       ↑ stale, but outside len

        len=3, so visible: [ 1 | 3 | 4]
```

The last `4` is still in memory but invisible - `len` is now 3. This is fine, but note that `append` modifies the original backing array in place. If anything else holds a reference to the original slice, it will see the mutation.

- **Cost: O(n)** - Every element after index `i` gets shifted left by one.

Removes value `2`.

For preserving order, this is fine.
For performance when order does not matter:

```go
s[i] = s[len(s)-1]
s = s[:len(s)-1]
```

### Order-ignoring delete (swap & shrink)

```go
s[i] = s[len(s)-1]
s = s[:len(s)-1]
```

Copy the last element into the slot being deleted, then shrink `len` by one. No shifting.

```bash
Before:  [ 1 | 2 | 3 | 4 ]
                ↑ delete index 1

Step 1:  [ 1 | 4 | 3 | 4 ]   ← last element copied to index 1
Step 2:  [ 1 | 4 | 3 ]       ← len shrunk by 1
```

- **Cost: O(1)** - two assignments regardless of slice size. This is right choice when building a pool, work queue, or any unordered collection where deletions are frequent.

## Filter Pattern

```go
nums := []int{1, 2, 3, 4, 5}
result := nums[:0]

for _, n := range nums{
    if n%2 == 0 {
        result = append(result, n)
    }
}
```

`nums[:0]` creates a new slice header with `len=0` but **the same backing array and same capacity** as `nums`.

```bash
nums:    ptr→[1|2|3|4|5]  len=5  cap=5
result:  ptr→[1|2|3|4|5]  len=0  cap=5
                ↑ same backing array

# As the loop runs and `append` adds matching elements, they overwrite the front of the shared array:

Iteration 1: n=1, odd, skip
Iteration 2: n=2, even → result = [2]        array: [2|2|3|4|5]
Iteration 3: n=3, odd, skip
Iteration 4: n=4, even → result = [2,4]      array: [2|4|3|4|5]
Iteration 5: n=5, odd, skip

Final result: [2, 4]
```

Efficient, but can be tricky because it reuses original backing array.

**Safer for readability**:

```go
var result []int
for _, n := range nums{
    if n%2 == 0{
        result = append(result, n)
    }
}
```

> [!IMPORTANT]
> Prefer readable code unleass performance justifies complexity
