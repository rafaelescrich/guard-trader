package main

import (
	binance "github.com/adshao/go-binance"
)

func main() {

	var (
		apiKey    = "WkB0gJY4xpVxT2OoShZUMnBANybsJyolz1B4XZxIda7OWKE0EUUP8zH9cwnnmH6f"
		secretKey = "nvtY4lwzDZuXyta2q04JqLZZsACh0jsfKnUo5tIS96kX8gM6B22KH7BiSbnLU03L"
	)
	client := binance.NewClient(apiKey, secretKey)

	allOrders(client, "GTOBTC")

}
