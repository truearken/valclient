package valclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func NewClient() (*ValClient, error) {
	lockfile, err := getLockFile()
	if err != nil {
		return nil, err
	}

	authResp, err := authenticate(lockfile.Port, lockfile.Password)
	if err != nil {
		return nil, err
	}

	logfile, err := readLogfile()
	if err != nil {
		return nil, err
	}

	header := http.Header{}
	header.Add("Authorization", fmt.Sprintf("Bearer %s", authResp.AccessToken))
	header.Add("X-Riot-Entitlements-JWT", authResp.Token)
	header.Add("X-Riot-ClientPlatform", CLIENT_PLATFORM_B64)
	header.Add("X-Riot-ClientVersion", logfile.ClientVersion)

	return &ValClient{
		Region: logfile.Region,
		Shard:  logfile.Shard,
		Header: header,
		Player: &Player{
			Uuid: authResp.Subject,
		},
	}, nil
}

type GetPlayerLoadoutRequest struct {
	Subject           string               `json:"Subject"`
	Version           int                  `json:"Version"`
	Guns              []*Gun               `json:"Guns"`
	ActiveExpressions []*ActiveExpressions `json:"ActiveExpressions"`
	Identity          *Identity            `json:"Identity"`
	Incognito         bool                 `json:"Incognito"`
}

func (c *ValClient) GetPlayerLoadout() (*GetPlayerLoadoutRequest, error) {
	url := fmt.Sprintf("https://pd.%s.a.pvp.net/personalization/v3/players/%s/playerloadout", c.Shard, c.Player.Uuid)
	loadout := new(GetPlayerLoadoutRequest)

	err := c.RunRequest(http.MethodGet, url, nil, loadout)
	if err != nil {
		return nil, err
	}

	return loadout, nil
}

type SetPlayerLoadoutRequest struct {
	Guns              []*Gun               `json:"Guns"`
	ActiveExpressions []*ActiveExpressions `json:"ActiveExpressions"`
	Identity          *Identity            `json:"Identity"`
	Incognito         bool                 `json:"Incognito"`
}

func (c *ValClient) SetPlayerLoadout(loadout *SetPlayerLoadoutRequest) (*GetPlayerLoadoutRequest, error) {
	url := fmt.Sprintf("https://pd.%s.a.pvp.net/personalization/v3/players/%s/playerloadout", c.Shard, c.Player.Uuid)
	responseloadout := new(GetPlayerLoadoutRequest)

	err := c.RunRequest(http.MethodPut, url, loadout, responseloadout)
	if err != nil {
		return nil, err
	}

	return responseloadout, nil
}

type OwnedItems struct {
	ItemTypeID   string `json:"ItemTypeID"`
	Entitlements []struct {
		TypeID string `json:"TypeID"`
		ItemID string `json:"ItemID"`
	} `json:"Entitlements"`
}

func (c *ValClient) GetOwnedItems(itemType ItemTypeId) (*OwnedItems, error) {
	url := fmt.Sprintf("https://pd.%s.a.pvp.net/store/v1/entitlements/%s/%s", c.Shard, c.Player.Uuid, itemType)
	ownedItems := new(OwnedItems)

	err := c.RunRequest(http.MethodGet, url, nil, ownedItems)
	if err != nil {
		return nil, err
	}

	return ownedItems, nil
}

func (c *ValClient) RunRequest(method, url string, in any, out any) error {
	body := new(bytes.Buffer)
	if in != nil {
		bytes, err := json.Marshal(in)
		if err != nil {
			return err
		}
		_, err = body.Write(bytes)
		if err != nil {
			return err
		}
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}
	req.Header = c.Header

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("Is VALORANT running? error occurred while running local request: " + string(bytes))
	}

	if err := json.Unmarshal(bytes, out); err != nil {
		return err
	}

	return nil
}
