package cmd

import (
	"context"

	"github.com/abiosoft/ishell"
	m "github.com/guard-trader/manager"
)

var shell *ishell.Shell

// Run the shell
func Run(guard *m.Manager, cancel context.CancelFunc) {
	shell = ishell.New()
	addCmd(shell, guard)
	go shell.Run()
}

func addCmd(shell *ishell.Shell, guard *m.Manager) {
	shell.AddCmd(&ishell.Cmd{
		Name: "accs",
		Help: "Get accounts",
		Func: getAccounts(guard),
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "register",
		Help: "Register an account. 1-Name 2-ClientType",
		Func: registerAcc(guard),
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "service",
		Help: "Information of service of exchange.1-Exchange Name",
		Func: infoService(guard),
	})
}
