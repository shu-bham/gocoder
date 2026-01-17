package service

import (
	"gocoder/lld/stock-broker/entities"
	"sync"
)

type portfolioManager struct {
	portfolios map[int]*entities.UserPortfolio
	mu         sync.RWMutex
}

func NewPortfolioManager() *portfolioManager {
	return &portfolioManager{
		portfolios: make(map[int]*entities.UserPortfolio),
	}
}

func (p *portfolioManager) UpdatePortfolio(userId int, instrumentId int, qty int, price float64, action entities.OrderKind) {
	p.mu.Lock()
	defer p.mu.Unlock()

	portfolio, exists := p.portfolios[userId]
	if !exists {
		portfolio = &entities.UserPortfolio{UserId: userId, Holdings: []entities.Holdings{}}
		p.portfolios[userId] = portfolio
	}

	// Logic to find holding and update slice
	found := false
	for i, h := range portfolio.Holdings {
		if h.InstrumentID == instrumentId {
			if action == entities.OrderTypeBuy {
				// Weighted Average Price
				totalValue := (float64(h.TotalQty) * h.AvgPrice) + (float64(qty) * price)
				portfolio.Holdings[i].TotalQty += uint(qty)
				portfolio.Holdings[i].AvgPrice = totalValue / float64(portfolio.Holdings[i].TotalQty)
			} else {
				portfolio.Holdings[i].TotalQty -= uint(qty)
			}
			found = true
			break
		}
	}

	if !found && action == entities.OrderTypeBuy {
		portfolio.Holdings = append(portfolio.Holdings, entities.Holdings{
			InstrumentID: instrumentId,
			TotalQty:     uint(qty),
			AvgPrice:     price,
		})
	}
}
