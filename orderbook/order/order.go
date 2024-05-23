package order

import "time"

var internal_id = 0

const BUY_SIDE = 0
const SELL_SIDE = 1

type Order struct {
	id        int
	clientId  string
	timestamp int
	side      int
	price     int
	volume    int
}

type IOrder interface {
	GetID() int
	GetClientID() string
	GetTimestamp() int
	GetSide() int
	GetPrice() int
	GetVolume() int
	SetVolume(v int)
}

func NewOrder(clientId string, side, price, volume int) *Order {
	internal_id++
	return &Order{
		id:        internal_id,
		clientId:  clientId,
		timestamp: int(time.Now().Unix()),
		side:      side,
		price:     price,
		volume:    volume,
	}
}

func (o *Order) GetID() int {
	return o.id
}

func (o *Order) GetClientID() string {
	return o.clientId
}

func (o *Order) GetTimestamp() int {
	return o.timestamp

}

func (o *Order) GetSide() int {
	return o.side

}

func (o *Order) GetPrice() int {
	return o.price

}

func (o *Order) GetVolume() int {
	return o.volume

}

func (o *Order) SetVolume(v int) {
	if v >= 0 {
		o.volume = v
	}
}
