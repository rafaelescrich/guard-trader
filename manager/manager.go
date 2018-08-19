package manager

import (
	"fmt"

	ent "github.com/guard-trader/entities"
	rp "github.com/guard-trader/repository"
)

// Manager of trades
type Manager struct {
	accounts []ent.Account
	clients  []ent.Exchange
	//indicators []ent.Indicator
	//analysis   ent.Analysis
	Badger *rp.Badger
}

// NewManager creates new manager
func NewManager() *Manager {
	manager := Manager{}
	manager.Badger = rp.OpenBadger()
	return &manager
}

// WakeUp and work
func (m *Manager) WakeUp() {
	m.LoadAccounts()
	select {}
}

// NewClient create new exchange
func (m *Manager) NewClient(accName string, client int) {
	account, _ := m.GetAccByName(accName)
	account.ClientType(client)

	switch client {
	case ent.BINANCE:
		binance := &ent.Binance{}
		binance.AccessExchange(account)
		fmt.Println("Access: ", binance.GetName())
		m.clients = append(m.clients, binance)

	// Comented for tests binance cli first
	case ent.BITFINEX:
		bitfinex := &ent.Bitfinex{}
		bitfinex.AccessExchange(account)
		m.clients = append(m.clients, bitfinex)
	}
}

// GetExchangeByName search exchanges by name
func (m *Manager) GetExchangeByName(name string) (ent.Exchange, error) {
	for _, exchange := range m.clients {
		if exchange.GetName() == name {
			return exchange, nil
		}
	}
	return nil, fmt.Errorf("exchange not exist")
}

// GetExchangesByAcc search and returns exchanges by name
func (m *Manager) GetExchangesByAcc(name string) []ent.Exchange {
	exchanges := []ent.Exchange{}
	for _, exchange := range m.clients {
		if exchange.GetName() == name {
			exchanges = append(exchanges, exchange)
		}
	}
	return exchanges
}

// GetAccounts get all accounts for a manager
func (m *Manager) GetAccounts() []ent.Account {
	return m.accounts
}

// GetAccByName search for account with same name
func (m *Manager) GetAccByName(accName string) (ent.Account, error) {
	for _, account := range m.accounts {
		if account.GetName() == accName {
			return account, nil
		}
	}
	return ent.Account{}, fmt.Errorf("account not exists")
}

// LoadAccounts load accounts on manager
func (m *Manager) LoadAccounts() {
	accs, err := m.Badger.GetAccounts()
	if err != nil {
		return
	}
	m.accounts = accs
}

// AddAccount to manager
func (m *Manager) AddAccount(acc ent.Account) {
	m.accounts = append(m.accounts, acc)
}
