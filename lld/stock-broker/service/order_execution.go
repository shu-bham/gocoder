package service

import "gocoder/lld/stock-broker/entities"

type IExecutionService interface {
	Execute(order entities.Order) bool
}
