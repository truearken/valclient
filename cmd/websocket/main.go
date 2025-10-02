package main

import (
	"encoding/json"
	"log/slog"

	"github.com/truearken/valclient/valclient"
)

func main() {
	client, err := valclient.NewClient()
	if err != nil {
		panic(err)
	}

	ws, err := client.GetLocalWebsocket()
	if err != nil {
		panic(err)
	}
	defer ws.Close()

	if err := ws.SubscribeEvent("OnJsonApiEvent_chat_v4_presences"); err != nil {
		panic(err)
	}

	events := make(chan valclient.LocalWebsocketEventData)
	go func() {
		if err := ws.Read(events); err != nil {
			panic(err)
		}
	}()

	for event := range events {
		dataBytes, err := json.Marshal(event.Data)
		if err != nil {
			panic(err)
		}

		slog.Info("data", "data", string(dataBytes))
	}
}
