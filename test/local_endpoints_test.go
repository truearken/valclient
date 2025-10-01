package test

import (
	"testing"
)

func TestGetHelp(t *testing.T) {
	help, err := client.GetHelp()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if help == nil {
		t.Fatal("expected help not to be empty")
	}
}
