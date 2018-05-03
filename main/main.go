package main

import (
	"fmt"

	shell "github.com/guard-trader/cmd"
	c "github.com/guard-trader/config"
	"github.com/guard-trader/core"
)

func main() {

	guard, err := core.NewGuard(c.Config.Binance.ApiKey, c.Config.Binance.SecretKey)
	if err != nil {
		fmt.Println(err)
	}
	shell.Run(guard, cancel)
	guard.WakeUp()

}
