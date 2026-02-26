// Package cmd — คำสั่ง CLI สำหรับ git pull ใน workspace
package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/darknamer/black-vault-lib"
	"github.com/spf13/cobra"
)

var (
	gitPullRemote string
	gitPullBranch string
	gitPullRebase bool
)

// gitPullCmd — run git pull to fetch and merge/rebase changes from remote
var gitPullCmd = &cobra.Command{
	Use:   "git-pull [group/repo]",
	Short: "Run git pull in workspace",
	Long:  "Run git pull inside the workspace of the specified repo. You can specify remote/branch and the --rebase option.",
	Args:  cobra.ExactArgs(1),
	RunE:  runGitPull,
}

func init() {
	gitPullCmd.Flags().StringVarP(&gitPullRemote, "remote", "r", "", "Remote name (e.g. origin).")
	gitPullCmd.Flags().StringVarP(&gitPullBranch, "branch", "b", "", "Branch name to pull (e.g. main).")
	gitPullCmd.Flags().BoolVar(&gitPullRebase, "rebase", false, "Use --rebase instead of the default merge.")
}

func runGitPull(cmd *cobra.Command, args []string) error {
	repoPath := args[0]
	svc, err := blackvault.NewService()
	if err != nil {
		return err
	}
	ctx := context.Background()
	out, err := svc.GitPull(ctx, repoPath, gitPullRemote, gitPullBranch, gitPullRebase)
	if err != nil {
		if strings.TrimSpace(out) != "" {
			fmt.Println(out)
		}
		return err
	}
	fmt.Print(out)
	return nil
}

