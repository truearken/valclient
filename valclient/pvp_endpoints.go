package valclient

import (
	"fmt"
	"net/http"
	"time"
)

type GetContentResponse struct {
	DisabledIDs []any `json:"DisabledIDs"`
	Seasons     []struct {
		ID        string      `json:"ID"`
		Name      string      `json:"Name"`
		Type      ContentType `json:"Type"`
		StartTime time.Time   `json:"StartTime"`
		EndTime   time.Time   `json:"EndTime"`
		IsActive  bool        `json:"IsActive"`
	} `json:"Seasons"`
	Events []struct {
		ID        string    `json:"ID"`
		Name      string    `json:"Name"`
		StartTime time.Time `json:"StartTime"`
		EndTime   time.Time `json:"EndTime"`
		IsActive  bool      `json:"IsActive"`
	} `json:"Events"`
}

func (c *ValClient) GetContentRequest() (*GetContentResponse, error) {
	url := fmt.Sprintf("https://shared.%s.a.pvp.net/content-service/v3/content", c.Shard)
	content := new(GetContentResponse)

	if err := c.RunRequest(http.MethodGet, url, nil, content); err != nil {
		return nil, err
	}

	return content, nil
}

type GetAccountXpResponse struct {
	Version  int    `json:"Version"`
	Subject  string `json:"Subject"`
	Progress struct {
		Level int `json:"Level"`
		XP    int `json:"XP"`
	} `json:"Progress"`
	History []struct {
		ID            string `json:"ID"`
		MatchStart    string `json:"MatchStart"`
		StartProgress struct {
			Level int `json:"Level"`
			XP    int `json:"XP"`
		} `json:"StartProgress"`
		EndProgress struct {
			Level int `json:"Level"`
			XP    int `json:"XP"`
		} `json:"EndProgress"`
		XPDelta   int `json:"XPDelta"`
		XPSources []struct {
			ID     XpSourceId `json:"ID"`
			Amount int        `json:"Amount"`
		} `json:"XPSources"`
		XPMultipliers []any `json:"XPMultipliers"`
	} `json:"History"`
	LastTimeGrantedFirstWin   string `json:"LastTimeGrantedFirstWin"`
	NextTimeFirstWinAvailable string `json:"NextTimeFirstWinAvailable"`
}

func (c *ValClient) GetAccountXp() (*GetAccountXpResponse, error) {
	url := fmt.Sprintf("https://pd.%s.a.pvp.net/account-xp/v1/players/%s", c.Shard, c.Player.Uuid)
	accountXp := new(GetAccountXpResponse)

	if err := c.RunRequest(http.MethodGet, url, nil, accountXp); err != nil {
		return nil, err
	}

	return accountXp, nil
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

	if err := c.RunRequest(http.MethodGet, url, nil, loadout); err != nil {
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

	if err := c.RunRequest(http.MethodPut, url, loadout, responseloadout); err != nil {
		return nil, err
	}

	return responseloadout, nil
}

type GetOwnedItemsResponse struct {
	ItemTypeID   ItemTypeID `json:"ItemTypeID"`
	Entitlements []struct {
		TypeID string `json:"TypeID"`
		ItemID string `json:"ItemID"`
	} `json:"Entitlements"`
}

func (c *ValClient) GetOwnedItems(itemType ItemTypeID) (*GetOwnedItemsResponse, error) {
	url := fmt.Sprintf("https://pd.%s.a.pvp.net/store/v1/entitlements/%s/%s", c.Shard, c.Player.Uuid, itemType)
	ownedItems := new(GetOwnedItemsResponse)

	err := c.RunRequest(http.MethodGet, url, nil, ownedItems)
	if err != nil {
		return nil, err
	}

	return ownedItems, nil
}
