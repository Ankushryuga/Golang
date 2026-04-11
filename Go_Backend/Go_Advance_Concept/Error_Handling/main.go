package errorhandling

import (
	"errors"
	"fmt"
	"log"
)

type Truck struct {
	id string
}

var (
	ErrorNotImplemented = errors.New("not implemented")
	ErrorTruckNotFound  = errors.New("Truck Not found")
)

func (t *Truck) LoadCargo() error {
	return ErrorTruckNotFound
}

// processTruck handles the loading and unloading of a truck.
func processTruck(truck Truck) error {
	fmt.Printf("Processing truck: %s\n", truck.id)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %w", err)
	}

	return ErrorNotImplemented
}

func main() {
	trucks := []Truck{
		{id: "Truck-1"},
		{id: "Truck-2"},
		{id: "Truck-3"},
	}

	for _, truck := range trucks {
		fmt.Printf("Truck %s arrived.\n", truck.id)

		if err := processTruck(truck); err != nil {
			if errors.Is(err, ErrorNotImplemented) {
				// if not implmented we do this
			}

			if errors.Is(err, ErrorTruckNotFound) {
				// if truck not found we do this
			}

			// another approach is it use switch.
			switch err {
			case ErrorNotImplemented:
				return
			case ErrorTruckNotFound:
				return
			default:
				log.Fatalf("Error not found")
			}

			log.Fatalf("Error processing truck: %s", err)
		}
	}

}
