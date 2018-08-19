package entities

// Exchange client
type Exchange interface {
	AccessExchange(account Account)
	GetName() string
	GetAccountName(accName string) string
	CreateOrder(symbol, price, qtd string)
	GetOrder(symbol string, orderID int64) Operation
	GetOrdersBySymbol(symbol string) []Operation
	ListOpenOrders(symbol string) []Operation
	CancelOrder(symbol string, orderID int64)
	ShowDepth(symbol string)
	ListTickerPrices()
	ListKlines(symbol, time string)
	GetAccount()
	TimeService()
	GetNewHistoryTrades(symbol string)
	GetInfoService()
}
