package client

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

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

type Lockfile struct {
	Name     string
	Pid      string
	Port     string
	Password string
	Protocol string
}

type Player struct {
	Uuid string
	Name string
	Tag  string
}

type ValClient struct {
	Shard  Shard
	Region Region
	Player *Player
	Header http.Header
}

func NewValClient(region Region) (*ValClient, error) {
	lockfile, err := getLockFile()
	if err != nil {
		return nil, err
	}

	authResp, err := authenticate(lockfile.Port, lockfile.Password)
	if err != nil {
		return nil, err
	}

	clientVersion, err := getClientVersion()
	if err != nil {
		return nil, err
	}

	header := http.Header{}
	header.Add("Authorization", fmt.Sprintf("Bearer %s", authResp.AccessToken))
	header.Add("X-Riot-Entitlements-JWT", authResp.Token)
	header.Add("X-Riot-ClientPlatform", CLIENT_PLATFORM_B64)
	header.Add("X-Riot-ClientVersion", clientVersion)

	return &ValClient{
		Region: region,
		Shard:  ShardForRegion[region],
		Header: header,
		Player: &Player{
			Uuid: authResp.Subject,
		},
	}, nil
}

func (c *ValClient) GetPlayerLoadout() (string, error) {
	url := fmt.Sprintf("https://pd.%s.a.pvp.net/personalization/v2/players/%s/playerloadout", c.Shard, c.Player.Uuid)
	resp, err := c.RunRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (c *ValClient) RunRequest(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header = c.Header

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type AuthenticateResponse struct {
	AccessToken string `json:"accessToken"`
	Subject     string `json:"subject"`
	Token       string `json:"token"`
}

func authenticate(port, password string) (*AuthenticateResponse, error) {
	authResp := new(AuthenticateResponse)
	if err := runLocalRequest(port, password, "/entitlements/v1/token", authResp); err != nil {
		return nil, err
	}

	return authResp, nil
}

func runLocalRequest(port, password, endpoint string, out any) error {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://127.0.0.1:%s%s", port, endpoint), nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth("riot", password)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, out)
	if err != nil {
		return err
	}

	return nil
}

func getLockFile() (*Lockfile, error) {
	lockfilePath := fmt.Sprintf("%s\\%s", os.Getenv("LOCALAPPDATA"), "Riot Games\\Riot Client\\Config\\lockfile")
	lockfile, err := os.Open(lockfilePath)
	if err != nil {
		return nil, err
	}

	fileData, err := io.ReadAll(lockfile)
	if err != nil {
		return nil, err
	}
	lockfileKeys := strings.Split(string(fileData), ":")
	lockfileStruct := &Lockfile{
		Name:     lockfileKeys[0],
		Pid:      lockfileKeys[1],
		Port:     lockfileKeys[2],
		Password: lockfileKeys[3],
		Protocol: lockfileKeys[4],
	}

	return lockfileStruct, err
}

type GetVersionResponse struct {
	Data struct {
		RiotClientVersion string `json:"riotClientVersion"`
	} `json:"data"`
}

func getClientVersion() (string, error) {
	resp, err := http.Get("https://valorant-api.com/v1/version")
	if err != nil {
		return "", err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	versResp := new(GetVersionResponse)
	if err := json.Unmarshal(bytes, versResp); err != nil {
		return "", err
	}

	return versResp.Data.RiotClientVersion, nil
}
