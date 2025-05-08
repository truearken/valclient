package valclient

import "net/http"

const CLIENT_PLATFORM_B64 = "ew0KCSJwbGF0Zm9ybVR5cGUiOiAiUEMiLA0KCSJwbGF0Zm9ybU9TIjogIldpbmRvd3MiLA0KCSJwbGF0Zm9ybU9TVmVyc2lvbiI6ICIxMC4wLjE5MDQyLjEuMjU2LjY0Yml0IiwNCgkicGxhdGZvcm1DaGlwc2V0IjogIlVua25vd24iDQp9"

type Region string

const (
	REGION_NA    Region = "na"
	REGION_EU    Region = "eu"
	REGION_LATAM Region = "latam"
	REGION_BR    Region = "br"
	REGION_AP    Region = "ap"
	REGION_KR    Region = "kr"
)

type Shard string

const (
	SHARD_NA Shard = Shard(REGION_NA)
	SHARD_EU Shard = Shard(REGION_EU)
	SHARD_AP Shard = Shard(REGION_AP)
	SHARD_KR Shard = Shard(REGION_KR)
)

var ShardForRegion = map[Region]Shard{
	REGION_NA:    SHARD_NA,
	REGION_EU:    SHARD_EU,
	REGION_LATAM: SHARD_NA,
	REGION_BR:    SHARD_NA,
	REGION_AP:    SHARD_AP,
	REGION_KR:    SHARD_KR,
}

type ItemTypeId string

const (
	ITEM_TYPE_AGENTS        ItemTypeId = "01bb38e1-da47-4e6a-9b3d-945fe4655707"
	ITEM_TYPE_CONTRACTS     ItemTypeId = "f85cb6f7-33e5-4dc8-b609-ec7212301948"
	ITEM_TYPE_SPRAYS        ItemTypeId = "d5f120f8-ff8c-4aac-92ea-f2b5acbe9475"
	ITEM_TYPE_GUN_BUDDIES   ItemTypeId = "dd3bf334-87f3-40bd-b043-682a57a8dc3a"
	ITEM_TYPE_CARDS         ItemTypeId = "3f296c07-64c3-494c-923b-fe692a4fa1bd"
	ITEM_TYPE_SKINS         ItemTypeId = "e7c63390-eda7-46e0-bb7a-a6abdacd2433"
	ITEM_TYPE_SKIN_VARIANTS ItemTypeId = "3ad1b2b2-acdb-4524-852f-954a76ddae0a"
	ITEM_TYPE_TITLES        ItemTypeId = "de7caa6b-adf7-4588-bbd1-143831e786c6"
)

type Player struct {
	Uuid string
}

type ValClient struct {
	Shard  Shard
	Region Region
	Player *Player
	Header http.Header
}

type Gun struct {
	ID              string `json:"ID"`
	CharmInstanceID string `json:"CharmInstanceID,omitempty"`
	CharmID         string `json:"CharmID,omitempty"`
	CharmLevelID    string `json:"CharmLevelID,omitempty"`
	SkinID          string `json:"SkinID"`
	SkinLevelID     string `json:"SkinLevelID"`
	ChromaID        string `json:"ChromaID"`
	Attachments     []any  `json:"Attachments"`
}

type Spray struct {
	EquipSlotID string `json:"EquipSlotID"`
	SprayID     string `json:"SprayID"`
}

type Identity struct {
	PlayerCardID           string `json:"PlayerCardID"`
	PlayerTitleID          string `json:"PlayerTitleID"`
	AccountLevel           int    `json:"AccountLevel"`
	PreferredLevelBorderID string `json:"PreferredLevelBorderID"`
	HideAccountLevel       bool   `json:"HideAccountLevel"`
}
