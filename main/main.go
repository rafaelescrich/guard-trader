package main

import (
	"context"

	shell "github.com/guard-trader/cmd"
	mng "github.com/guard-trader/manager"
)

func main() {

	manager := mng.NewManager()

	_, cancel := context.WithCancel(context.Background())

	shell.Run(manager, cancel)
	manager.WakeUp()
}
