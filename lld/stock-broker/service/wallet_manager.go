package service

type IWalletManager interface {
	AddFunds(userId int, amount float64) bool
	DeductMoney(userId int, amount float64) bool
}
