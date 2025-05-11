package valclient

import (
	"net/http"
)

type GetOwnedItemsResponse struct {
	ItemTypeID   ItemTypeID `json:"ItemTypeID"`
	Entitlements []struct {
		TypeID string `json:"TypeID"`
		ItemID string `json:"ItemID"`
	} `json:"Entitlements"`
}

func (c *ValClient) GetOwnedItems(itemType ItemTypeID) (*GetOwnedItemsResponse, error) {
	url := c.buildUrl("https://pd.{shard}.a.pvp.net/store/v1/entitlements/{puuid}/{ItemTypeID}", "{ItemTypeID}", string(itemType))
	ownedItems := new(GetOwnedItemsResponse)

	err := c.RunRequest(http.MethodGet, url, nil, ownedItems)
	if err != nil {
		return nil, err
	}

	return ownedItems, nil
}
