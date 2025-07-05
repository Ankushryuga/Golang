## NOTE:::
    =>
    range only works properly if the channel is closed, because it will keep waiting for new value if it's stil open..


## Context Switching in Go:
    => In Go, Context switching happens b/w goroutines, which are user-level lightweight threads.

## How Go handles it:
    =>
    1. Go Uses a scheduler that is part of the Go runtime.
    2. It Maps thousands of goroutines onto a smaller number of OS threads
    3. Context switching b/w goroutines are much cheaper than OS Threads.


## Goroutine Scheduling (M:N Model):
    =>
    1. "M" OS threads
    2. "N" goroutines


# Go Context Switching is Cheaper b/c:
    =>
    1. No need to ask the OS to switch threads.
    2. Go's scheduler keeps most data in user-space.
    3. Stack sizes are small and grow dynamically (starting at 2KB).


