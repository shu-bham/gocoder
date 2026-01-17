package service

import "gocoder/lld/stock-broker/entities"

type OrderValidator interface {
	Validate(userId int, order entities.Order) (bool, error)
}
