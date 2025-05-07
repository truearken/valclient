package client

import (
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

// func (c *ValClient) SetPlayerLoadout(loadout io.Reader) error {
// 	url := fmt.Sprintf("https://pd.%s.a.pvp.net/personalization/v2/players/%s/playerloadout", c.Shard, c.Player.Uuid)
// 	resp, err := c.RunRequest(http.MethodPut, url, loadout)
// 	if err != nil {
// 		return err
// 	}
// }

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
