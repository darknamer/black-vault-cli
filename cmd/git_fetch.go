// Package cmd — คำสั่ง CLI สำหรับ git fetch ใน workspace
package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/darknamer/black-vault-lib"
	"github.com/spf13/cobra"
)

var (
	gitFetchRemote  string
	gitFetchRefspec string
	gitFetchAll     bool
	gitFetchPrune   bool
)

// gitFetchCmd — run git fetch to get updates from remote
var gitFetchCmd = &cobra.Command{
	Use:   "git-fetch [group/repo]",
	Short: "Run git fetch in workspace",
	Long:  "Run git fetch inside the workspace of the specified repo. You can use --all, --prune or specify remote/refspec.",
	Args:  cobra.ExactArgs(1),
	RunE:  runGitFetch,
}

func init() {
	gitFetchCmd.Flags().StringVarP(&gitFetchRemote, "remote", "r", "", "Remote name (e.g. origin).")
	gitFetchCmd.Flags().StringVar(&gitFetchRefspec, "refspec", "", "Refspec to fetch.")
	gitFetchCmd.Flags().BoolVar(&gitFetchAll, "all", false, "Fetch all remotes (equivalent to git fetch --all).")
	gitFetchCmd.Flags().BoolVar(&gitFetchPrune, "prune", false, "Remove refs that no longer exist on the remote (--prune).")
}

func runGitFetch(cmd *cobra.Command, args []string) error {
	repoPath := args[0]
	svc, err := blackvault.NewService()
	if err != nil {
		return err
	}
	ctx := context.Background()
	out, err := svc.GitFetch(ctx, repoPath, gitFetchRemote, gitFetchRefspec, gitFetchAll, gitFetchPrune)
	if err != nil {
		if strings.TrimSpace(out) != "" {
			fmt.Println(out)
		}
		return err
	}
	fmt.Print(out)
	return nil
}

