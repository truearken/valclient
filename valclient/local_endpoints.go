package valclient

import "net/http"

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
