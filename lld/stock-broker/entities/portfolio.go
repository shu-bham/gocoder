package entities

import "time"

type UserAccount struct {
	Id     int
	User   User
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

type OrderType uint

const (
	OrderTypeBuy OrderType = iota
	OrderTypeSell
)

type UserFund struct {
	user   User
	amount float64
}
