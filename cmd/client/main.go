package main

import (
	"log"

	"github.com/truearken/valclient/valclient"
)

func main() {
	client, err := valclient.NewClient()
	if err != nil {
		panic(err)
	}

	loadout, err := client.GetPlayerLoadout()
	if err != nil {
		panic(err)
	}

	log.Print(loadout)
}
