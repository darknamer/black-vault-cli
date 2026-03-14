package test

import (
	"testing"

	"github.com/darknamer/black-vault-cli/cmd"
	"github.com/spf13/cobra"
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

// expectedTopLevelCommands must match cmd/root.go init() — all commands added to root.
var expectedTopLevelCommands = []string{
	"open", "close", "status", "serve", "config", "install-git",
	"git-commit", "git-add", "git-fetch", "git-pull", "git-push",
	"git-branch-create", "git-branch-switch", "git-branch-rename",
	"git-branch-delete", "git-branch-set-upstream",
	"git-merge", "git-remote", "git-flow", "version",
}

func TestRootCmd_ExpectedTopLevelCommandsPresent(t *testing.T) {
	root := cmd.RootCmd()
	names := make(map[string]bool)
	for _, c := range root.Commands() {
		names[c.Name()] = true
	}
	for _, name := range expectedTopLevelCommands {
		if !names[name] {
			t.Errorf("root command missing expected subcommand %q", name)
		}
	}
}

// collectCommandPaths returns all command paths (e.g. ["config","get"], ["git-remote","list"]).
func collectCommandPaths(c *cobra.Command, path []string) [][]string {
	var out [][]string
	for _, sub := range c.Commands() {
		p := append(append([]string{}, path...), sub.Name())
		out = append(out, p)
		out = append(out, collectCommandPaths(sub, p)...)
	}
	return out
}

func TestRootCmd_AllCommandsHaveMetadata(t *testing.T) {
	root := cmd.RootCmd()
	paths := collectCommandPaths(root, nil)
	for _, path := range paths {
		c := root
		for _, name := range path {
			var next *cobra.Command
			for _, sub := range c.Commands() {
				if sub.Name() == name {
					next = sub
					break
				}
			}
			if next == nil {
				t.Fatalf("command not found: %v", path)
			}
			c = next
		}
		if c.Use == "" {
			t.Errorf("command %v has empty Use", path)
		}
		if c.Short == "" {
			t.Errorf("command %v has empty Short", path)
		}
	}
}

func TestRootCmd_HelpSucceedsForAllCommands(t *testing.T) {
	root := cmd.RootCmd()
	// Root help
	root.SetArgs([]string{"--help"})
	if err := root.Execute(); err != nil {
		t.Errorf("root --help: %v", err)
	}
	// Each top-level and nested path: path + "--help"
	paths := collectCommandPaths(root, nil)
	for _, path := range paths {
		args := append(append([]string{}, path...), "--help")
		root.SetArgs(args)
		if err := root.Execute(); err != nil {
			t.Errorf("help for %v: %v", path, err)
		}
	}
}

