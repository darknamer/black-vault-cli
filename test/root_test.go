package test

import (
	"testing"

	"github.com/darknamer/black-vault-cli/cmd"
)

func TestRootCmd_Metadata(t *testing.T) {
	rootCmd := cmd.RootCmd()

	if rootCmd.Use != "blackvault" {
		t.Fatalf("expected Use=%q, got %q", "blackvault", rootCmd.Use)
	}

	if rootCmd.Short == "" {
		t.Fatalf("expected Short to be non-empty")
	}

	if rootCmd.Long == "" {
		t.Fatalf("expected Long to be non-empty")
	}
}

