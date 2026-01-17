package service

import (
	"gocoder/lld/stock-broker/entities"
	"sync"
)

type walletManager struct {
	funds map[int]*entities.UserFund
	mu    sync.Mutex
}

func NewWalletManager() *walletManager {
	return &walletManager{
		funds: make(map[int]*entities.UserFund),
	}
}

func (w *walletManager) AddFunds(userId int, amount float64) bool {
	w.mu.Lock()
	defer w.mu.Unlock()

	fund, exists := w.funds[userId]
	if !exists {
		// In real app, maybe error or auto-create
		fund = &entities.UserFund{UserId: userId, Balance: 0}
		w.funds[userId] = fund
	}
	fund.Balance += amount
	return true
}

func (w *walletManager) DeductMoney(userId int, amount float64) bool {
	w.mu.Lock()
	defer w.mu.Unlock()

	fund, exists := w.funds[userId]
	if !exists || fund.Balance < amount {
		return false
	}
	fund.Balance -= amount
	return true
}
