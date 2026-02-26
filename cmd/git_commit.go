// Package cmd — คำสั่ง CLI สำหรับ commit การเปลี่ยนแปลงใน workspace
package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/darknamer/black-vault-lib"
	"github.com/spf13/cobra"
)

var (
	gitCommitMessage string
	gitCommitNoAdd   bool
)

// gitCommitCmd — commit การเปลี่ยนแปลง (ดีฟอลต์ทำ git add -A ก่อน)
var gitCommitCmd = &cobra.Command{
	Use:   "git-commit [group/repo]",
	Short: "Commit changes in the workspace.",
	Long:  "Run git add -A (by default) and then git commit -m <message> in the workspace of the specified repo.",
	Args:  cobra.ExactArgs(1),
	RunE:  runGitCommit,
}

func init() {
	gitCommitCmd.Flags().StringVarP(&gitCommitMessage, "message", "m", "", "Commit message.")
	gitCommitCmd.Flags().BoolVar(&gitCommitNoAdd, "no-add", false, "Do not run git add -A before commit.")
}

func runGitCommit(cmd *cobra.Command, args []string) error {
	repoPath := args[0]
	if strings.TrimSpace(gitCommitMessage) == "" {
		return fmt.Errorf("Commit message is required (use -m or --message).")
	}
	svc, err := blackvault.NewService()
	if err != nil {
		return err
	}
	ctx := context.Background()
	out, err := svc.GitCommit(ctx, repoPath, gitCommitMessage, !gitCommitNoAdd)
	if err != nil {
		// แสดง output จาก git ด้วย เพื่อช่วย debug
		if strings.TrimSpace(out) != "" {
			fmt.Println(out)
		}
		return err
	}
	fmt.Print(out)
	return nil
}

