package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"

	"github.com/truearken/valclient/valclient"
)

type Root struct {
	Presences []Presence `json:"presences"`
}

type Presence struct {
	Puuid   string `json:"puuid"`
	Private string `json:"private"`
}

type MatchPresenceData struct {
	SessionLoopState string `json:"sessionLoopState"`
	ProvisioningFlow string `json:"provisioningFlow"`
	MatchMap         string `json:"matchMap"`
	QueueID          string `json:"queueId"`
}

type PrivateData struct {
	MatchPresenceData MatchPresenceData `json:"matchPresenceData"`
}

// gives information about the players game state
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

		var root Root
		if err := json.Unmarshal(dataBytes, &root); err != nil {
			log.Fatal("Failed to unmarshal root JSON:", err)
		}

		if len(root.Presences) == 0 {
			log.Fatal("No presences found")
		}

		uuid := root.Presences[0].Puuid
		if uuid != client.Player.Uuid {
			continue
		}

		decodedPrivate, err := base64.StdEncoding.DecodeString(root.Presences[0].Private)
		if err != nil {
			log.Fatal("Failed to decode base64:", err)
		}

		var privateData PrivateData
		if err := json.Unmarshal(decodedPrivate, &privateData); err != nil {
			log.Fatal("Failed to unmarshal private JSON:", err)
		}

		fmt.Printf("MatchPresenceData:\n")
		fmt.Println("  State:", privateData.MatchPresenceData.SessionLoopState)
		fmt.Println("  Flow:", privateData.MatchPresenceData.ProvisioningFlow)
		fmt.Println("  Map:", privateData.MatchPresenceData.MatchMap)
		fmt.Println("  Queue:", privateData.MatchPresenceData.QueueID)
	}
}
