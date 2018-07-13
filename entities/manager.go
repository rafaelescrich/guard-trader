package entities

import (
	"fmt"

	"github.com/guard-trader/pairs"

	"github.com/boltdb/bolt"
)

// Manager of trades
type Manager struct {
	accounts []Account
	clients  []Exchange
	db       *bolt.DB
}

// NewManager creates new manager
func NewManager() *Manager {
	manager := Manager{}
	return &manager
}

// WakeUp and work
func (m *Manager) WakeUp() {

	m.NewClient("binance")

	m.clients[0].CreateOrder(pairs.BTCUSD, "5.853,8", "0,0001")

}

// NewClient create new exchange
func (m *Manager) NewClient(accName string) {
	account, _ := m.GetAcc(accName)

	switch account.ExchangeType {
	case BINANCE:
		binance := &Binance{}
		binance.CreateExchange(account)
		m.clients = append(m.clients, binance)

	case BITFINEX:
		bitfinex := &Bitfinex{}
		bitfinex.CreateExchange(account)
		m.clients = append(m.clients, bitfinex)
	}
}

// GetExchangesByName search and returns exchanges by name
func (m *Manager) GetExchangesByName(name string) []Exchange {
	exchanges := []Exchange{}
	for _, exchange := range m.clients {
		if exchange.GetName() == name {
			exchanges = append(exchanges, exchange)
		}
	}
	return exchanges
}

// GetExchangesByAcc search and returns exchanges by name
func (m *Manager) GetExchangesByAcc(name string) []Exchange {
	exchanges := []Exchange{}
	for _, exchange := range m.clients {
		if exchange.GetName() == name {
			exchanges = append(exchanges, exchange)
		}
	}
	return exchanges
}

// GetAccounts get all accounts for a manager
func (m *Manager) GetAccounts() []Account {
	return m.accounts
}

// GetAcc search for account with same name
func (m *Manager) GetAcc(accName string) (Account, error) {
	for _, account := range m.accounts {
		if account.Name == accName {
			return account, nil
		}
	}
	return Account{}, fmt.Errorf("account not exists")
}
