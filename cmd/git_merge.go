// Package cmd — คำสั่ง CLI สำหรับ git merge แบบง่าย
package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/darknamer/black-vault-lib"
	"github.com/spf13/cobra"
)

var (
	gitMergeNoFF    bool
	gitMergeSquash  bool
	gitMergeMessage string
)

// gitMergeCmd — merge one branch into the current branch in the workspace
var gitMergeCmd = &cobra.Command{
	Use:   "git-merge [group/repo] [branch]",
	Short: "Run git merge to merge a branch into the current branch in the workspace.",
	Long:  "Run git merge <branch> inside the specified workspace. Supports --no-ff, --squash and -m for the merge message.",
	Args:  cobra.ExactArgs(2),
	RunE:  runGitMerge,
}

func init() {
	gitMergeCmd.Flags().BoolVar(&gitMergeNoFF, "no-ff", false, "Use --no-ff to force creating a merge commit.")
	gitMergeCmd.Flags().BoolVar(&gitMergeSquash, "squash", false, "Use --squash to squash commits (you must commit manually after the merge).")
	gitMergeCmd.Flags().StringVarP(&gitMergeMessage, "message", "m", "", "Merge message (-m).")
}

func runGitMerge(cmd *cobra.Command, args []string) error {
	repoPath := args[0]
	branch := strings.TrimSpace(args[1])
	if branch == "" {
		return fmt.Errorf("You must specify the branch to merge.")
	}
	svc, err := blackvault.NewService()
	if err != nil {
		return err
	}
	ctx := context.Background()
	out, err := svc.GitMerge(ctx, repoPath, branch, gitMergeNoFF, gitMergeSquash, gitMergeMessage)
	if err != nil {
		if strings.TrimSpace(out) != "" {
			fmt.Println(out)
		}
		return err
	}
	fmt.Print(out)
	return nil
}

