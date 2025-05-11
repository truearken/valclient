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

func TestGetOwnedItems(t *testing.T) {
	client, err := valclient.NewClient()
	if err != nil {
		t.Fatalf("unable to create client: %v", err)
	}

	ownedItems, err := client.GetOwnedItems(valclient.ITEM_TYPE_AGENTS)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if ownedItems.ItemTypeID != valclient.ITEM_TYPE_AGENTS {
		t.Fatal("expected ItemTypeID to be ITEM_TYPE_AGENTS")
	}

	if len(ownedItems.Entitlements) == 0 {
		t.Fatal("expected Entitlements not to be empty")
	}
}

func TestGetContent(t *testing.T) {
	client, err := valclient.NewClient()
	if err != nil {
		t.Fatalf("unable to create client: %v", err)
	}

	content, err := client.GetContentRequest()
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
		t.Fatalf("expected Seasons not to be empty")
	}

	if accountXp.Progress.Level == 0 {
		t.Fatalf("expected Level not to be 0")
	}
}
