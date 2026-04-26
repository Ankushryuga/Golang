package main

// Exampple 1: Change Name
type User_Pointer struct {
	Name string
}

func (u *User_Pointer) ChangeName(newName string) {
	u.Name = newName
}

// Example 2: Counter increment
type Counter struct {
	Value int
}

func (c *Counter) Increment() {
	c.Value++
}

// Example 3: Bank Account Deposit and Withdraw
type BankAccount struct {
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
}

func (b *BankAccount) Withdraw(amount float64) {
	b.Balance -= amount
}

func (b BankAccount) GetBalance() float64 {
	return b.Balance
}
