package test

import (
	"testing"

	"github.com/truearken/valclient/valclient"
)

func TestGetOwnedItems(t *testing.T) {
	t.Run("agents", func(t *testing.T) {
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
	})

	t.Run("buddies", func(t *testing.T) {
		ownedItems, err := client.GetOwnedItems(valclient.ITEM_TYPE_GUN_BUDDIES)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}

		if ownedItems.ItemTypeID != valclient.ITEM_TYPE_GUN_BUDDIES {
			t.Fatal("expected ItemTypeID to be ITEM_TYPE_GUN_BUDDIES")
		}

		if len(ownedItems.Entitlements) == 0 {
			t.Fatal("expected Entitlements not to be empty")
		}

		if ownedItems.Entitlements[0].InstanceID == nil {
			t.Fatal("expected InstanceID not to be empty")
		}
	})
}
