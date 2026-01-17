package service

import (
	"gocoder/lld/stock-broker/entities"
	"sync"
)

// Observer Interface
type PriceObserver interface {
	OnPriceUpdate(feed entities.LiveFeed)
}

// Subject Interface
type MarketDataSubject interface {
	Subscribe(o PriceObserver)
	Unsubscribe(o PriceObserver)
	Notify(feed entities.LiveFeed)
}

// Implementation of Exchange Connector as Singleton + Subject
type StockExchange struct {
	observers []PriceObserver
	mu        sync.Mutex
}

var instance *StockExchange
var once sync.Once

// GetExchangeInstance returns the Singleton instance
func GetExchangeInstance() *StockExchange {
	once.Do(func() {
		instance = &StockExchange{
			observers: make([]PriceObserver, 0),
		}
	})
	return instance
}

// Implementation of IExchangeConnector methods
func (e *StockExchange) GetRealTimePrice(symbol string) entities.LiveFeed {
	// Mock implementation
	return entities.LiveFeed{InstrumentId: 1, LastTradedPrice: 150.0}
}

func (e *StockExchange) SendOrderToExchange(order entities.Order) (bool, error) {
	// Mock execution
	return true, nil
}

// Observer Pattern Methods
func (e *StockExchange) Subscribe(o PriceObserver) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.observers = append(e.observers, o)
}

func (e *StockExchange) Unsubscribe(o PriceObserver) {
	e.mu.Lock()
	defer e.mu.Unlock()
	// logic to remove observer
}

func (e *StockExchange) Notify(feed entities.LiveFeed) {
	e.mu.Lock()
	defer e.mu.Unlock()
	for _, observer := range e.observers {
		go observer.OnPriceUpdate(feed) // Notify asynchronously
	}
}

// Simulation method to trigger an update
func (e *StockExchange) SimulateMarketUpdate(feed entities.LiveFeed) {
	e.Notify(feed)
}
