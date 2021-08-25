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

}
