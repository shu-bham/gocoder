package service

import "gocoder/lld/stock-broker/entities"

type ExchangeConnector interface {
	GetRealTimePrice(symbol string) entities.LiveFeed
	SendOrderToExchange(order entities.Order) (bool, error)
}
