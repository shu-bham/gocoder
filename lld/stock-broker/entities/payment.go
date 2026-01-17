package entities

type PaymentMethod uint

const (
	PaymentMethodCreditCard PaymentMethod = iota
	PaymentMethodUPI
	PaymentMethodNetbanking
)

type PaymentService interface {
	AddFund(userId int, amount float64) bool
}
