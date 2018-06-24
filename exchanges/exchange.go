package exchanges

import (
	acc "github.com/guard-trader/account"
	op "github.com/guard-trader/operation"
)

// TODO: return of errors
// Exchange client
type Exchange interface {
	CreateExchange(account acc.Account)
	GetName() string
	GetAccountName(accName string) string
	CreateOrder(symbol, price, qtd string)
	GetOrder(symbol string, orderID int64) op.Operation
	GetOrdersBySymbol(symbol string) []op.Operation
	ListOpenOrders(symbol string) []op.Operation
	CancelOrder(symbol string, orderID int64)
	ShowDepth(symbol string)
	ListTickerPrices()
	ListKlines(symbol, time string)
	GetAccount()
	//GetOperations() []op.Operation
	//GetClient() *interface{}
}
