package entities

import (
	"context"
	"fmt"

	binance "github.com/adshao/go-binance"
)

// Binance struct bot
type Binance struct {
	name    string
	accName string
	client  *binance.Client
	orders  []Operation
}

// CreateExchange create new client connection
func (b *Binance) CreateExchange(account Account) {
	b.client = binance.NewClient(account.GetApiKey(), account.GetSecretKey())
	b.name = "binance"
	b.orders = []Operation{}
}

func (b *Binance) GetName() string {
	return b.name
}

func (b *Binance) GetAccountName(accName string) string {
	return b.accName
}

func (b *Binance) GetOrdersBySymbol(symbol string) []Operation {
	orders, err := b.client.NewListOrdersService().Symbol(symbol).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return []Operation{}
	}
	fmt.Println(orders)
	// TODO: operation <= order
	opr := []Operation{}
	return opr
}

func (b *Binance) CreateOrder(symbol, price, qtd string) {
	order, err := b.client.NewCreateOrderService().Symbol(symbol).
		Side(binance.SideTypeBuy).Type(binance.OrderTypeLimit).
		TimeInForce(binance.TimeInForceGTC).Quantity(qtd).
		Price(price).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(order)
}

func (b *Binance) GetOrder(symbol string, orderID int64) Operation {
	order, err := b.client.NewGetOrderService().Symbol(symbol).
		OrderID(orderID).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return Operation{}
	}
	fmt.Println(order)
	// TODO: operation <= order
	opr := Operation{}
	return opr
}

func (b *Binance) CancelOrder(symbol string, orderID int64) {
	_, err := b.client.NewCancelOrderService().Symbol(symbol).
		OrderID(orderID).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (b *Binance) ListOpenOrders(symbol string) []Operation {
	openOrders, err := b.client.NewListOpenOrdersService().Symbol(symbol).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return []Operation{}
	}
	for _, o := range openOrders {
		fmt.Println(o)
	}
	// TODO: operation <= order
	opr := []Operation{}
	return opr
}

func (b *Binance) ListTickerPrices() {
	prices, err := b.client.NewListPricesService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, p := range prices {
		fmt.Println(p)
	}

}

func (b *Binance) ShowDepth(symbol string) {
	res, err := b.client.NewDepthService().Symbol("LTCBTC").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

}

func (b *Binance) ListKlines(symbol, time string) {
	klines, err := b.client.NewKlinesService().Symbol(symbol).
		Interval(time).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, k := range klines {
		fmt.Println(k)
	}
}

func (b *Binance) GetAccount() {
	res, err := b.client.NewGetAccountService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

/* func (b *Binance) GetOrdersBySymbol(symbol string) []op.Operation {
	ops := []op.Operation{}
	for _, op := range b.orders {
		if op.GetSymbol() == symbol {
			ops = append(ops, op)
		}
	}
	return ops
} */

/* // List Aggregate Trades
func (b *Binance) listAggregateTrades(symbol string, start, end int64) {
	trades, err := client.NewAggTradesService().
		Symbol(symbol).StartTime(start).EndTime(end).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, t := range trades {
		fmt.Println(t)
	}
} */

/* // Start User Stream
func (b *Binance) startUserStream() {
	res, err := client.NewStartUserStreamService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
*/
