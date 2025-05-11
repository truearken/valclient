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
		t.Fatal("expected active expressions not to be empty")
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
		t.Fatalf("expected no error, got %v", err)
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
