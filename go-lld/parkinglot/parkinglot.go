package main

import (
	"errors"
	"fmt"
)

type Vehicle interface {
	GetNumber() string
	GetType() string
}

type Car struct {
	Number string
}

// receiver:
func (c Car) GetNumber() string {
	return c.Number
}

func (c Car) GetType() string {
	return "Car"
}

// Bike:
type Bike struct {
	Number string
}

func (b Bike) GetNumber() string {
	return b.Number
}

func (b Bike) GetType() string {
	return "Bike"
}

// slot represents: parking space.
type ParkingSlot struct {
	SlotNumber int
	Occupied   bool
	Vehicle    Vehicle
}

// parking lot represent entire parking system.
type ParkingLot struct {
	Slots []ParkingSlot
}

func NewParkingLot(capacity int) *ParkingLot {
	slots := make([]ParkingSlot, capacity)
	for i := 0; i < capacity; i++ {
		slots[i] = ParkingSlot{SlotNumber: i + 1}
	}
	return &ParkingLot{Slots: slots}
}

//Park vehicle parks a vehicle and return assigned slot.

func (p *ParkingLot) ParkVehicle(v Vehicle) (int, error) {
	for i := range p.Slots {
		if !p.Slots[i].Occupied {
			p.Slots[i].Vehicle = v
			p.Slots[i].Occupied = true
			return p.Slots[i].SlotNumber, nil
		}
	}
	return 0, errors.New("no available slot")
}

func (p *ParkingLot) RemoveVehicle(slotNumber int) error {
	if slotNumber < 1 || slotNumber > len(p.Slots) {
		return errors.New("invalid slot number")
	}
	if !p.Slots[slotNumber-1].Occupied {
		return errors.New("slot already empty")
	}
	p.Slots[slotNumber-1].Vehicle = nil
	p.Slots[slotNumber-1].Occupied = false
	return nil
}

func (p *ParkingLot) DisplayStatus() {
	fmt.Println("Slot\tVehicle Type \tVehicle Number")
	for _, slot := range p.Slots {
		if slot.Occupied {
			fmt.Printf("%d\t%s\t%s\n", slot.SlotNumber, slot.Vehicle.GetType(), slot.Vehicle.GetNumber())
		}
	}
}

func main() {
	lot := NewParkingLot(5)
	car := Car{Number: "KA-3939-3HH-4923"}
	bike := Bike{Number: "BH-2229-BH-1234"}

	slot1, _ := lot.ParkVehicle(car)
	fmt.Printf("Parked car at slot %d\n", slot1)

	slot2, _ := lot.ParkVehicle(bike)
	fmt.Printf("Parked bike at slot %d\n", slot2)

	lot.DisplayStatus()

	_ = lot.RemoveVehicle(slot1)
	fmt.Println("Removed vehicle from slot ", slot1)

	lot.DisplayStatus()
}
