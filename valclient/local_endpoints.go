package valclient

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/websocket"
)

type GetHelpResponse struct {
	Events    map[string]string `json:"Events"`
	Functions map[string]string `json:"Functions"`
	Types     map[string]string `json:"Types"`
}

func (c *ValClient) GetHelp() (*GetHelpResponse, error) {
	help := new(GetHelpResponse)
	if err := c.RunLocalRequest(http.MethodGet, "/help", nil, help); err != nil {
		return nil, err
	}

	return help, nil
}

type LocalWebsocket struct {
	Conn *websocket.Conn
}

func (c *ValClient) GetLocalWebsocket() (*LocalWebsocket, error) {
	dialer := *websocket.DefaultDialer
	dialer.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	url := fmt.Sprintf("wss://127.0.0.1:%s/", c.Local.Port)
	header := http.Header{}
	header.Set("Authorization", "Basic "+basicAuth("riot", c.Local.Password))

	conn, _, err := dialer.Dial(url, header)
	if err != nil {
		return nil, err
	}
	lws := &LocalWebsocket{Conn: conn}
	return lws, nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func (lws *LocalWebsocket) Close() error {
	return lws.Conn.Close()
}

type WampMessage []any

// a list of events can be found in the GetHelp request
func (lws *LocalWebsocket) SubscribeEvent(event string) error {
	msg := WampMessage{}
	msg = append(msg, 5)
	msg = append(msg, event)
	return lws.Conn.WriteJSON(msg)
}

func (lws *LocalWebsocket) UnsubscribeEvent(event string) error {
	msg := WampMessage{}
	msg = append(msg, 6)
	msg = append(msg, event)
	return lws.Conn.WriteJSON(msg)
}

type LocalWebsocketApiEvent struct {
	OpCode     int                      `json:"-"`
	Event      string                   `json:"-"`
	Payload    *LocalWebsocketEventData `json:"-"`
	RawPayload json.RawMessage          `json:"-"`
}

type LocalWebsocketEventData struct {
	Data      map[string]any `json:"data"`
	EventType string         `json:"eventType"`
	URI       string         `json:"uri"`
}

func (lws *LocalWebsocket) Read(events chan<- *LocalWebsocketApiEvent) error {
	for {
		rawArr := []json.RawMessage{}
		err := lws.Conn.ReadJSON(&rawArr)
		if err == io.ErrUnexpectedEOF {
			continue
		}
		if err != nil {
			return err
		}

		var opCode int
		var event string
		eventData := new(LocalWebsocketEventData)
		if err := json.Unmarshal(rawArr[0], &opCode); err != nil {
			return err
		}
		if err := json.Unmarshal(rawArr[1], &event); err != nil {
			return err
		}
		// some events don't contain json
		_ = json.Unmarshal(rawArr[2], eventData)

		events <- &LocalWebsocketApiEvent{OpCode: opCode, Event: event, Payload: eventData, RawPayload: rawArr[2]}
	}
}
