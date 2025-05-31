package valclient

import "net/http"

type GetPreGamePlayerResponse struct {
	Subject string `json:"Subject"`
	MatchID string `json:"MatchID"`
	Version int    `json:"Version"`
}

func (c *ValClient) GetPreGamePlayer() (*GetPreGamePlayerResponse, error) {
	url := c.BuildUrl("https://glz-{region}-1.{shard}.a.pvp.net/pregame/v1/players/{puuid}")
	prePlayer := new(GetPreGamePlayerResponse)

	if err := c.RunRequest(http.MethodGet, url, nil, prePlayer); err != nil {
		return nil, err
	}

	return prePlayer, nil
}

type GetPreGameMatchResponse struct {
	ID      string `json:"ID"`
	Version int    `json:"Version"`
	Teams   []struct {
		TeamID  TeamID `json:"TeamID"`
		Players []struct {
			Subject                 string                  `json:"Subject"`
			CharacterID             string                  `json:"CharacterID"`
			CharacterSelectionState CharacterSelectionState `json:"CharacterSelectionState"`
			PregamePlayerState      PregamePlayerState      `json:"PregamePlayerState"`
			CompetitiveTier         int                     `json:"CompetitiveTier"`
			PlayerIdentity          struct {
				Subject                string `json:"Subject"`
				PlayerCardID           string `json:"PlayerCardID"`
				PlayerTitleID          string `json:"PlayerTitleID"`
				AccountLevel           int    `json:"AccountLevel"`
				PreferredLevelBorderID string `json:"PreferredLevelBorderID"`
				Incognito              bool   `json:"Incognito"`
				HideAccountLevel       bool   `json:"HideAccountLevel"`
			} `json:"PlayerIdentity"`
			SeasonalBadgeInfo struct {
				SeasonID        string `json:"SeasonID"`
				NumberOfWins    int    `json:"NumberOfWins"`
				WinsByTier      any    `json:"WinsByTier"`
				Rank            int    `json:"Rank"`
				LeaderboardRank int    `json:"LeaderboardRank"`
			} `json:"SeasonalBadgeInfo"`
			IsCaptain bool `json:"IsCaptain"`
		} `json:"Players"`
	} `json:"Teams"`
	AllyTeam *struct {
		TeamID  TeamID `json:"TeamID"`
		Players []struct {
			Subject                 string                  `json:"Subject"`
			CharacterID             string                  `json:"CharacterID"`
			CharacterSelectionState CharacterSelectionState `json:"CharacterSelectionState"`
			PregamePlayerState      PregamePlayerState      `json:"PregamePlayerState"`
			CompetitiveTier         int                     `json:"CompetitiveTier"`
			PlayerIdentity          struct {
				Subject                string `json:"Subject"`
				PlayerCardID           string `json:"PlayerCardID"`
				PlayerTitleID          string `json:"PlayerTitleID"`
				AccountLevel           int    `json:"AccountLevel"`
				PreferredLevelBorderID string `json:"PreferredLevelBorderID"`
				Incognito              bool   `json:"Incognito"`
				HideAccountLevel       bool   `json:"HideAccountLevel"`
			} `json:"PlayerIdentity"`
			SeasonalBadgeInfo struct {
				SeasonID        string `json:"SeasonID"`
				NumberOfWins    int    `json:"NumberOfWins"`
				WinsByTier      any    `json:"WinsByTier"`
				Rank            int    `json:"Rank"`
				LeaderboardRank int    `json:"LeaderboardRank"`
			} `json:"SeasonalBadgeInfo"`
			IsCaptain bool `json:"IsCaptain"`
		} `json:"Players"`
	} `json:"AllyTeam"`
	EnemyTeam *struct {
		TeamID  TeamID `json:"TeamID"`
		Players []struct {
			Subject                 string                  `json:"Subject"`
			CharacterID             string                  `json:"CharacterID"`
			CharacterSelectionState CharacterSelectionState `json:"CharacterSelectionState"`
			PregamePlayerState      PregamePlayerState      `json:"PregamePlayerState"`
			CompetitiveTier         int                     `json:"CompetitiveTier"`
			PlayerIdentity          struct {
				Subject                string `json:"Subject"`
				PlayerCardID           string `json:"PlayerCardID"`
				PlayerTitleID          string `json:"PlayerTitleID"`
				AccountLevel           int    `json:"AccountLevel"`
				PreferredLevelBorderID string `json:"PreferredLevelBorderID"`
				Incognito              bool   `json:"Incognito"`
				HideAccountLevel       bool   `json:"HideAccountLevel"`
			} `json:"PlayerIdentity"`
			SeasonalBadgeInfo struct {
				SeasonID        string `json:"SeasonID"`
				NumberOfWins    int    `json:"NumberOfWins"`
				WinsByTier      any    `json:"WinsByTier"`
				Rank            int    `json:"Rank"`
				LeaderboardRank int    `json:"LeaderboardRank"`
			} `json:"SeasonalBadgeInfo"`
			IsCaptain bool `json:"IsCaptain"`
		} `json:"Players"`
	} `json:"EnemyTeam"`
	ObserverSubjects     []any  `json:"ObserverSubjects"`
	MatchCoaches         []any  `json:"MatchCoaches"`
	EnemyTeamSize        int    `json:"EnemyTeamSize"`
	EnemyTeamLockCount   int    `json:"EnemyTeamLockCount"`
	PregameState         string `json:"PregameState"`
	LastUpdated          string `json:"LastUpdated"`
	MapID                string `json:"MapID"`
	MapSelectPool        []any  `json:"MapSelectPool"`
	BannedMapIDs         []any  `json:"BannedMapIDs"`
	CastedVotes          any    `json:"CastedVotes"`
	MapSelectSteps       []any  `json:"MapSelectSteps"`
	MapSelectStep        int    `json:"MapSelectStep"`
	Team1                TeamID `json:"Team1"`
	GamePodID            string `json:"GamePodID"`
	Mode                 string `json:"Mode"`
	VoiceSessionID       string `json:"VoiceSessionID"`
	MUCName              string `json:"MUCName"`
	TeamMatchToken       string `json:"TeamMatchToken"`
	QueueID              string `json:"QueueID"`
	ProvisioningFlowID   string `json:"ProvisioningFlowID"`
	IsRanked             bool   `json:"IsRanked"`
	PhaseTimeRemainingNS int    `json:"PhaseTimeRemainingNS"`
	StepTimeRemainingNS  int    `json:"StepTimeRemainingNS"`
	AltModesFlagADA      bool   `json:"altModesFlagADA"`
	TournamentMetadata   any    `json:"TournamentMetadata"`
	RosterMetadata       any    `json:"RosterMetadata"`
}

func (c *ValClient) GetPreGameMatch() (*GetPreGameMatchResponse, error) {
	prePlayer, err := c.GetPreGamePlayer()
	if err != nil {
		return nil, err
	}

	url := c.BuildUrl("https://glz-{region}-1.{shard}.a.pvp.net/pregame/v1/matches/{matchId}", "{matchId}", prePlayer.MatchID)
	preMatch := new(GetPreGameMatchResponse)

	if err := c.RunRequest(http.MethodGet, url, nil, preMatch); err != nil {
		return nil, err
	}

	return preMatch, nil
}
