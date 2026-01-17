package service

import "gocoder/lld/stock-broker/entities"

type IPortfolioManager interface {
	UpdatePortfolio(userId int, instrumentId int, qty int, price float64, action entities.OrderKind)
}
