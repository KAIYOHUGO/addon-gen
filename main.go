package main

import (
	"encoding/json"
	"os"

	"github.com/creasty/defaults"
)

func main() {
	config := new(Config)
	{
		cf, err := os.Open("./manifest.json")
		if err != nil {
			panic(err)
		}
		json.NewDecoder(cf).Decode(config)
		defaults.MustSet(config)
	}
	if err := copyfolder(config); err != nil {
		panic(err)
	}

	// compile py to js
	if err := compile(config, "./scripts/server"); err != nil {
		panic(err)
	}
	if err := compile(config, "./scripts/client"); err != nil {
		panic(err)
	}
}
