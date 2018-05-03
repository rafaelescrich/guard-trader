package core

import binance "github.com/adshao/go-binance"

// Guard struct bot
type Guard struct {
	client *binance.Client
}

// NewGuard create new guard trader
func NewGuard(apiKey, secretKey string) (*Guard, error) {
	guard := Guard{
		client: binance.NewClient(apiKey, secretKey),
	}
	return &guard, nil
}

// WakeUp init guard process
func (g Guard) WakeUp() {
	// TODO process
}
