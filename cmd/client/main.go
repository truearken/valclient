package main

import (
	"log"

	valclient "github.com/truearken/valclient/client"
)

func main() {
	client, err := valclient.NewClient()
	if err != nil {
		panic(err)
	}

	log.Print(client.Header.Get("X-Riot-Clientversion"))

	loadout, err := client.GetPlayerLoadout()
	if err != nil {
		panic(err)
	}

	log.Print(loadout)
}
