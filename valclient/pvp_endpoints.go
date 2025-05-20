package valclient

import (
	"fmt"
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
	url := c.BuildUrl("https://shared.{shard}.a.pvp.net/content-service/v3/content")
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
	url := c.BuildUrl("https://pd.{shard}.a.pvp.net/account-xp/v1/players/{puuid}")
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
	url := c.BuildUrl("https://pd.{shard}.a.pvp.net/personalization/v3/players/{puuid}/playerloadout")
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
	url := c.BuildUrl("https://pd.{shard}.a.pvp.net/personalization/v3/players/{puuid}/playerloadout")
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
	url := c.BuildUrl("https://pd.{shard}.a.pvp.net/mmr/v1/players/{puuid}")
	ownedItems := new(GetPlayerMmrResponse)

	if err := c.RunRequest(http.MethodGet, url, nil, ownedItems); err != nil {
		return nil, err
	}

	return ownedItems, nil
}

type GetMatchHistoryResponse struct {
	Subject    string `json:"Subject"`
	BeginIndex int    `json:"BeginIndex"`
	EndIndex   int    `json:"EndIndex"`
	Total      int    `json:"Total"`
	History    []struct {
		MatchID       string  `json:"MatchID"`
		GameStartTime int     `json:"GameStartTime"`
		QueueID       QueueID `json:"QueueID"`
	} `json:"History"`
}

/*
All parameters are optional, default values are
- startIndex: 0
- endIndex: 20
- queue: (not passed, results in all queues)
*/
func (c *ValClient) GetMatchHistory(startIndex, endIndex int, queue QueueID) (*GetMatchHistoryResponse, error) {
	if endIndex == 0 {
		endIndex = 20
	}
	additionalParams := []string{
		"{startIndex}", fmt.Sprint(startIndex),
		"{endIndex}", fmt.Sprint(endIndex),
	}

	baseUrl := "https://pd.{shard}.a.pvp.net/match-history/v1/history/{puuid}?startIndex={startIndex}&endIndex={endIndex}"

	if queue != "" {
		baseUrl += "&queue={queue}"
		additionalParams = append(additionalParams, []string{
			"{queue}", string(queue),
		}...)
	}

	url := c.BuildUrl(baseUrl, additionalParams...)
	matchHistory := new(GetMatchHistoryResponse)

	if err := c.RunRequest(http.MethodGet, url, nil, matchHistory); err != nil {
		return nil, err
	}

	return matchHistory, nil
}

type GetMatchDetailsResponse struct {
	MatchInfo struct {
		MatchID                     string             `json:"matchId"`
		MapID                       string             `json:"mapId"`
		GamePodID                   string             `json:"gamePodId"`
		GameLoopZone                string             `json:"gameLoopZone"`
		GameServerAddress           string             `json:"gameServerAddress"`
		GameVersion                 string             `json:"gameVersion"`
		GameLengthMillis            int                `json:"gameLengthMillis,omitempty"`
		GameStartMillis             int                `json:"gameStartMillis"`
		ProvisioningFlowID          ProvisioningFlowID `json:"provisioningFlowID"`
		IsCompleted                 bool               `json:"isCompleted"`
		CustomGameName              string             `json:"customGameName"`
		ForcePostProcessing         bool               `json:"forcePostProcessing"`
		QueueID                     QueueID            `json:"queueID"`
		GameMode                    string             `json:"gameMode"`
		IsRanked                    bool               `json:"isRanked"`
		IsMatchSampled              bool               `json:"isMatchSampled"`
		SeasonID                    string             `json:"seasonId"`
		CompletionState             CompletionState    `json:"completionState"`
		PlatformType                PlatformType       `json:"platformType"`
		PremierMatchInfo            struct{}           `json:"premierMatchInfo"`
		PartyRRPenalties            map[string]int     `json:"partyRRPenalties"`
		ShouldMatchDisablePenalties bool               `json:"shouldMatchDisablePenalties"`
	} `json:"matchInfo"`
	Players []struct {
		Subject      string `json:"subject"`
		GameName     string `json:"gameName"`
		TagLine      string `json:"tagLine"`
		PlatformInfo struct {
			PlatformType      PlatformType `json:"platformType"`
			PlatformOS        string       `json:"platformOS"`
			PlatformOSVersion string       `json:"platformOSVersion"`
			PlatformChipset   string       `json:"platformChipset"`
		} `json:"platformInfo"`
		TeamID      TeamID `json:"teamId"`
		PartyID     string `json:"partyId"`
		CharacterID string `json:"characterId"`
		Stats       *struct {
			Score          int `json:"score"`
			RoundsPlayed   int `json:"roundsPlayed"`
			Kills          int `json:"kills"`
			Deaths         int `json:"deaths"`
			Assists        int `json:"assists"`
			PlaytimeMillis int `json:"playtimeMillis"`
			AbilityCasts   *struct {
				GrenadeCasts  int `json:"grenadeCasts"`
				Ability1Casts int `json:"ability1Casts"`
				Ability2Casts int `json:"ability2Casts"`
				UltimateCasts int `json:"ultimateCasts"`
			} `json:"abilityCasts"`
		} `json:"stats"`
		RoundDamage []struct {
			Round    int    `json:"round"`
			Receiver string `json:"receiver"`
			Damage   int    `json:"damage"`
		} `json:"roundDamage"`
		CompetitiveTier        int    `json:"competitiveTier"`
		IsObserver             bool   `json:"isObserver"`
		PlayerCard             string `json:"playerCard"`
		PlayerTitle            string `json:"playerTitle"`
		PreferredLevelBorder   string `json:"preferredLevelBorder,omitempty"`
		AccountLevel           int    `json:"accountLevel"`
		SessionPlaytimeMinutes int    `json:"sessionPlaytimeMinutes,omitempty"`
		XPModifications        []struct {
			Value float64 `json:"Value"`
			ID    string  `json:"ID"`
		} `json:"xpModifications"`
		BehaviorFactors *struct {
			AFKRounds                   int     `json:"afkRounds"`
			Collisions                  float64 `json:"collisions,omitempty"`
			CommsRatingRecovery         int     `json:"commsRatingRecovery"`
			DamageParticipationOutgoing int     `json:"damageParticipationOutgoing"`
			FriendlyFireIncoming        int     `json:"friendlyFireIncoming,omitempty"`
			FriendlyFireOutgoing        int     `json:"friendlyFireOutgoing,omitempty"`
			MouseMovement               int     `json:"mouseMovement,omitempty"`
			StayedInSpawnRounds         int     `json:"stayedInSpawnRounds,omitempty"`
		} `json:"behaviorFactors"`
		NewPlayerExperienceDetails *struct {
			BasicMovement struct {
				IdleTimeMillis              int `json:"idleTimeMillis"`
				ObjectiveCompleteTimeMillis int `json:"objectiveCompleteTimeMillis"`
			} `json:"basicMovement"`
			BasicGunSkill struct {
				IdleTimeMillis              int `json:"idleTimeMillis"`
				ObjectiveCompleteTimeMillis int `json:"objectiveCompleteTimeMillis"`
			} `json:"basicGunSkill"`
			AdaptiveBots struct {
				AdaptiveBotAverageDurationMillisAllAttempts  int `json:"adaptiveBotAverageDurationMillisAllAttempts"`
				AdaptiveBotAverageDurationMillisFirstAttempt int `json:"adaptiveBotAverageDurationMillisFirstAttempt"`
				KillDetailsFirstAttempt                      any `json:"killDetailsFirstAttempt"`
				IdleTimeMillis                               int `json:"idleTimeMillis"`
				ObjectiveCompleteTimeMillis                  int `json:"objectiveCompleteTimeMillis"`
			} `json:"adaptiveBots"`
			Ability struct {
				IdleTimeMillis              int `json:"idleTimeMillis"`
				ObjectiveCompleteTimeMillis int `json:"objectiveCompleteTimeMillis"`
			} `json:"ability"`
			BombPlant struct {
				IdleTimeMillis              int `json:"idleTimeMillis"`
				ObjectiveCompleteTimeMillis int `json:"objectiveCompleteTimeMillis"`
			} `json:"bombPlant"`
			DefendBombSite struct {
				Success                     bool `json:"success"`
				IdleTimeMillis              int  `json:"idleTimeMillis"`
				ObjectiveCompleteTimeMillis int  `json:"objectiveCompleteTimeMillis"`
			} `json:"defendBombSite"`
			SettingStatus struct {
				IsMouseSensitivityDefault bool `json:"isMouseSensitivityDefault"`
				IsCrosshairDefault        bool `json:"isCrosshairDefault"`
			} `json:"settingStatus"`
			VersionString string `json:"versionString"`
		} `json:"newPlayerExperienceDetails"`
	} `json:"players"`
	Bots    []any `json:"bots"`
	Coaches []struct {
		Subject string `json:"subject"`
		TeamID  TeamID `json:"teamId"`
	} `json:"coaches"`
	Teams []struct {
		TeamID       TeamID `json:"teamId"`
		Won          bool   `json:"won"`
		RoundsPlayed int    `json:"roundsPlayed"`
		RoundsWon    int    `json:"roundsWon"`
		NumPoints    int    `json:"numPoints"`
	} `json:"teams"`
	RoundResults []struct {
		RoundNum             int           `json:"roundNum"`
		RoundResult          RoundResult   `json:"roundResult"`
		RoundCeremony        RoundCeremony `json:"roundCeremony"`
		WinningTeam          TeamID        `json:"winningTeam"`
		BombPlanter          string        `json:"bombPlanter,omitempty"`
		BombDefuser          TeamID        `json:"bombDefuse,omitempty"`
		PlantRoundTime       int           `json:"plantRoundTim,omitempty"`
		PlantPlayerLocations []struct {
			Subject     string  `json:"subject"`
			ViewRadians float64 `json:"viewRadians"`
			Location    struct {
				X int `json:"x"`
				Y int `json:"y"`
			} `json:"location"`
		} `json:"plantPlayerLocations"`
		PlantLocation struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"plantLocation"`
		PlantSite             PlantSite `json:"plantSite"`
		DefuseRoundTime       *int      `json:"defuseRoundTime"`
		DefusePlayerLocations []struct {
			Subject     string  `json:"subject"`
			ViewRadians float64 `json:"viewRadians"`
			Location    struct {
				X int `json:"x"`
				Y int `json:"y"`
			} `json:"location"`
		} `json:"defusePlayerLocations"`
		DefuseLocation struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"defuseLocation"`
		PlayerStats []struct {
			Subject string `json:"subject"`
			Kills   []struct {
				GameTime       int    `json:"gameTime"`
				RoundTime      int    `json:"roundTime"`
				Killer         string `json:"killer"`
				Victim         string `json:"victim"`
				VictimLocation struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"victimLocation"`
				Assistants      []string `json:"assistants"`
				PlayerLocations []struct {
					Subject     string  `json:"subject"`
					ViewRadians float64 `json:"viewRadians"`
					Location    struct {
						X int `json:"x"`
						Y int `json:"y"`
					} `json:"location"`
				} `json:"playerLocations"`
				FinishingDamage struct {
					DamageType          DamageType `json:"damageType"`
					DamageItem          DamageItem `json:"damageItem"`
					IsSecondaryFireMode bool       `json:"isSecondaryFireMode"`
				} `json:"finishingDamage"`
			} `json:"kills"`
			Damage []struct {
				Receiver  string `json:"receiver"`
				Damage    int    `json:"damage"`
				Legshots  int    `json:"legshots"`
				Bodyshots int    `json:"bodyshots"`
				Headshots int    `json:"headshots"`
			} `json:"damage"`
			Score   int `json:"score"`
			Economy struct {
				LoadoutValue int    `json:"loadoutValue"`
				Weapon       string `json:"weapon"`
				Armor        string `json:"armor"`
				Remaining    int    `json:"remaining"`
				Spent        int    `json:"spent"`
			} `json:"economy"`
			Ability struct {
				GrenadeEffects  any `json:"grenadeEffects"`
				Ability1Effects any `json:"ability1Effects"`
				Ability2Effects any `json:"ability2Effects"`
				UltimateEffects any `json:"ultimateEffects"`
			} `json:"ability"`
			WasAfk        bool `json:"wasAfk"`
			WasPenalized  bool `json:"wasPenalized"`
			StayedInSpawn bool `json:"stayedInSpawn"`
		} `json:"playerStats"`
		RoundResultCode RoundResultCode `json:"roundResultCode"`
		PlayerEconomies []struct {
			Subject      string `json:"subject"`
			LoadoutValue int    `json:"loadoutValue"`
			Weapon       string `json:"weapon"`
			Armor        string `json:"armor"`
			Remaining    int    `json:"remaining"`
			Spent        int    `json:"spent"`
		} `json:"playerEconomies"`
		PlayerScores []struct {
			Subject string `json:"subject"`
			Score   int    `json:"score"`
		} `json:"playerScores"`
	} `json:"roundResults"`
	Kills []struct {
		GameTime       int    `json:"gameTime"`
		RoundTime      int    `json:"roundTime"`
		Killer         string `json:"killer"`
		Victim         string `json:"victim"`
		VictimLocation struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"victimLocation"`
		Assistants      []string `json:"assistants"`
		PlayerLocations []struct {
			Subject     string  `json:"subject"`
			ViewRadians float64 `json:"viewRadians"`
			Location    struct {
				X int `json:"x"`
				Y int `json:"y"`
			} `json:"location"`
		} `json:"playerLocations"`
		FinishingDamage struct {
			DamageType          DamageType `json:"damageType"`
			DamageItem          DamageItem `json:"damageItem"`
			IsSecondaryFireMode bool       `json:"isSecondaryFireMode"`
		} `json:"finishingDamage"`
		Round int `json:"round"`
	} `json:"kills"`
}

func (c *ValClient) GetMatchDetails(matchId string) (*GetMatchDetailsResponse, error) {
	url := c.BuildUrl("https://pd.{shard}.a.pvp.net/match-details/v1/matches/{matchID}", "{matchID}", matchId)
	matchDetails := new(GetMatchDetailsResponse)

	if err := c.RunRequest(http.MethodGet, url, nil, matchDetails); err != nil {
		return nil, err
	}

	return matchDetails, nil
}

type GetCompetitiveUpdatesResponse struct {
	Version int    `json:"Version"`
	Subject string `json:"Subject"`
	Matches []struct {
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
	} `json:"Matches"`
}

/*
All parameters are optional, default values are
- startIndex: 0
- endIndex: 20
- queue: (not passed, results in all queues)
*/
func (c *ValClient) GetCompetitiveUpdates(startIndex, endIndex int, queue QueueID) (*GetCompetitiveUpdatesResponse, error) {
	if endIndex == 0 {
		endIndex = 20
	}
	additionalParams := []string{
		"{startIndex}", fmt.Sprint(startIndex),
		"{endIndex}", fmt.Sprint(endIndex),
	}

	baseUrl := "https://pd.{shard}.a.pvp.net/mmr/v1/players/{puuid}/competitiveupdates?startIndex={startIndex}&endIndex={endIndex}&queue={queue}"

	if queue != "" {
		baseUrl += "&queue={queue}"
		additionalParams = append(additionalParams, []string{
			"{queue}", string(queue),
		}...)
	}

	url := c.BuildUrl(baseUrl, additionalParams...)
	compUpdates := new(GetCompetitiveUpdatesResponse)

	if err := c.RunRequest(http.MethodGet, url, nil, compUpdates); err != nil {
		return nil, err
	}

	return compUpdates, nil
}

type GetLeaderboardResponse struct {
	Deployment string `json:"Deployment"`
	QueueID    string `json:"QueueID"`
	SeasonID   string `json:"SeasonID"`
	Players    []struct {
		PlayerCardID    string `json:"PlayerCardID"`
		TitleID         string `json:"TitleID"`
		IsBanned        bool   `json:"IsBanned"`
		IsAnonymized    bool   `json:"IsAnonymized"`
		Puuid           string `json:"puuid"`
		GameName        string `json:"gameName"`
		TagLine         string `json:"tagLine"`
		LeaderboardRank int    `json:"leaderboardRank"`
		RankedRating    int    `json:"rankedRating"`
		NumberOfWins    int    `json:"numberOfWins"`
		CompetitiveTier int    `json:"competitiveTier"`
	} `json:"Players"`
	TotalPlayers          int `json:"totalPlayers"`
	ImmortalStartingPage  int `json:"immortalStartingPage"`
	ImmortalStartingIndex int `json:"immortalStartingIndex"`
	TopTierRRThreshold    int `json:"topTierRRThreshold"`
	TierDetails           map[string]struct {
		RankedRatingThreshold int `json:"rankedRatingThreshold"`
		StartingPage          int `json:"startingPage"`
		StartingIndex         int `json:"startingIndex"`
	} `json:"tierDetails"`
	StartIndex int    `json:"startIndex"`
	Query      string `json:"query"`
}

/*
seasonId is mandatory, others are optional. default values are:
- startIndex: 0
- size: 510 (amount of entries in the leaderboard to return)
- query: (not passed, returns all players. otherwise a player name can be passed)
*/
func (c *ValClient) GetLeaderboard(region Region, startIndex int, seasonId string, size int, query string) (*GetLeaderboardResponse, error) {
	if size == 0 {
		size = 510
	}
	additionalParams := []string{
		"{startIndex}", fmt.Sprint(startIndex),
		"{seasonId}", seasonId,
		"{size}", fmt.Sprint(size),
	}

	baseUrl := "https://pd.{shard}.a.pvp.net/mmr/v1/leaderboards/affinity/{region}/queue/competitive/season/{seasonId}?startIndex={startIndex}&size={size}"

	if query != "" {
		baseUrl += "&query={query}"
		additionalParams = append(additionalParams, []string{
			"{query}", query,
		}...)
	}

	// pretty hacky but whatever, it works
	oldRegion, oldShard := c.Region, c.Shard
	c.Region = region
	c.Shard = ShardForRegion[region]
	url := c.BuildUrl(baseUrl, additionalParams...)
	c.Region = oldRegion
	c.Shard = oldShard

	leaderboard := new(GetLeaderboardResponse)

	if err := c.RunRequest(http.MethodGet, url, nil, leaderboard); err != nil {
		return nil, err
	}

	return leaderboard, nil
}

type GetConfigResponse struct {
	LastApplication string `json:"LastApplication"`
	Collapsed       struct {
		AresMocEntitlement                          string `json:"ARES_MOC_ENTITLEMENT"`
		ClientIconsEnabled                          string `json:"CLIENT.ICONS.ENABLED"`
		ClientLeaderboardsEnabled                   string `json:"CLIENT_LEADERBOARDS_ENABLED"`
		GameAllowConsole                            string `json:"GAME_ALLOW_CONSOLE"`
		GameAllowDeveloperMenu                      string `json:"GAME_ALLOW_DEVELOPER_MENU"`
		GameDisabledDeathcam                        string `json:"GAME_DISABLED_DEATHCAM"`
		GameDisabledSkinsWeapons                    string `json:"GAME_DISABLED_SKINS_WEAPONS"`
		GamePerfreportingEnabled                    string `json:"GAME_PERFREPORTING_ENABLED"`
		GameRemoteMoveInterpEnabled                 string `json:"GAME_REMOTE_MOVE_INTERP_ENABLED"`
		GameRoamingSettingsEnabled                  string `json:"GAME_ROAMINGSETTINGS_ENABLED"`
		GameRoamingSettingsKey                      string `json:"GAME_ROAMINGSETTINGS_KEY"`
		GameRoamingSettingsStorageUrl               string `json:"GAME_ROAMINGSETTINGS_STORAGEURL"`
		MapPreloadingEnabled                        string `json:"MAP_PRELOADING_ENABLED"`
		NamecheckPlatformRegion                     string `json:"NAMECHECK_PLATFORM_REGION"`
		NamecheckPlatformUrl                        string `json:"NAMECHECK_PLATFORM_URL"`
		RosterRealm                                 string `json:"ROSTER_REALM"`
		SecurityWatermarkEnabled                    string `json:"SECURITY_WATERMARK_ENABLED"`
		SecurityWatermarkMaxOpacity                 string `json:"SECURITY_WATERMARK_MAX_OPACITY"`
		SecurityWatermarkMinOpacity                 string `json:"SECURITY_WATERMARK_MIN_OPACITY"`
		SecurityWatermarkTilingFactor               string `json:"SECURITY_WATERMARK_TILING_FACTOR"`
		ServiceUrlAccountXp                         string `json:"SERVICEURL_ACCOUNT_XP"`
		ServiceUrlAggStats                          string `json:"SERVICEURL_AGGSTATS"`
		ServiceUrlAvs                               string `json:"SERVICEURL_AVS"`
		ServiceUrlContent                           string `json:"SERVICEURL_CONTENT"`
		ServiceUrlContracts                         string `json:"SERVICEURL_CONTRACTS"`
		ServiceUrlContractDefinitions               string `json:"SERVICEURL_CONTRACT_DEFINITIONS"`
		ServiceUrlCoreGame                          string `json:"SERVICEURL_COREGAME"`
		ServiceUrlDailyTicket                       string `json:"SERVICEURL_DAILY_TICKET"`
		ServiceUrlFavorites                         string `json:"SERVICEURL_FAVORITES"`
		ServiceUrlGalbsQuery                        string `json:"SERVICEURL_GALBS_QUERY"`
		ServiceUrlLatency                           string `json:"SERVICEURL_LATENCY"`
		ServiceUrlLoginQueue                        string `json:"SERVICEURL_LOGINQUEUE"`
		ServiceUrlMassRewards                       string `json:"SERVICEURL_MASS_REWARDS"`
		ServiceUrlMatchDetails                      string `json:"SERVICEURL_MATCHDETAILS"`
		ServiceUrlMatchHistory                      string `json:"SERVICEURL_MATCHHISTORY"`
		ServiceUrlMatchmaking                       string `json:"SERVICEURL_MATCHMAKING"`
		ServiceUrlMmr                               string `json:"SERVICEURL_MMR"`
		ServiceUrlName                              string `json:"SERVICEURL_NAME"`
		ServiceUrlParty                             string `json:"SERVICEURL_PARTY"`
		ServiceUrlPatchNotes                        string `json:"SERVICEURL_PATCHNOTES"`
		ServiceUrlPersonalization                   string `json:"SERVICEURL_PERSONALIZATION"`
		ServiceUrlPlayerFeedback                    string `json:"SERVICEURL_PLAYERFEEDBACK"`
		ServiceUrlPreGame                           string `json:"SERVICEURL_PREGAME"`
		ServiceUrlPremier                           string `json:"SERVICEURL_PREMIER"`
		ServiceUrlProgression                       string `json:"SERVICEURL_PROGRESSION"`
		ServiceUrlPurchaseMerchant                  string `json:"SERVICEURL_PURCHASEMERCHANT"`
		ServiceUrlReplayCatalog                     string `json:"SERVICEURL_REPLAY_CATALOG"`
		ServiceUrlRestrictions                      string `json:"SERVICEURL_RESTRICTIONS"`
		ServiceUrlSession                           string `json:"SERVICEURL_SESSION"`
		ServiceUrlStore                             string `json:"SERVICEURL_STORE"`
		ServiceUrlTournaments                       string `json:"SERVICEURL_TOURNAMENTS"`
		ServiceTickerMessage                        string `json:"SERVICE_TICKER_MESSAGE"`
		ServiceTickerMessageDeDe                    string `json:"SERVICE_TICKER_MESSAGE.de-DE"`
		ServiceTickerMessageEsMx                    string `json:"SERVICE_TICKER_MESSAGE.es-MX"`
		ServiceTickerMessageFrFr                    string `json:"SERVICE_TICKER_MESSAGE.fr-FR"`
		ServiceTickerMessageItIt                    string `json:"SERVICE_TICKER_MESSAGE.it-IT"`
		ServiceTickerMessagePlPl                    string `json:"SERVICE_TICKER_MESSAGE.pl-PL"`
		ServiceTickerMessagePtBr                    string `json:"SERVICE_TICKER_MESSAGE.pt-BR"`
		ServiceTickerMessageRuRu                    string `json:"SERVICE_TICKER_MESSAGE.ru-RU"`
		ServiceTickerMessageTrTr                    string `json:"SERVICE_TICKER_MESSAGE.tr-TR"`
		ServiceTickerSeverity                       string `json:"SERVICE_TICKER_SEVERITY"`
		StoreScreenOfferRefreshMaxDelayMilliseconds string `json:"STORESCREEN_OFFERREFRESH_MAXDELAY_MILLISECONDS"`
		AvsEnabled                                  string `json:"avs.enabled"`
		CapLocation                                 string `json:"cap.location"`
		CharacterSelectDebugWidgetsHide             string `json:"characterselect.debugwidgets.hide"`
		ChatMutedWordsEnabled                       string `json:"chat.mutedwords.enabled"`
		ChatV3Enabled                               string `json:"chat.v3.enabled"`
		CollectionCharactersEnabled                 string `json:"collection.characters.enabled"`
		CompetitiveSeasonOffsetEndTime              string `json:"competitiveSeasonOffsetEndTime"`
		ConfigClientTelemetrySampleRate             string `json:"config.client.telemetry.samplerate"`
		ContentFilterEnabled                        string `json:"content.filter.enabled"`
		ContentMapsDisabled                         string `json:"content.maps.disabled"`
		EogWip                                      string `json:"eog.wip"`
		FriendsEnabled                              string `json:"friends.enabled"`
		GameUmgChatEnabled                          string `json:"game.umgchat.enabled"`
		HomescreenFeaturedQueues                    string `json:"homescreen.featuredQueues"`
		HomescreenPatchNotesBaseUrl                 string `json:"homescreen.patchnotes.baseURL"`
		HomescreenPromoEnabled                      string `json:"homescreen.promo.enabled"`
		HomescreenPromoKey                          string `json:"homescreen.promo.key"`
		HomescreenWebTileBaseUrl                    string `json:"homescreen.webtile.baseURL"`
		LoginQueueRegion                            string `json:"loginqueue.region"`
		MainMenuBarCollectionsEnabled               string `json:"mainmenubar.collections.enabled"`
		MainMenuBarDebugEnabled                     string `json:"mainmenubar.debug.enabled"`
		MainMenuBarProfileEnabled                   string `json:"mainmenubar.profile.enabled"`
		MainMenuBarProgressionEnabled               string `json:"mainmenubar.progression.enabled"`
		MainMenuBarShootingRangeEnabled             string `json:"mainmenubar.shootingrange.enabled"`
		MainMenuBarStoreEnabled                     string `json:"mainmenubar.store.enabled"`
		MatchDetailsDelay                           string `json:"match.details.delay"`
		NotificationsEnabled                        string `json:"notifications.enabled"`
		PartiesAutoBalanceEnabled                   string `json:"parties.auto.balance.enabled"`
		PartyObserversEnabled                       string `json:"party.observers.enabled"`
		PartyInvitesEnabled                         string `json:"partyinvites.enabled"`
		PatchAvailabilityEnabled                    string `json:"patchavailability.enabled"`
		PersonalizationEquipAnyLevelEnabled         string `json:"personalization.equipAnyLevel.enabled"`
		PersonalizationUseWidePlayerIdentityV2      string `json:"personalization.useWidePlayerIdentityV2"`
		PingUpdateInterval                          string `json:"ping.update.interval"`
		PingUseGamePodsFromParties                  string `json:"ping.useGamePodsFromParties"`
		PlatformFaultedLevel                        string `json:"platformFaulted.level"`
		PlayerFeedbackToolAccessUrl                 string `json:"playerfeedbacktool.accessurl"`
		PlayerFeedbackToolLocale                    string `json:"playerfeedbacktool.locale"`
		PlayerFeedbackToolShard                     string `json:"playerfeedbacktool.shard"`
		PlayerFeedbackToolShow                      string `json:"playerfeedbacktool.show"`
		PlayerFeedbackToolSurveyRequestRateFloat    string `json:"playerfeedbacktool.survey_request_rate_float"`
		PlayScreenPartyWidgetEnabled                string `json:"playscreen.partywidget.enabled"`
		PlayScreenPartyWidgetMatchmakingMaxSize     string `json:"playscreen.partywidget.matchmaking.maxsize"`
		PlayScreenPremierEnabled                    string `json:"playscreen.premier.enabled"`
		PremierConferencesFetchEnabled              string `json:"premier.conferences.fetch.enabled"`
		PremierLeaderboardTabEnabled                string `json:"premier.leaderboardTab.enabled"`
		PremierMatchHistoryTabEnabled               string `json:"premier.matchHistoryTab.enabled"`
		PremierPlayScreenFlowEnabled                string `json:"premier.playscreenflow.enabled"`
		PremierRosterEligibilityCheckEnabled        string `json:"premier.rosterEligibilityCheck.enabled"`
		PremierSeasonsActiveSeasonEnabled           string `json:"premier.seasons.activeseason.enabled"`
		PremierSeasonsFetchEnabled                  string `json:"premier.seasons.fetch.enabled"`
		QueueStatusEnabled                          string `json:"queue.status.enabled"`
		RChatInGameEnabled                          string `json:"rchat.ingame.enabled"`
		ReporterFeedbackFetchEnabled                string `json:"reporterfeedback.fetch.enabled"`
		ReporterFeedbackNotificationsEnabled        string `json:"reporterfeedback.notifications.enabled"`
		RestrictionsV2FetchEnabled                  string `json:"restrictions.v2.fetch.enabled"`
		RestrictionsV2WarningsEnabled               string `json:"restrictions.v2.warnings.enabled"`
		RiotWarningFetchEnabled                     string `json:"riotwarning.fetch.enabled"`
		RiotWarningNotificationsEnabled             string `json:"riotwarning.notifications.enabled"`
		RNetUseAuthenticatedVoice                   string `json:"rnet.useAuthenticatedVoice"`
		RussiaVoiceEnabled                          string `json:"russia.voice.enabled"`
		SettingsLiveDiagnosticsAllowedPlayers       string `json:"settings.livediagnostics.allowedplayers"`
		ShootingTestEnabled                         string `json:"shootingtest.enabled"`
		SkillRatingEnabled                          string `json:"skillrating.enabled"`
		SkillRatingInGameEnabled                    string `json:"skillrating.inGame.enabled"`
		SkillRatingPreGameEnabled                   string `json:"skillrating.preGame.enabled"`
		SocialPanelV6Enabled                        string `json:"social.panel.v6.enabled"`
		SocialViewControllerEnabled                 string `json:"socialviewcontroller.enabled"`
		SocialViewControllerV2Enabled               string `json:"socialviewcontroller.v2.enabled"`
		StoreIsXgpDisabled                          string `json:"store.isXgpDisabled"`
		StoreUseCurrencyInventoryModels             string `json:"store.use_currency_inventory_models"`
		StoreUsePlatformBundleDiscountedPrices      string `json:"store.use_platform_bundle_discounted_prices"`
		TelemetryRtpEventEndpoint                   string `json:"telemetry.rtp.eventendpoint"`
		TelemetryRtpRfc190Scope                     string `json:"telemetry.rtp.rfc190scope"`
		TempVoiceAllowMuting                        string `json:"temp.voice.allowmuting"`
		TournamentsEnabled                          string `json:"tournaments.enabled"`
		TournamentsPreGameEnabled                   string `json:"tournaments.pregame.enabled"`
		VanguardAccessUrl                           string `json:"vanguard.accessurl"`
		VanguardNetRequired                         string `json:"vanguard.netrequired"`
		VoiceClutchMuteEnabled                      string `json:"voice.clutchmute.enabled"`
		VoiceClutchMutePromptEnabled                string `json:"voice.clutchmute.prompt.enabled"`
		VoiceProvider                               string `json:"voice.provider"`
		WhisperEnabled                              string `json:"whisper.enabled"`
	} `json:"Collapsed"`
}

func (c *ValClient) GetConfig() (*GetConfigResponse, error) {
	url := c.BuildUrl("https://shared.{shard}.a.pvp.net/v1/config/{region}")
	config := new(GetConfigResponse)

	if err := c.RunRequest(http.MethodGet, url, nil, config); err != nil {
		return nil, err
	}

	return config, nil
}

type GetNamesResponse []*struct {
	DisplayName string `json:"DisplayName"`
	Subject     string `json:"Subject"`
	GameName    string `json:"GameName"`
	TagLine     string `json:"TagLine"`
}

func (c *ValClient) GetNames(puuids []string) (GetNamesResponse, error) {
	url := c.BuildUrl("https://pd.{shard}.a.pvp.net/name-service/v2/players")
	names := new(GetNamesResponse)

	if err := c.RunRequest(http.MethodPut, url, puuids, names); err != nil {
		return nil, err
	}

	return *names, nil
}
