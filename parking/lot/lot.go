package lot

const SMALL_SIZE = 1
const MEDIUM_SIZE = 2
const BIG_SIZE = 3

var internal_id int = 0

type ILot interface {
	GetID() int
	GetSize() int
	GetAvailability() bool
	SetAvailable()
	SetUnavailable()
}

// We'll have a Lot component that mainly identifies a lot size, id and availability.
type Lot struct {
	id, size  int
	available bool
}

// Constructor.
func NewLot(size int) *Lot {
	internal_id++
	return &Lot{
		internal_id,
		size,
		true,
	}
}

func (l *Lot) GetID() int {
	return l.id
}
func (l *Lot) GetSize() int {
	return l.size
}

func (l *Lot) GetAvailability() bool {
	return l.available
}

func (l *Lot) SetAvailable() {
	if !l.available {
		l.available = true
	}
}

func (l *Lot) SetUnavailable() {
	if l.available {
		l.available = false
	}
}
