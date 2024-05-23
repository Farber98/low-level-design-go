package main

import (
	"fmt"

	"github.com/Farber98/low-level-design/parking/lot"
	"github.com/Farber98/low-level-design/parking/parking"
	"github.com/Farber98/low-level-design/parking/vehicle"
)

func main() {

	// Create parking lots
	smallLots := []*lot.Lot{
		lot.NewLot(lot.SMALL_SIZE),
		lot.NewLot(lot.SMALL_SIZE),
		lot.NewLot(lot.SMALL_SIZE),
	}

	mediumLots := []*lot.Lot{
		lot.NewLot(lot.MEDIUM_SIZE),
		lot.NewLot(lot.MEDIUM_SIZE),
	}

	bigLots := []*lot.Lot{
		lot.NewLot(lot.BIG_SIZE),
	}

	// Create vehicles
	smallVehicle1 := vehicle.NewVehicle("small1", vehicle.SMALL_SIZE)
	smallVehicle2 := vehicle.NewVehicle("small2", vehicle.SMALL_SIZE)
	smallVehicle3 := vehicle.NewVehicle("small3", vehicle.SMALL_SIZE)
	smallVehicle4 := vehicle.NewVehicle("small4", vehicle.SMALL_SIZE)
	mediumVehicle1 := vehicle.NewVehicle("medium1", vehicle.MEDIUM_SIZE)
	mediumVehicle2 := vehicle.NewVehicle("medium2", vehicle.MEDIUM_SIZE)
	mediumVehicle3 := vehicle.NewVehicle("medium3", vehicle.MEDIUM_SIZE)
	bigVehicle1 := vehicle.NewVehicle("big1", vehicle.BIG_SIZE)
	bigVehicle2 := vehicle.NewVehicle("big2", vehicle.BIG_SIZE)

	parking := parking.NewParking(smallLots, mediumLots, bigLots)

	if parking.Park(smallVehicle1) {
		fmt.Println("PARKED: ", smallVehicle1.GetID())
	}

	if parking.Park(smallVehicle2) {
		fmt.Println("PARKED: ", smallVehicle2.GetID())
	}

	if parking.Park(smallVehicle3) {
		fmt.Println("PARKED: ", smallVehicle3.GetID())
	}

	if parking.Park(smallVehicle4) {
		fmt.Println("PARKED: ", smallVehicle4.GetID())
	}

	if parking.Unpark(smallVehicle4) {
		fmt.Println("UNPARKED: ", smallVehicle4.GetID())
	}

	if parking.Unpark(smallVehicle1) {
		fmt.Println("UNPARKED: ", smallVehicle1.GetID())
	}

	if parking.Park(smallVehicle4) {
		fmt.Println("PARKED: ", smallVehicle4.GetID())
	}

	if parking.Park(mediumVehicle1) {
		fmt.Println("PARKED: ", mediumVehicle1.GetID())
	}

	if parking.Park(mediumVehicle2) {
		fmt.Println("PARKED: ", mediumVehicle2.GetID())
	}

	if parking.Park(mediumVehicle3) {
		fmt.Println("PARKED: ", mediumVehicle3.GetID())
	}

	if parking.Unpark(mediumVehicle1) {
		fmt.Println("UNPARKED: ", mediumVehicle1.GetID())
	}

	if parking.Unpark(mediumVehicle3) {
		fmt.Println("UNPARKED: ", mediumVehicle3.GetID())
	}

	if parking.Park(mediumVehicle3) {
		fmt.Println("PARKED: ", mediumVehicle3.GetID())
	}

	if parking.Park(bigVehicle1) {
		fmt.Println("PARKED: ", bigVehicle1.GetID())
	}

	if parking.Park(bigVehicle2) {
		fmt.Println("PARKED: ", bigVehicle2.GetID())
	}

	if parking.Unpark(bigVehicle1) {
		fmt.Println("UNPARKED: ", bigVehicle1.GetID())
	}

	if parking.Park(bigVehicle2) {
		fmt.Println("PARKED: ", bigVehicle2.GetID())
	}
}
