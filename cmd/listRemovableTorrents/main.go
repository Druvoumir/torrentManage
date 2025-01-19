package main

import (
	"os"

	"patu.re/torrentManager/pkg/config"
	"patu.re/torrentManager/pkg/torrentManager"
)

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config.json"
	}

	conf, err := config.Load(configPath)
	if err != nil {
		panic(err)
	}

	torrentManager.ListRemovableTorrents(conf)
}
