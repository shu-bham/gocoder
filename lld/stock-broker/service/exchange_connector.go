package service

import "gocoder/lld/stock-broker/entities"

type IExchangeConnector interface {
	GetRealTimePrice(symbol string) entities.LiveFeed
	SendOrderToExchange(orderId entities.Order) (bool, error)
}
