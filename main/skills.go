package main

import (
	"context"
	"fmt"

	binance "github.com/adshao/go-binance"
)

// Create order
func createOrder(client *binance.Client, symbol, price, qtd string) {
	order, err := client.NewCreateOrderService().Symbol(symbol).
		Side(binance.SideTypeBuy).Type(binance.OrderTypeLimit).
		TimeInForce(binance.TimeInForceGTC).Quantity(qtd).
		Price(price).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(order)
}

// Get order
func getOrder(client *binance.Client, symbol string, orderID int64) {
	order, err := client.NewGetOrderService().Symbol(symbol).
		OrderID(orderID).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(order)
}

// Cancel order
func cancelOrder(client *binance.Client, symbol string, orderID int64) {
	_, err := client.NewCancelOrderService().Symbol(symbol).
		OrderID(orderID).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
}

// List open orders
func listOpenOrders(client *binance.Client, symbol string) {
	openOrders, err := client.NewListOpenOrdersService().Symbol(symbol).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, o := range openOrders {
		fmt.Println(o)
	}
}

//list all orders
func allOrders(client *binance.Client, symbol string) {
	orders, err := client.NewListOrdersService().Symbol(symbol).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, o := range orders {
		fmt.Println(o)
	}
}

// List Ticker Prices
func listTickerPrices(client *binance.Client) {
	prices, err := client.NewListPricesService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, p := range prices {
		fmt.Println(p)
	}

}

// Show Depth
func showDepth(client *binance.Client, symbol string) {
	res, err := client.NewDepthService().Symbol("LTCBTC").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

}

// List Klines
func listKlines(client *binance.Client, symbol, time string) {
	klines, err := client.NewKlinesService().Symbol(symbol).
		Interval(time).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, k := range klines {
		fmt.Println(k)
	}
}

// List Aggregate Trades
func listAggregateTrades(client *binance.Client, symbol string, start, end int64) {
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
}

// Get Account
func getAccount(client *binance.Client) {
	res, err := client.NewGetAccountService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

// Start User Stream
func startUserStream(client *binance.Client) {
	res, err := client.NewStartUserStreamService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
