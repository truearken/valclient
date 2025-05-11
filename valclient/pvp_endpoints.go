package valclient

import (
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

func (c *ValClient) GetContent() (*GetContentResponse, error) {
	url := c.buildUrl("https://shared.{shard}.a.pvp.net/content-service/v3/content")
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
	url := c.buildUrl("https://pd.{shard}.a.pvp.net/account-xp/v1/players/{puuid}")
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
	url := c.buildUrl("https://pd.{shard}.a.pvp.net/personalization/v3/players/{puuid}/playerloadout")
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
	url := c.buildUrl("https://pd.{shard}.a.pvp.net/personalization/v3/players/{puuid}/playerloadout")
	responseloadout := new(GetPlayerLoadoutRequest)

	if err := c.RunRequest(http.MethodPut, url, loadout, responseloadout); err != nil {
		return nil, err
	}

	return responseloadout, nil
}

type GetPlayerMmrResponse struct {
	Version                     int    `json:"Version"`
	Subject                     string `json:"Subject"`
	NewPlayerExperienceFinished bool   `json:"NewPlayerExperienceFinished"`
	QueueSkills                 map[string]struct {
		TotalGamesNeededForRating         int `json:"TotalGamesNeededForRating"`
		TotalGamesNeededForLeaderboard    int `json:"TotalGamesNeededForLeaderboard"`
		CurrentSeasonGamesNeededForRating int `json:"CurrentSeasonGamesNeededForRating"`
		SeasonalInfoBySeasonID            map[string]struct {
			SeasonID                   string         `json:"SeasonID"`
			NumberOfWins               int            `json:"NumberOfWins"`
			NumberOfWinsWithPlacements int            `json:"NumberOfWinsWithPlacements"`
			NumberOfGames              int            `json:"NumberOfGames"`
			Rank                       int            `json:"Rank"`
			CapstoneWins               int            `json:"CapstoneWins"`
			LeaderboardRank            int            `json:"LeaderboardRank"`
			CompetitiveTier            int            `json:"CompetitiveTier"`
			RankedRating               int            `json:"RankedRating"`
			WinsByTier                 map[string]int `json:"WinsByTier"`
			GamesNeededForRating       int            `json:"GamesNeededForRating"`
			TotalWinsNeededForRank     int            `json:"TotalWinsNeededForRank"`
		} `json:"SeasonalInfoBySeasonID"`
	} `json:"QueueSkills"`
	LatestCompetitiveUpdate struct {
		MatchID                      string              `json:"MatchID"`
		MapID                        string              `json:"MapID"`
		SeasonID                     string              `json:"SeasonID"`
		MatchStartTime               int                 `json:"MatchStartTime"`
		TierAfterUpdate              int                 `json:"TierAfterUpdate"`
		TierBeforeUpdate             int                 `json:"TierBeforeUpdate"`
		RankedRatingAfterUpdate      int                 `json:"RankedRatingAfterUpdate"`
		RankedRatingBeforeUpdate     int                 `json:"RankedRatingBeforeUpdate"`
		RankedRatingEarned           int                 `json:"RankedRatingEarned"`
		RankedRatingPerformanceBonus int                 `json:"RankedRatingPerformanceBonus"`
		CompetitiveMovement          CompetitiveMovement `json:"CompetitiveMovement"`
		AFKPenalty                   int                 `json:"AFKPenalty"`
	} `json:"LatestCompetitiveUpdate"`
	IsLeaderboardAnonymized bool `json:"IsLeaderboardAnonymized"`
	IsActRankBadgeHidden    bool `json:"IsActRankBadgeHidden"`
}

func (c *ValClient) GetPlayerMmr() (*GetPlayerMmrResponse, error) {
	url := c.buildUrl("https://pd.{shard}.a.pvp.net/mmr/v1/players/{puuid}")
	ownedItems := new(GetPlayerMmrResponse)

	err := c.RunRequest(http.MethodGet, url, nil, ownedItems)
	if err != nil {
		return nil, err
	}

	return ownedItems, nil
}
