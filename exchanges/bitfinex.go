package exchanges

import (
	bitfinex "github.com/bitfinexcom/bitfinex-api-go/v2/rest"
	acc "github.com/guard-trader/account"
	op "github.com/guard-trader/operation"
)

// Bitfinex struct bot
type Bitfinex struct {
	name    string
	accName string
	client  *bitfinex.Client
	orders  []op.Operation
}

// CreateExchange create new client connection
func (b *Bitfinex) CreateExchange(account acc.Account) {
	b.client = bitfinex.NewClient()
	b.name = "bitfinex"
	b.orders = []op.Operation{}
}

// GetClient returns client of api
func (b *Bitfinex) GetClient() *bitfinex.Client {
	return b.client
}

func (b *Bitfinex) GetName() string {
	return b.name
}

func (b *Bitfinex) GetAccountName(accName string) string {
	return b.accName
}

func (b *Bitfinex) GetOrdersBySymbol(symbol string) []op.Operation {
	opr := []op.Operation{}
	return opr
}

func (b *Bitfinex) CreateOrder(symbol, price, qtd string) {

}

func (b *Bitfinex) GetOrder(symbol string, orderID int64) op.Operation {
	opr := op.Operation{}
	return opr
}

func (b *Bitfinex) CancelOrder(symbol string, orderID int64) {

}

func (b *Bitfinex) ListOpenOrders(symbol string) []op.Operation {
	opr := []op.Operation{}
	return opr
}

func (b *Bitfinex) ListTickerPrices() {

}

func (b *Bitfinex) ShowDepth(symbol string) {

}

func (b *Bitfinex) ListKlines(symbol, time string) {

}

func (b *Bitfinex) GetAccount() {

}
