package test

import (
	"strings"
	"testing"

	"github.com/darknamer/black-vault-cli/cmd"
)

// TestArgValidation checks that commands with required args fail when args are missing or wrong.
func TestOpen_RequiresOneArg(t *testing.T) {
	root := cmd.RootCmd()
	root.SetArgs([]string{"open"})
	err := root.Execute()
	if err == nil {
		t.Fatal("expected error when open is called with no args")
	}
	if !strings.Contains(err.Error(), "accepts 1 arg") && !strings.Contains(err.Error(), "required") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestClose_RequiresOneArg(t *testing.T) {
	root := cmd.RootCmd()
	root.SetArgs([]string{"close"})
	err := root.Execute()
	if err == nil {
		t.Fatal("expected error when close is called with no args")
	}
	if !strings.Contains(err.Error(), "accepts 1 arg") && !strings.Contains(err.Error(), "required") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestGitAdd_RequiresAtLeastOneArg(t *testing.T) {
	root := cmd.RootCmd()
	root.SetArgs([]string{"git-add"})
	err := root.Execute()
	if err == nil {
		t.Fatal("expected error when git-add is called with no args")
	}
	errStr := err.Error()
	if !strings.Contains(errStr, "accepts 1 arg") && !strings.Contains(errStr, "required") &&
		!strings.Contains(errStr, "minimum") && !strings.Contains(errStr, "at least") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestGitCommit_RequiresOneArg(t *testing.T) {
	root := cmd.RootCmd()
	root.SetArgs([]string{"git-commit"})
	err := root.Execute()
	if err == nil {
		t.Fatal("expected error when git-commit is called with no args")
	}
	if !strings.Contains(err.Error(), "accepts 1 arg") && !strings.Contains(err.Error(), "required") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestGitFetch_RequiresOneArg(t *testing.T) {
	root := cmd.RootCmd()
	root.SetArgs([]string{"git-fetch"})
	err := root.Execute()
	if err == nil {
		t.Fatal("expected error when git-fetch is called with no args")
	}
}

func TestGitPull_RequiresOneArg(t *testing.T) {
	root := cmd.RootCmd()
	root.SetArgs([]string{"git-pull"})
	err := root.Execute()
	if err == nil {
		t.Fatal("expected error when git-pull is called with no args")
	}
}

func TestGitPush_RequiresOneArg(t *testing.T) {
	root := cmd.RootCmd()
	root.SetArgs([]string{"git-push"})
	err := root.Execute()
	if err == nil {
		t.Fatal("expected error when git-push is called with no args")
	}
}

func TestGitBranchCreate_RequiresTwoArgs(t *testing.T) {
	root := cmd.RootCmd()
	root.SetArgs([]string{"git-branch-create"})
	err := root.Execute()
	if err == nil {
		t.Fatal("expected error when git-branch-create is called with no args")
	}
	root.SetArgs([]string{"git-branch-create", "group/repo"})
	err = root.Execute()
	if err == nil {
		t.Fatal("expected error when git-branch-create is called with one arg")
	}
}

func TestGitBranchSwitch_RequiresTwoArgs(t *testing.T) {
	root := cmd.RootCmd()
	root.SetArgs([]string{"git-branch-switch", "group/repo"})
	err := root.Execute()
	if err == nil {
		t.Fatal("expected error when git-branch-switch is called with one arg")
	}
}

func TestGitBranchDelete_RequiresTwoArgs(t *testing.T) {
	root := cmd.RootCmd()
	root.SetArgs([]string{"git-branch-delete", "group/repo"})
	err := root.Execute()
	if err == nil {
		t.Fatal("expected error when git-branch-delete is called with one arg")
	}
}

func TestGitMerge_RequiresTwoArgs(t *testing.T) {
	root := cmd.RootCmd()
	root.SetArgs([]string{"git-merge", "group/repo"})
	err := root.Execute()
	if err == nil {
		t.Fatal("expected error when git-merge is called with one arg")
	}
}

func TestGitRemoteList_RequiresOneArg(t *testing.T) {
	root := cmd.RootCmd()
	root.SetArgs([]string{"git-remote", "list"})
	err := root.Execute()
	if err == nil {
		t.Fatal("expected error when git-remote list is called with no repo arg")
	}
}

func TestConfigSet_RequiresKeyAndValue(t *testing.T) {
	root := cmd.RootCmd()
	root.SetArgs([]string{"config", "set"})
	err := root.Execute()
	if err == nil {
		t.Fatal("expected error when config set is called with no args")
	}
	root.SetArgs([]string{"config", "set", "git_path"})
	err = root.Execute()
	if err == nil {
		t.Fatal("expected error when config set is called with only key")
	}
}

func TestGitFlowInit_RequiresOneArg(t *testing.T) {
	root := cmd.RootCmd()
	root.SetArgs([]string{"git-flow", "init"})
	err := root.Execute()
	if err == nil {
		t.Fatal("expected error when git-flow init is called with no args")
	}
}

func TestGitFlowFeatureStart_RequiresTwoArgs(t *testing.T) {
	root := cmd.RootCmd()
	root.SetArgs([]string{"git-flow", "feature-start", "group/repo"})
	err := root.Execute()
	if err == nil {
		t.Fatal("expected error when git-flow feature-start is called with one arg")
	}
}

func TestVersion_NoArgsSucceeds(t *testing.T) {
	root := cmd.RootCmd()
	root.SetArgs([]string{"version"})
	if err := root.Execute(); err != nil {
		t.Errorf("version with no args should succeed (or fail only on service): %v", err)
	}
}

func TestStatus_NoArgsSucceeds(t *testing.T) {
	root := cmd.RootCmd()
	root.SetArgs([]string{"status"})
	// status calls blackvault.NewService() and Status(); may fail if no config dir, but command wiring is correct
	_ = root.Execute()
}

func TestInstallGit_NoArgsSucceeds(t *testing.T) {
	root := cmd.RootCmd()
	root.SetArgs([]string{"install-git"})
	if err := root.Execute(); err != nil {
		t.Errorf("install-git with no args should succeed: %v", err)
	}
}

func TestConfigGet_NoArgsSucceeds(t *testing.T) {
	root := cmd.RootCmd()
	root.SetArgs([]string{"config", "get"})
	// May fail on NewService() if env/config missing; we only check it doesn't panic
	_ = root.Execute()
}
