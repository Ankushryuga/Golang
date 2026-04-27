package main

type User_Non_Addressable struct {
	Name string
}

func (u *User_Non_Addressable) ChangeName(name string) {
	u.Name = name
}

func GetUser() User_Non_Addressable {
	// GetUser() returns a temporary value. It is not addressable.
	return User_Non_Addressable{Name: "Kush"}
}
