package cmd

import (
	"context"

	"github.com/abiosoft/ishell"
	"github.com/guard-trader/core"
)

var shell *ishell.Shell

// Run the shell
func Run(guard *core.Guard, cancel context.CancelFunc) {
	shell = ishell.New()
	//	monitoreInterrupt(node, cancel)
	addCmd(shell, guard)
	go shell.Run()
}

func addCmd(shell *ishell.Shell, guard *core.Guard) {
	//state := node.StateMgr
	shell.AddCmd(&ishell.Cmd{
		Name: "allorders",
		Help: "Allorders get all orders.",
		Func: getAllOrders(),
	})
}
