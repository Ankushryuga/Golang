package main

import "fmt"

// ✅ Struct definition
type Person struct {
	Name    string
	Age     int
	Id      int
	Vehicle string
}

// ✅ Interface with method declarations (without structs)
type interfaceExample interface {
	Information(u Person) Person
	VehicleName(car string) string
}

// ✅ Value receiver method — returns updated value
func (p Person) Information(u Person) Person {
	p.Name = u.Name
	p.Age = u.Age
	p.Id = u.Id
	p.Vehicle = u.Vehicle
	return p
}

// ✅ Pointer receiver method — modifies the original object
func (p *Person) SameButDifferent(u Person) {
	p.Name = u.Name
	p.Age = u.Age
	p.Id = u.Id
	p.Vehicle = u.Vehicle
}

// ✅ Return vehicle name (value receiver)
func (p Person) VehicleName(car string) string {
	return car
}

// ✅ Another method variation
func (p Person) VehicleButDifferent(car string) string {
	return "This is a " + car
}

func main() {
	// ✅ Create a Person instance correctly
	p := Person{
		Name:    "Ankush",
		Age:     27,
		Id:      130,
		Vehicle: "Car",
	}

	fmt.Println("Original Person:", p)

	// ✅ Use Information method
	updated := p.Information(Person{
		Name:    "Raj",
		Age:     30,
		Id:      200,
		Vehicle: "Bike",
	})
	fmt.Println("Updated (Information):", updated)

	// ✅ Use pointer receiver method
	p.SameButDifferent(Person{
		Name:    "Karan",
		Age:     28,
		Id:      300,
		Vehicle: "Scooter",
	})
	fmt.Println("After SameButDifferent:", p)

	// ✅ Call VehicleName method
	veh := p.VehicleName("Bus")
	fmt.Println("VehicleName:", veh)

	// ✅ Call VehicleButDifferent
	veh2 := p.VehicleButDifferent("Truck")
	fmt.Println("VehicleButDifferent:", veh2)
}
