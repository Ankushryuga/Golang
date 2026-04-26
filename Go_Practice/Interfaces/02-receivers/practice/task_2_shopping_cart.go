package main

import "fmt"

type Cart struct {
	Items []string
}

// Pointer Receiver -> modifies underlying slice
func (cart *Cart) AddItem(item string) {
	// this is how increase the size of slice
	cart.Items = append(cart.Items, item)
}

// Pointer receiver -> removes last item safely
func (cart *Cart) RemoveLastItem() {
	if len(cart.Items) == 0 {
		return
	}
	// this is how we remove the last item from slice
	cart.Items = cart.Items[:len(cart.Items)-1]
}

func (cart Cart) PrintItems() {
	for _, item := range cart.Items {
		fmt.Println(item)
	}
}
