package service

import "gocoder/lld/stock-broker/entities"

type StockBroker struct {
	Validator         IOrderValidator
	ExchangeConnector IExchangeConnector
	ExecutionService  IExecutionService
	WalletManager     IWalletManager
	PortfolioManager  IPortfolioManager
}

type IStockBrokerService interface {
	PlaceOrder(userId int, order entities.Order) bool
	GetPortfolio(userId int) entities.UserPortfolio
	AddFunds(userId int, amount float64) bool
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

}
