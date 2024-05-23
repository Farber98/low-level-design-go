package vehicle

const SMALL_SIZE = 1
const MEDIUM_SIZE = 2
const BIG_SIZE = 3

// Vehicle Interface: GetID(), GetSize()
type IVehicle interface {
	GetID() string
	GetSize() int
}

// We'll have a Vehicle component that mainly identifies a car size and id.
type Vehicle struct {
	id   string
	size int
}

// Constructor.
func NewVehicle(id string, size int) *Vehicle {
	return &Vehicle{
		id,
		size,
	}
}

func (v *Vehicle) GetID() string {
	return v.id
}
func (v *Vehicle) GetSize() int {
	return v.size
}
