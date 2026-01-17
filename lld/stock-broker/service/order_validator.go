package service

import "gocoder/lld/stock-broker/entities"

type IOrderValidator interface {
	Validate(userId int, order entities.Order) (bool, error)
}
