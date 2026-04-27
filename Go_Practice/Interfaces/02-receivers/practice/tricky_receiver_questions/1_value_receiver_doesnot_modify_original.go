package main

// Problem 1: Value receiver Does not Modify Original

type User struct {
	Name string
}

// Here ChangeName uses a value receiver. It modifies only a copy.
func (u User) ChangeName(name string) {
	u.Name = name
}
