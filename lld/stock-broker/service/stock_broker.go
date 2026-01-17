package service

import "gocoder/lld/stock-broker/entities"

type StockBroker struct {
	Validator         OrderValidator
	ExchangeConnector ExchangeConnector
	ExecutionService  ExecutionService
	WalletManager     WalletManager
	PortfolioManager  PortfolioManager
}

type StockBrokerService interface {
	PlaceOrder(userId int, order entities.Order) bool
	GetPortfolio(userId int) entities.UserPortfolio
	AddFunds(userId int, amount float64) bool
}

func NewStockBroker(v OrderValidator, e ExchangeConnector, ex ExecutionService, w WalletManager, p PortfolioManager) *StockBroker {
	return &StockBroker{
		Validator:         v,
		ExchangeConnector: e,
		ExecutionService:  ex,
		WalletManager:     w,
		PortfolioManager:  p,
	}
}

func (s *StockBroker) PlaceOrder(userId int, order entities.Order) bool {
	_, err := s.Validator.Validate(userId, order)
	if err != nil {
		return false
	}
	totalCost := order.Price * float64(order.Units)
	if order.Kind == entities.OrderTypeBuy {
		s.WalletManager.DeductMoney(userId, totalCost)
	}
	done := s.ExecutionService.Execute(order)
	if done {
		s.PortfolioManager.UpdatePortfolio(userId, order.Instrument.ID, int(order.Units), order.Price, order.Kind)
	} else {
		if order.Kind == entities.OrderTypeBuy {
			s.WalletManager.AddFunds(userId, totalCost)
		}
		return false
	}
	return done
}

func (s *StockBroker) GetPortfolio(userId int) entities.UserPortfolio {

}

func (s *StockBroker) AddFunds(userId int, amount float64) bool {
	return s.WalletManager.AddFunds(userId, amount)
}

// OnPriceUpdate implements IPriceObserver
func (s *StockBroker) OnPriceUpdate(feed entities.LiveFeed) {
	// In a real app, check 'Open Orders' to see if any Limit Order matches this price
	// For now, we just acknowledge the update
	// fmt.Printf("Broker received price update for Instrument %d: %f\n", feed.InstrumentId, feed.LastTradedPrice)
}
