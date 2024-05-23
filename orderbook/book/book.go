package book

import (
	"fmt"

	"github.com/Farber98/low-level-design/orderbook/order"
)

type Book struct {
	orderIdToOrder       map[int]*order.Order
	ask                  AskHeap
	askPriceToPriceLevel map[int]*PriceLevel
	askPriceToVolume     map[int]int
	bid                  BidHeap
	bidPriceToPriceLevel map[int]*PriceLevel
	bidPriceToVolume     map[int]int
}

type IBook interface {
	PlaceOrder(o *order.Order) bool
	RemoveOrder(orderId int) bool
	GetVolumeAtPrice(price, side int) int
}

func NewBook() *Book {
	return &Book{
		orderIdToOrder:       make(map[int]*order.Order, 0),
		ask:                  make(AskHeap, 0),
		askPriceToPriceLevel: make(map[int]*PriceLevel, 0),
		askPriceToVolume:     make(map[int]int, 0),
		bid:                  make(BidHeap, 0),
		bidPriceToPriceLevel: make(map[int]*PriceLevel, 0),
		bidPriceToVolume:     make(map[int]int, 0),
	}
}

func (b *Book) PlaceOrder(o *order.Order) bool {
	// Figure side of books
	sameBook, oppositeBook := b.figureSameAndOppositeBook(o.GetSide())

	// Map order to order id
	b.orderIdToOrder[o.GetID()] = o

	// While order has volume, oppositte book isn't empty and the next order to be filled has volume and it's a price we are willing to trade
	for o.GetVolume() > 0 &&
		oppositeBook.Len() > 0 &&
		oppositeBook.Peek() != nil &&
		oppositeBook.Peek().getNextOrder() != nil &&
		oppositeBook.Peek().getNextOrder().GetPrice() <= o.GetPrice() {

		// Grab the other order we are going to trade against
		otherOrder := oppositeBook.Peek().getNextOrder()
		// Set trade price and trade volume
		tradePrice, tradeVolume := otherOrder.GetPrice(), min(o.GetVolume(), otherOrder.GetVolume())
		// Print the trade
		fmt.Println("Made by %v, Taken by %v, %v shares @ %v price", otherOrder.GetClientID(), o.GetClientID(), tradeVolume, tradePrice)

		// adjuste both orders after the trade
		otherOrder.SetVolume(otherOrder.GetVolume() - tradeVolume)
		o.SetVolume(o.GetVolume() - tradeVolume)

		// If the other order got to 0 volume, remove it
		if otherOrder.GetVolume() == 0 {
			b.RemoveOrder(otherOrder.GetID())
		}

	}

	// If after the loop our order still has volume, place on book
	if o.GetVolume() > 0 {
		addToOrderBook(o)
	}

	return false
}

func (b *Book) figureSameAndOppositeBook(side int) (IHeap, IHeap) {
	var sameBook, oppositeBook IHeap
	switch side {
	case order.BUY_SIDE:
		sameBook, oppositeBook = &b.ask, &b.bid
	case order.SELL_SIDE:
		sameBook, oppositeBook = &b.bid, &b.ask
	}

	return sameBook, oppositeBook
}

func (b *Book) RemoveOrder(orderId int) bool {
	return false
}

/*
func (b *Book) GetVolumeAtPrice(price, side int) int {

} */

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
