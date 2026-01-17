package service

import "gocoder/lld/stock-broker/entities"

type ExecutionService interface {
	Execute(order entities.Order) bool
}
