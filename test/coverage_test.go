package test

import (
	"testing"

	"github.com/darknamer/black-vault-cli/cmd"
)

// TestCoverage_CollectsCmdPackage ensures the cmd package is exercised so that
// "go test -cover ./..." or "make test-cover" reports coverage for cmd.
// Run coverage with: make test-cover, make test-cover-profile, or make test-cover-html.
func TestCoverage_CollectsCmdPackage(t *testing.T) {
	root := cmd.RootCmd()
	if root == nil {
		t.Fatal("RootCmd() must not return nil")
	}
	if root.Use != "blackvault" {
		t.Errorf("RootCmd().Use = %q, want blackvault", root.Use)
	}
}
