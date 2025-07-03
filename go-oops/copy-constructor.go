package main

import "fmt"

type User struct{
    Name string
    Email string
    Age int
}

//Copy constructor..
func NewUserCopy(u *User) *User{
    return &User{
        Name:u.Name,
        Email: u.Email,
        Age: u.Age,
    }
}

func main(){
    original := &User{Name:"Ankush", Email:"ankush.raj@example.com", Age:27}
    
    copy := NewUserCopy(original)
    copy.Name="Copied Constructor"
    
    fmt.Println("Original: ", original.Name)
    fmt.Println("Copy: ", copy.Name)
}
