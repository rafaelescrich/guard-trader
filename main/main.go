package main

import (
	"fmt"

	shell "github.com/guard-trader/cmd"
	mng "github.com/guard-trader/manager"
)

func main() {

	manager, err := mng.NewManager()
	if err != nil {
		fmt.Println(err)
	}
	shell.Run(manager, cancel)
	manager.WakeUp()
}
