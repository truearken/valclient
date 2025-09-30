package valclient

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type ValClient struct {
	Shard  Shard
	Region Region
	Player *struct {
		Uuid string
	}
	Header http.Header
}

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
		Player: &struct{ Uuid string }{
			Uuid: authResp.Subject,
		},
	}, nil
}

var retried = false

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
		if !retried {
			c, err = NewClient()
			if err != nil {
				return err
			}
			if err := c.RunRequest(method, url, in, out); err != nil {
				return errors.New("Is VALORANT running? error occurred while running local request: " + string(bytes))
			}
		}
		return errors.New("Is VALORANT running? error occurred while running local request: " + string(bytes))
	}

	if err := json.Unmarshal(bytes, out); err != nil {
		return err
	}

	return nil
}

/*
Automatically replaces shard, region and puuid in string.
For additional parameters use the corresponding argument
*/
func (c *ValClient) BuildUrl(urlWithParams string, additionalParams ...string) string {
	params := []string{"{shard}", string(c.Shard), "{region}", string(c.Region), "{puuid}", c.Player.Uuid}
	params = append(params, additionalParams...)
	r := strings.NewReplacer(params...)
	return r.Replace(urlWithParams)
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
		return errors.New("Is VALORANT running? error occurred while running local request: " + err.Error())
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("Is VALORANT running? error occurred while running local request: " + string(bytes))
	}

	err = json.Unmarshal(bytes, out)
	if err != nil {
		return err
	}

	return nil
}

type LockfileData struct {
	Name     string
	Pid      string
	Port     string
	Password string
	Protocol string
}

func getLockFile() (*LockfileData, error) {
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
	lockfileStruct := &LockfileData{
		Name:     lockfileKeys[0],
		Pid:      lockfileKeys[1],
		Port:     lockfileKeys[2],
		Password: lockfileKeys[3],
		Protocol: lockfileKeys[4],
	}

	return lockfileStruct, err
}

type LogFileData struct {
	Region        Region
	Shard         Shard
	ClientVersion string
}

func readLogfile() (*LogFileData, error) {
	logfilePath := fmt.Sprintf("%s\\%s", os.Getenv("LOCALAPPDATA"), "VALORANT\\Saved\\Logs\\ShooterGame.log")
	file, err := os.Open(logfilePath)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	regex, err := regexp.Compile("https://glz-(.+?)-1.(.+?).a.pvp.net")
	if err != nil {
		return nil, err
	}

	shardRegionMatches := regex.FindSubmatch(bytes)
	if len(shardRegionMatches) != 3 {
		return nil, errors.New("no valid shard/region matches found in log file")
	}

	regex, err = regexp.Compile("CI server version: (release-[^-]*)(.*)")
	if err != nil {
		return nil, err
	}

	clientVersionMatches := regex.FindSubmatch(bytes)
	if len(clientVersionMatches) != 3 {
		return nil, errors.New("no valid shard/region matches found in log file")
	}

	clientVersionString := fmt.Sprintf("%s-shipping%s", clientVersionMatches[1], clientVersionMatches[2])

	return &LogFileData{
		Region:        Region(shardRegionMatches[1]),
		Shard:         Shard(shardRegionMatches[2]),
		ClientVersion: strings.TrimSpace(clientVersionString),
	}, nil
}
