package cmd

import (
	"fmt"
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/guard-trader/manager"
)

func getAccounts(guard *manager.Manager) func(*ishell.Context) {
	return func(c *ishell.Context) {
		if len(c.Args) != 0 {
			return
		}
		accs := guard.GetAccounts()
		for _, acc := range accs {
			fmt.Println("Exchange: ", acc.GetName())
		}
	}

}

func registerAcc(guard *manager.Manager) func(*ishell.Context) {
	return func(c *ishell.Context) {
		if len(c.Args) != 2 {
			fmt.Println("Name and client type of account is required")
			return
		}
		accName := c.Args[0]
		client, err := strconv.Atoi(c.Args[1])
		if err != nil {
			fmt.Println("Invalid client type")
			return
		}
		//account := ent.NewAccount()
		//guard.AddAccount(account)
		guard.NewClient(accName, client)
	}
}

func infoService(guard *manager.Manager) func(*ishell.Context) {
	return func(c *ishell.Context) {
		if len(c.Args) != 1 {
			fmt.Println("Name of account is required")
			return
		}
		exchange, err := guard.GetExchangeByName(c.Args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		exchange.GetInfoService()
	}
}
