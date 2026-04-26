package main

import "fmt"

// Example 1: Print User Details Example
type User struct {
	Name string
	Age  int
}

func (u User) PrintDetails() {
	fmt.Println("Name:", u.Name)
	fmt.Println("Age:", u.Age)
}

// Example 2: Area of Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Example 3: Full Name Method
type Employee struct {
	FirstName string
	LastName  string
}

func (e Employee) FullName() string {
	return e.FirstName + " " + e.LastName
}

func main() {

	// Example 1: Print user details
	user := User{Name: "Ankush", Age: 25}
	user.PrintDetails()

	// Example 2: Area of Rectangle
	dimension := Rectangle{Width: 30, Height: 40}
	fmt.Println(dimension.Area())

	// Example 3: Full Name method
	employee := Employee{FirstName: "Ankush", LastName: "Raj"}
	fmt.Println(employee.FullName())

	// Pointer Receiver Example
	// Example 4: Change Name
	user_pointer := User_Pointer{Name: "Ankush"}
	user_pointer.ChangeName("New Name Changed")
	fmt.Println("Pointer Receiver", user_pointer.Name)

	// Example 5: Coutner Increment:
	counter := Counter{Value: 0}
	counter.Increment()
	counter.Increment()
	fmt.Println(counter.Value) // 2

	// Example 6: Bank Account Deposit and Withdraw
	bank := BankAccount{Balance: 5000}
	// Use pointer receiver for "Deposit" and "Withdraw", because they modify balance
	bank.Deposit(5444)
	bank.Withdraw(7000)
	// Use value receiver for "GetBalance", becuase it only reads balance.
	fmt.Println("Bank Balance", bank.GetBalance())

	// Example 7: Mix student receiver Example
	student := Student{Name: "Ankush", Marks: 50}
	student.PrintResult()
	student.AddGraceMarks(40)
	student.PrintResult()
	// Note: Pointer receiver modifies the original struct.

}
