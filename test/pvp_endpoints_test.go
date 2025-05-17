package test

import (
	"testing"

	"github.com/truearken/valclient/valclient"
)

func TestGetPlayerLoadout(t *testing.T) {
	client, err := valclient.NewClient()
	if err != nil {
		t.Fatalf("unable to create client: %v", err)
	}

	loadout, err := client.GetPlayerLoadout()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if loadout.Identity == nil {
		t.Fatal("expected Identity not to be null")
	}
	if len(loadout.Guns) == 0 {
		t.Fatal("expected Guns not to be empty")
	}
	if len(loadout.ActiveExpressions) == 0 {
		t.Fatal("expected ActiveExpressions not to be empty")
	}

	if loadout.Subject != client.Player.Uuid {
		t.Errorf("Expected Subject %s, got %s", client.Player.Uuid, loadout.Subject)
	}
}

func TestSetPlayerLoadout(t *testing.T) {
	client, err := valclient.NewClient()
	if err != nil {
		t.Fatalf("unable to create client: %v", err)
	}

	loadout, err := client.GetPlayerLoadout()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	loadoutDiff, err := client.SetPlayerLoadout(&valclient.SetPlayerLoadoutRequest{
		Guns:              loadout.Guns,
		ActiveExpressions: loadout.ActiveExpressions,
		Identity:          loadout.Identity,
		Incognito:         !loadout.Incognito,
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if loadout.Incognito == loadoutDiff.Incognito {
		t.Fatal("expected loadout.Incognito to be different then loadoutDiff.Incognito")
	}

	loadoutDiff, err = client.SetPlayerLoadout(&valclient.SetPlayerLoadoutRequest{
		Guns:              loadout.Guns,
		ActiveExpressions: loadout.ActiveExpressions,
		Identity:          loadout.Identity,
		Incognito:         loadout.Incognito,
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if loadout.Incognito != loadoutDiff.Incognito {
		t.Fatal("expected loadout.Incognito to be loadoutDiff.Incognito")
	}
}

func TestGetPlayerMmr(t *testing.T) {
	client, err := valclient.NewClient()
	if err != nil {
		t.Fatalf("unable to create client: %v", err)
	}

	playerMmr, err := client.GetPlayerMmr()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(playerMmr.QueueSkills) == 0 {
		t.Fatalf("expected QueueSkills not to be empty")
	}
}

func TestGetContent(t *testing.T) {
	client, err := valclient.NewClient()
	if err != nil {
		t.Fatalf("unable to create client: %v", err)
	}

	content, err := client.GetContent()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(content.Seasons) == 0 {
		t.Fatalf("expected Seasons not to be empty")
	}

	if len(content.Events) == 0 {
		t.Fatalf("expected Events not to be empty")
	}
}

func TestGetAccountXp(t *testing.T) {
	client, err := valclient.NewClient()
	if err != nil {
		t.Fatalf("unable to create client: %v", err)
	}

	accountXp, err := client.GetAccountXp()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(accountXp.History) == 0 {
		t.Fatalf("expected History not to be empty")
	}

	if accountXp.Progress.Level == 0 {
		t.Fatalf("expected Level not to be 0")
	}
}

func TestGetMatchHistory(t *testing.T) {
	client, err := valclient.NewClient()
	if err != nil {
		t.Fatalf("unable to create client: %v", err)
	}

	history, err := client.GetMatchHistory(0, 0, valclient.QUEUE_ALL)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(history.History) == 0 {
		t.Fatalf("expected all History not to be empty. make sure you have at least one game played")
	}

	history, err = client.GetMatchHistory(0, 0, valclient.QUEUE_DEATHMATCH)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(history.History) == 0 {
		t.Fatalf("expected deathmatch History not to be empty. make sure you have at least one deathmatch game played")
	}

	if history.History[0].QueueID != valclient.QUEUE_DEATHMATCH {
		t.Fatalf("expected QueueID to be QUEUE_DEATHMATCH")
	}
}

func TestGetMatchDetails(t *testing.T) {
	client, err := valclient.NewClient()
	if err != nil {
		t.Fatalf("unable to create client: %v", err)
	}

	history, err := client.GetMatchHistory(0, 0, valclient.QUEUE_ALL)
	if err != nil {
		t.Fatalf("expected no error when getting history, got: %v", err)
	}

	if len(history.History) == 0 {
		t.Fatalf("expected all History not to be empty. make sure you have at least one game played")
	}

	matchDetails, err := client.GetMatchDetails(history.History[0].MatchID)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if matchDetails.MatchInfo.MatchID == "" {
		t.Fatalf("expected matchId not to be empty")
	}
}

func TestGetCompetitiveUpdates(t *testing.T) {
	client, err := valclient.NewClient()
	if err != nil {
		t.Fatalf("unable to create client: %v", err)
	}

	compUpdates, err := client.GetCompetitiveUpdates(0, 0, valclient.QUEUE_COMPETITIVE)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(compUpdates.Matches) == 0 {
		t.Fatalf("expected compUpdates not to be empty. make sure you have at least one competitive game played")
	}
}

func TestGetLeaderboard(t *testing.T) {
	client, err := valclient.NewClient()
	if err != nil {
		t.Fatalf("unable to create client: %v", err)
	}

	oldShard := client.Shard

	v25act2 := "16118998-4705-5813-86dd-0292a2439d90"
	playerName := "arkeN" // it's me !! :)

	leaderbaord, err := client.GetLeaderboard(valclient.SHARD_EU, 0, v25act2, 0, playerName)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if client.Shard != oldShard {
		t.Fatalf("expected shard to remain the same")
	}

	if len(leaderbaord.Players) == 0 {
		t.Fatalf("expected players not to be empty")
	}

	firstMatch := leaderbaord.Players[0].GameName
	if firstMatch != playerName {
		t.Fatalf("expected player to be arkeN, got: %s", firstMatch)
	}
}
