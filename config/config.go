package config

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/BurntSushi/toml"
)

type config struct {
	Binance struct {
		APIKey, SecretKey string
	}
	Database struct {
		Path string
	}
}

// Config holds the configuration file(config.toml) contents
var Config config

func init() {
	_, filename, _, _ := runtime.Caller(0)
	directory := path.Dir(filename)
	LoadConfig(directory + "/config.toml")
}

// LoadConfig loads a file to populate config struct
func LoadConfig(configFile string) {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		fmt.Println("Config file \"" + configFile + "\" does not exist.")
		return
	} else if err != nil {
		fmt.Println(err.Error())
		return
	}

	if _, err := toml.DecodeFile(configFile, &Config); err != nil {
		fmt.Println(err.Error())
		return
	}
}
