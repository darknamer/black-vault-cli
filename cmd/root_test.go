package cmd

import (
	"testing"
)

// TestRootCmd_andExecuteHelp ensures the root command and --help path are
// exercised for coverage when running "go test -cover ./cmd/..." or "make test-cover".
func TestRootCmd_andExecuteHelp(t *testing.T) {
	root := RootCmd()
	if root == nil {
		t.Fatal("RootCmd() must not return nil")
	}
	root.SetArgs([]string{"--help"})
	if err := root.Execute(); err != nil {
		t.Errorf("Execute(--help): %v", err)
	}
}
