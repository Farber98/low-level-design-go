package parking

import (
	"github.com/Farber98/low-level-design/parking/lot"
	"github.com/Farber98/low-level-design/parking/vehicle"
)

// We'll have a Parking component that will contain an amount of lots of different sizes
// Through this component we'll be able to register and remove cars from lots, given their size and lot availability.
// Parking Interface: Park(Vehicle), Unpark(Vehicle)
// Constructor(): Receives lots amount of each type and initializes the map.
// Internal Logic:
// Given a car, check if we having available lots of that size
// If we have, place it, update the availability of that lot and update the map of cars in lots
// If we don't have, return err

type IParking interface {
	Park(vehicle *vehicle.Vehicle) bool
	Unpark(vehicle *vehicle.Vehicle) bool
}

type Parking struct {
	smallLots  []*lot.Lot
	mediumLots []*lot.Lot
	bigLots    []*lot.Lot
	smallLotsToVehiclesMp,
	mediumLotsToVehiclesMp,
	bigLotsToVehiclesMp map[*lot.Lot]*vehicle.Vehicle
	vehiclesToLots map[*vehicle.Vehicle]*lot.Lot
}

// Constructor.
func NewParking(smallLots, mediumLots, bigLots []*lot.Lot) IParking {
	return &Parking{
		smallLots:              smallLots,
		mediumLots:             mediumLots,
		bigLots:                bigLots,
		smallLotsToVehiclesMp:  make(map[*lot.Lot]*vehicle.Vehicle, len(smallLots)),
		mediumLotsToVehiclesMp: make(map[*lot.Lot]*vehicle.Vehicle, len(mediumLots)),
		bigLotsToVehiclesMp:    make(map[*lot.Lot]*vehicle.Vehicle, len(bigLots)),
		vehiclesToLots:         make(map[*vehicle.Vehicle]*lot.Lot, len(smallLots)+len(mediumLots)+len(bigLots)),
	}
}

func (p *Parking) Park(v *vehicle.Vehicle) bool {
	// Try to get available lot.
	lot := p.getAvailableLot(v.GetSize())
	if lot == nil {
		return false
	}

	// Establish mapping from lot to vehicle and vehicle to lot for quick lookup.
	p.mapLotToVehicleAndVehicleToLot(lot, v)

	// Remove the lot from the available lots slice
	p.removeLotFromAvailableLots(v.GetSize())

	// Car was parked
	return true

}

func (p *Parking) getAvailableLot(size int) *lot.Lot {
	// Get the slice of available lots for the corresponding size
	var availableLots []*lot.Lot
	switch size {
	case vehicle.SMALL_SIZE:
		availableLots = p.smallLots
	case vehicle.MEDIUM_SIZE:
		availableLots = p.mediumLots
	case vehicle.BIG_SIZE:
		availableLots = p.bigLots
	}

	// Check if there are available lots of the given size
	if len(availableLots) == 0 {
		return nil
	}

	// Return last one
	return availableLots[len(availableLots)-1]
}

func (p *Parking) mapLotToVehicleAndVehicleToLot(lot *lot.Lot, v *vehicle.Vehicle) {
	switch v.GetSize() {
	case vehicle.SMALL_SIZE:
		p.smallLotsToVehiclesMp[lot] = v
	case vehicle.MEDIUM_SIZE:
		p.mediumLotsToVehiclesMp[lot] = v
	case vehicle.BIG_SIZE:
		p.bigLotsToVehiclesMp[lot] = v
	}

	p.vehiclesToLots[v] = lot
}

func (p *Parking) removeLotFromAvailableLots(size int) {
	// Return last one for complexity purposes avoiding shifts.
	switch size {
	case vehicle.SMALL_SIZE:
		p.smallLots = p.smallLots[:len(p.smallLots)-1]
	case vehicle.MEDIUM_SIZE:
		p.mediumLots = p.mediumLots[:len(p.mediumLots)-1]
	case vehicle.BIG_SIZE:
		p.bigLots = p.bigLots[:len(p.bigLots)-1]
	}
}

func (p *Parking) Unpark(v *vehicle.Vehicle) bool {
	// Just need to unpark that vehicle

	// Get lot of that vehicle from map
	lot, ok := p.vehiclesToLots[v]
	if !ok {
		return false
	}

	// Unmap lot to vehicle
	p.unmapLotToVehicleAndVehicleToLotAndPutBackAvailableLot(lot, v)

	return true
}

func (p *Parking) unmapLotToVehicleAndVehicleToLotAndPutBackAvailableLot(lot *lot.Lot, v *vehicle.Vehicle) {
	switch v.GetSize() {
	case vehicle.SMALL_SIZE:
		p.smallLotsToVehiclesMp[lot] = nil
		p.smallLots = append(p.smallLots, lot)
	case vehicle.MEDIUM_SIZE:
		p.mediumLotsToVehiclesMp[lot] = nil
		p.mediumLots = append(p.mediumLots, lot)
	case vehicle.BIG_SIZE:
		p.bigLotsToVehiclesMp[lot] = nil
		p.bigLots = append(p.bigLots, lot)
	}

	// unmap vehicle to lot
	delete(p.vehiclesToLots, v)
}
