package main

// Problem 2: Pointer Receiver Can be called on Value
type Counter struct {
	Value int
}

func (c *Counter) Increment() {
	c.Value++
}
