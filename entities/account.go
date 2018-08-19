package entities

type ClientType int

const (
	BINANCE = iota
	BITFINEX
)

// Account struct represents accounts of exchanges
type Account struct {
	ID        string
	name      string
	exchange  int
	apiKey    string
	secretKey string
}

// NewAccount acc
func NewAccount() *Account {
	acc := Account{}
	return &acc
}

// GetAPIKey api key
func (acc Account) GetAPIKey() string {
	return acc.apiKey
}

// GetAPISecretKey secret key
func (acc Account) GetAPISecretKey() string {
	return acc.secretKey
}

// GetName name
func (acc Account) GetName() string {
	return acc.name
}

// ClientType api exchange
func (acc *Account) ClientType(exchange int) {
	acc.exchange = exchange
}
