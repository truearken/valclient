package test

import (
	"testing"

	"github.com/truearken/valclient/valclient"
)

func TestGetOwnedItems(t *testing.T) {
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
