package book

import (
	"github.com/Farber98/low-level-design/orderbook/order"
)

// An Item is a price level in the order book with a queue of orders at that price.
type PriceLevel struct {
	price  int            // The price of the item (priority in the heap).
	orders []*order.Order // Queue of orders at this price level.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

func NewPriceLevel(price int, o *order.Order) *PriceLevel {
	return &PriceLevel{
		price:  price,
		orders: []*order.Order{o},
	}
}

func (p *PriceLevel) getNextOrder() *order.Order {
	var o *order.Order

	if len(p.orders) > 0 {
		o = p.orders[0]
	}

	return o
}

// IHeap is an interface that abstracts the common operations for BidHeap and AskHeap.
type IHeap interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
	Push(x *PriceLevel)
	Pop() *PriceLevel
	update(item *PriceLevel, price int, orders []*order.Order)
	Peek() *PriceLevel
}

// A BidHeap implements heap.Interface and holds Items for the bid side (min heap).
type BidHeap []*PriceLevel

func (h BidHeap) Len() int           { return len(h) }
func (h BidHeap) Less(i, j int) bool { return h[i].price < h[j].price }
func (h BidHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i]; h[i].index = i; h[j].index = j }
func (h *BidHeap) Push(item *PriceLevel) {
	n := len(*h)
	item.index = n
	*h = append(*h, item)
}
func (h *BidHeap) Pop() *PriceLevel {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*h = old[0 : n-1]
	return item
}

// Peek returns the top element of the heap without removing it.
func (h BidHeap) Peek() *PriceLevel {
	if len(h) == 0 {
		return nil
	}
	return h[0]
}

// An AskHeap implements heap.Interface and holds Items for the ask side (max heap).
type AskHeap []*PriceLevel

func (h AskHeap) Len() int           { return len(h) }
func (h AskHeap) Less(i, j int) bool { return h[i].price > h[j].price }
func (h AskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i]; h[i].index = i; h[j].index = j }
func (h *AskHeap) Push(item *PriceLevel) {
	n := len(*h)
	item.index = n
	*h = append(*h, item)
}
func (h *AskHeap) Pop() *PriceLevel {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*h = old[0 : n-1]
	return item
}

// Peek returns the top element of the heap without removing it.
func (h AskHeap) Peek() *PriceLevel {
	if len(h) == 0 {
		return nil
	}
	return h[0]
}
