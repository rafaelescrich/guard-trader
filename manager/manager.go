package manager

import (
	"fmt"

	"github.com/boltdb/bolt"
	acc "github.com/guard-trader/account"
	exc "github.com/guard-trader/exchanges"
)

// Manager of trades
type Manager struct {
	accounts []acc.Account
	clients  []exc.Exchange
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
	//TODO: init exchanges
	// Get active orders from exchanges
}

// NewClient create new exchange
func (m *Manager) NewClient(accName string) {
	account, _ := m.GetAcc(accName)

	switch account.ExchangeType {
	case acc.BINANCE:
		binance := &exc.Binance{}
		binance.CreateExchange(account)
		m.clients = append(m.clients, binance)

	case acc.BITFINEX:
		bitfinex := &exc.Bitfinex{}
		bitfinex.CreateExchange(account)
		m.clients = append(m.clients, bitfinex)
	}
}

// GetExchangesByName search and returns exchanges by name
func (m *Manager) GetExchangesByName(name string) []exc.Exchange {
	exchanges := []exc.Exchange{}
	for _, exchange := range m.clients {
		if exchange.GetName() == name {
			exchanges = append(exchanges, exchange)
		}
	}
	return exchanges
}

// GetExchangesByAcc search and returns exchanges by name
func (m *Manager) GetExchangesByAcc(name string) []exc.Exchange {
	exchanges := []exc.Exchange{}
	for _, exchange := range m.clients {
		if exchange.GetName() == name {
			exchanges = append(exchanges, exchange)
		}
	}
	return exchanges
}

// GetAccounts get all accounts for a manager
func (m *Manager) GetAccounts() []acc.Account {
	return m.accounts
}

// GetAcc search for account with same name
func (m *Manager) GetAcc(accName string) (acc.Account, error) {
	for _, account := range m.accounts {
		if account.Name == accName {
			return account, nil
		}
	}
	return acc.Account{}, fmt.Errorf("account not exists")
}
