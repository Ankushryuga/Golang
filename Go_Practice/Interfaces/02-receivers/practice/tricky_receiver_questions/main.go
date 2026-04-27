package main

import "fmt"

func main() {
	// Problem 1: Value receiver Does not modify original.
	user := User{Name: "Ankush"}
	user.ChangeName("Raj")
	fmt.Println(user.Name) // O/P: Ankush

	// Problem 2: Pointer Receier can be called on Value
	counter := Counter{Value: 1}
	counter.Increment()
	fmt.Println(counter.Value) // 2

	// Problem 3: Pointer Receiver Cannot be called on non-addressable value
	// GetUser().ChangeName("Raj")	// Compiler error because Go cannot do : &GetUser()
	/**
	// Correct Version
	user := GetUser()
	user.ChangeName("Raj")
	*/

	// Problem 4: Nil Pointer Receiver Can be valid
	var user_nil *User_Nil_Pointer_Receiver
	user_nil.PrintName() // o/p: user is nil

	// Problem 5: Method Set and Interface Confusion
	var p Printer
	user_method_set := User_method_set_interface_example{Name: "Ankush"}
	// p=user_method_set // Wrong, because "Print()" is defined on "*User_method_set_interface_example", not "User_method_set_interface_example", so only "*User_method_set_interface_example" implements "printer"
	p = &user_method_set
	p.Print()

	// Problem 6: Value Receiver Method is Available to Both Value and Pointer.
	value_receiver := Value_Receiver{Name: "Ankush"}
	value_receiver.Print() // Ankush
	ptr := &value_receiver
	ptr.Print() // Ankush

	// Problem 7: Pointer Receiver with map value:
	// ############# Wrong Approach
	/**
	// this will give compile error
	pointer_receiver := map[string]Pointer_Receiver_With_Map_Value{
		"one": {Name: "Ankush"},
	}
	pointer_receiver["one"].ChangeName["Raj"]

	// why?
	// Map index values are not addressable
	// Go cannot take the address of:
	// pointer_receiver["one"]
	*/
	// ############# Correct Approach
	/**
	// Appraoch 1:
	pointer_receive := pointer_receiver["one"]
	pointer_receive.ChangeName("Raj")
	pointer_receiver["one"]=pointer_receive

	// Appraoch 2: Store pointers in the map.
	pointer_receiver := map[string]*Pointer_Receiver_With_Map_Value{
	"one": {Name: "Ankush"},
	}
	pointer_receiver["one"].ChangeName("Raj")
	*/
}
