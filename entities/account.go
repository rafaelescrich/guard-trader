package entities

type ClientType int

const (
	BINANCE = iota
	BITFINEX
)

// Account struct represents accounts of exchanges
type Account struct {
	ID           string
	Name         string
	ExchangeType ClientType
	apiKey       string
	secretKey    string
}

func (acc *Account) GetApiKey() string {
	return acc.apiKey
}

func (acc *Account) GetSecretKey() string {
	return acc.secretKey
}
