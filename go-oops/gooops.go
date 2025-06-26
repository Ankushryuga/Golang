package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Speak() {
	fmt.Println(a.Name, "makes s sound")
}

type Dog struct {
	Animal //embedded struct.(inherit behaviour)
	Breed  string
}

func (d Dog) Speak() {
	fmt.Println(d.Name, "barks! Breed:", d.Breed)
}

type Speaker interface {
	Speak()
}

func makeItSpeak(s Speaker) {
	s.Speak()
}

func main() {
	a := Animal{Name: "Go oops"}
	d := Dog{Animal: Animal{Name: "Broun"}, Breed: "Lavi"}


	makeItSpeak(a)
	makeItSpeak(d)
}
