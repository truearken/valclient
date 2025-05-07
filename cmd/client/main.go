package main

import (
	"log"

	"github.com/truearken/valclient/client"
)

func main() {
	client, err := client.NewValClient(client.REGION_EU)
	if err != nil {
		panic(err)
	}

	loadout, err := client.GetPlayerLoadout()
	if err != nil {
		panic(err)
	}

	log.Print(loadout)
}
