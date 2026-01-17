package entities

import (
	"time"
)

type UserAccount struct {
	Id     int
	UserId int
	Status AccountStatus
}

type AccountStatus uint

const (
	AccountStatusActive AccountStatus = iota
	AccountStatusDisabled
)

type Order struct {
	Instrument     Instrument
	Price          float64
	Status         OrderStatus
	Kind           OrderKind
	Type           OrderType
	Units          uint
	OrderTimeStamp time.Time
}

type OrderStatus uint

const (
	OrderStatusActive OrderStatus = iota
	OrderStatusCancelled
	OrderStatusRejected
	OrderStatusCompleted
)

type OrderKind uint

const (
	OrderTypeBuy OrderKind = iota
	OrderTypeSell
)

type OrderType uint

const (
	OrderTypeLimit OrderType = iota
	OrderTypeMarket
)

type UserPortfolio struct {
	UserId   int
	Holdings []Holdings
}

type Holdings struct {
	InstrumentID int
	TotalQty     uint
	AvgPrice     float64
}

type OrderBook struct {
	InstrumentId int
	buyOrders    []Order
	sellOrders   []Order
}
