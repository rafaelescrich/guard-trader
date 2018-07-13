package entities

import (
	bitfinex "github.com/bitfinexcom/bitfinex-api-go/v2/rest"
)

// Bitfinex struct bot
type Bitfinex struct {
	name    string
	accName string
	client  *bitfinex.Client
	orders  []Operation
}

// CreateExchange create new client connection
func (b *Bitfinex) CreateExchange(account Account) {
	b.client = bitfinex.NewClient()
	b.name = "bitfinex"
	b.orders = []Operation{}
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

func (b *Bitfinex) GetOrdersBySymbol(symbol string) []Operation {
	opr := []Operation{}
	return opr
}

func (b *Bitfinex) CreateOrder(symbol, price, qtd string) {

}

func (b *Bitfinex) GetOrder(symbol string, orderID int64) Operation {
	opr := Operation{}
	return opr
}

func (b *Bitfinex) CancelOrder(symbol string, orderID int64) {

}

func (b *Bitfinex) ListOpenOrders(symbol string) []Operation {
	opr := []Operation{}
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
