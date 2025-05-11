package valclient

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

type ItemTypeID string

const (
	ITEM_TYPE_AGENTS        ItemTypeID = "01bb38e1-da47-4e6a-9b3d-945fe4655707"
	ITEM_TYPE_CONTRACTS     ItemTypeID = "f85cb6f7-33e5-4dc8-b609-ec7212301948"
	ITEM_TYPE_SPRAYS        ItemTypeID = "d5f120f8-ff8c-4aac-92ea-f2b5acbe9475"
	ITEM_TYPE_GUN_BUDDIES   ItemTypeID = "dd3bf334-87f3-40bd-b043-682a57a8dc3a"
	ITEM_TYPE_CARDS         ItemTypeID = "3f296c07-64c3-494c-923b-fe692a4fa1bd"
	ITEM_TYPE_SKINS         ItemTypeID = "e7c63390-eda7-46e0-bb7a-a6abdacd2433"
	ITEM_TYPE_SKIN_VARIANTS ItemTypeID = "3ad1b2b2-acdb-4524-852f-954a76ddae0a"
	ITEM_TYPE_TITLES        ItemTypeID = "de7caa6b-adf7-4588-bbd1-143831e786c6"
)

type ContentType string

const (
	CONTENT_TYPE_EPISODE ContentType = "episode"
	CONTENT_TYPE_ACT     ContentType = "act"
)

type XpSourceId string

const (
	XP_SOURCE_ID_TIME_PLAYED      XpSourceId = "time-played"
	XP_SOURCE_ID_MATCH_WIN        XpSourceId = "match-win"
	XP_SOURCE_ID_FIRST_WON_OF_DAY XpSourceId = "first-win-of-the-day"
)

type CompetitiveMovement string

const (
	COMPETITIVE_MOVEMENT_UNKNOWN CompetitiveMovement = "MOVEMENT_UNKNOWN"
)

type ProvisioningFlowID string

const (
	PROVISIONING_FLOW_ID_MATCHMAKING ProvisioningFlowID = "Matchmaking"
	PROVISIONING_FLOW_ID_CUSTOM_GAME ProvisioningFlowID = "CustomGame"
)

type CompletionState string

const (
	COMPLETION_STATE_SURRENDERED CompletionState = "Surrendered"
	COMPLETION_STATE_COMPLETED   CompletionState = "Completed"
	COMPLETION_STATE_VOTE_DRAW   CompletionState = "VoteDraw"
	COMPLETION_STATE_EMPTY       CompletionState = ""
)

type PlatformType string

const (
	PLATFORM_TYPE_PC PlatformType = "PC"
)

type TeamID string

const (
	TEAM_ID_BLUE TeamID = "Blue"
	TEAM_ID_RED  TeamID = "Red"
)

type RoundResult string

const (
	ROUND_RESULT_ELIMINATED          RoundResult = "Eliminated"
	ROUND_RESULT_BOMB_DETONATED      RoundResult = "Bomb detonated"
	ROUND_RESULT_BOMB_DEFUSED        RoundResult = "Bomb defused"
	ROUND_RESULT_SURRENDERED         RoundResult = "Surrendered"
	ROUND_RESULT_ROUND_TIMER_EXPIRED RoundResult = "Round timer expired"
)

type RoundCeremony string

const (
	ROUND_CEREMONY_DEFAULT  RoundCeremony = "CeremonyDefault"
	ROUND_CEREMONY_TEAM_ACE RoundCeremony = "CeremonyTeamAce"
	ROUND_CEREMONY_FLAWLESS RoundCeremony = "CeremonyFlawless"
	ROUND_CEREMONY_CLOSER   RoundCeremony = "CeremonyCloser"
	ROUND_CEREMONY_CLUTCH   RoundCeremony = "CeremonyClutch"
	ROUND_CEREMONY_THRIFTY  RoundCeremony = "CeremonyThrifty"
	ROUND_CEREMONY_ACE      RoundCeremony = "CeremonyAce"
	ROUND_CEREMONY_EMPTY    RoundCeremony = ""
)

type PlantSite string

const (
	PLANT_SITE_A     PlantSite = "A"
	PLANT_SITE_B     PlantSite = "B"
	PLANT_SITE_C     PlantSite = "C"
	PLANT_SITE_EMPTY PlantSite = ""
)

type RoundResultCode string

const (
	ROUND_RESULT_CODE_ELIMINATION RoundResultCode = "Elimination"
	ROUND_RESULT_CODE_DETONATE    RoundResultCode = "Detonate"
	ROUND_RESULT_CODE_DEFUSE      RoundResultCode = "Defuse"
	ROUND_RESULT_CODE_SURRENDERED RoundResultCode = "Surrendered"
	ROUND_RESULT_CODE_EMPTY       RoundResultCode = ""
)

type DamageType string

const (
	DAMAGE_TYPE_WEAPON  DamageType = "Weapon"
	DAMAGE_TYPE_BOMB    DamageType = "Bomb"
	DAMAGE_TYPE_ABILITY DamageType = "Ability"
	DAMAGE_TYPE_FALL    DamageType = "Fall"
	DAMAGE_TYPE_MELEE   DamageType = "Melee"
	DAMAGE_TYPE_INVALID DamageType = "Invalid"
	DAMAGE_TYPE_EMPTY   DamageType = ""
)

type DamageItem string

const (
	DAMAGE_ITEM_ULTIMATE        DamageItem = "Ultimate"
	DAMAGE_ITEM_ABILITY1        DamageItem = "Ability1"
	DAMAGE_ITEM_ABILITY2        DamageItem = "Ability2"
	DAMAGE_ITEM_GRENADE_ABILITY DamageItem = "GrenadeAbility"
	DAMAGE_ITEM_PRIMARY         DamageItem = "Primary"
	DAMAGE_ITEM_EMPTY           DamageItem = ""
)

type QueueID string

const (
	QUEUE_COMPETITIVE     QueueID = "competitive"
	QUEUE_DEATHMATCH      QueueID = "deathmatch"
	QUEUE_TEAM_DEATHMATCH QueueID = "hurm"
	QUEUE_CUSTOM          QueueID = "custom"
	QUEUE_ESCALATION      QueueID = "ggteam"
	QUEUE_NEW_MAP         QueueID = "newmap"
	QUEUE_REPLICATION     QueueID = "onefa"
	QUEUE_PREMIER         QueueID = "premier"
	QUEUE_SNOWBALL_FIGHT  QueueID = "snowball"
	QUEUE_SPIKERUSH       QueueID = "spikerush"
	QUEUE_SWIFTPLAY       QueueID = "swiftplay"
	QUEUE_UNRATED         QueueID = "unrated"
)

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

type ActiveExpressions struct {
	TypeID  string `json:"TypeID"`
	AssetID string `json:"AssetID"`
}

type Identity struct {
	PlayerCardID           string `json:"PlayerCardID"`
	PlayerTitleID          string `json:"PlayerTitleID"`
	AccountLevel           int    `json:"AccountLevel"`
	PreferredLevelBorderID string `json:"PreferredLevelBorderID"`
	HideAccountLevel       bool   `json:"HideAccountLevel"`
}
