package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrorNotImplemented = errors.New("not implemented")
	ErrorTruckNotFound  = errors.New("truck not found")
)

// interface::
type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}
type NormalTruck struct {
	id    string
	cargo int
}

func (t *NormalTruck) LoadCargo() error {
	return nil
}

func (t *NormalTruck) UnloadCargo() error {
	return nil
}

type ElectricTruct struct {
	id    string
	cargo int
}

func processTruck(truck Truck) error {
	err := truck.LoadCargo()
	if err != nil {
		return fmt.Errorf("Error loading cargo: %w", err)
	}

	err = truck.UnloadCargo()
	if err != nil {
		return fmt.Errorf("Error unloading cargo: %w", err)
	}

}

func main() {
	fmt.Println("Interface example for backend")

	trucks := []*NormalTruck{
		{id: "Truck - 1"},
		{id: "Truck - 2"},
		{id: "Truck - 3"},
	}

	eTrucks := []ElectricTruct{
		{id: "Electrick-truck-1"},
	}

	for _, truck := range trucks {
		fmt.Printf("Trucks %s arrived. \n", truck.id)

		err := processTruck(truck)
		if err != nil {
			log.Fatalf("Error processing truck: %s", err)
		}

		fmt.Printf("Truck %s departed. \n", truck.id)
	}
}
