package client

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

type Player struct {
	Uuid string
}

type ValClient struct {
	Shard  Shard
	Region Region
	Player *Player
	Header http.Header
}
