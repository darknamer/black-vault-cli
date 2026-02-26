// Package cmd — คำสั่ง CLI สำหรับ git push ใน workspace
package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/darknamer/black-vault-lib"
	"github.com/spf13/cobra"
)

var (
	gitPushRemote      string
	gitPushBranch      string
	gitPushSetUpstream bool
	gitPushForce       bool
)

// gitPushCmd — run git push to send changes to remote
var gitPushCmd = &cobra.Command{
	Use:   "git-push [group/repo]",
	Short: "Run git push in workspace",
	Long:  "Run git push inside the workspace of the specified repo. You can specify remote/branch, -u (--set-upstream) and --force.",
	Args:  cobra.ExactArgs(1),
	RunE:  runGitPush,
}

func init() {
	gitPushCmd.Flags().StringVarP(&gitPushRemote, "remote", "r", "", "Remote name (e.g. origin). If empty, git's default is used.")
	gitPushCmd.Flags().StringVarP(&gitPushBranch, "branch", "b", "", "Branch name to push (e.g. main).")
	gitPushCmd.Flags().BoolVarP(&gitPushSetUpstream, "set-upstream", "u", false, "Use -u to set upstream (equivalent to git push -u).")
	gitPushCmd.Flags().BoolVar(&gitPushForce, "force", false, "Use --force to push forcibly (use with care).")
}

func runGitPush(cmd *cobra.Command, args []string) error {
	repoPath := args[0]
	svc, err := blackvault.NewService()
	if err != nil {
		return err
	}
	ctx := context.Background()
	out, err := svc.GitPush(ctx, repoPath, gitPushRemote, gitPushBranch, gitPushSetUpstream, gitPushForce)
	if err != nil {
		if strings.TrimSpace(out) != "" {
			fmt.Println(out)
		}
		return err
	}
	fmt.Print(out)
	return nil
}

