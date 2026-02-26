// Package cmd — คำสั่ง CLI สำหรับจัดการ branch (สร้าง/สลับ/เปลี่ยนชื่อ/ลบ/ตั้ง upstream)
package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/darknamer/black-vault-lib"
	"github.com/spf13/cobra"
)

var gitBranchForceDelete bool

var gitBranchCreateCmd = &cobra.Command{
	Use:   "git-branch-create [group/repo] [branch]",
	Short: "Create a new branch and check it out.",
	Long:  "Equivalent to git checkout -b <branch> in the specified workspace.",
	Args:  cobra.ExactArgs(2),
	RunE:  runGitBranchCreate,
}

var gitBranchSwitchCmd = &cobra.Command{
	Use:   "git-branch-switch [group/repo] [branch]",
	Short: "Switch to an existing branch (checkout).",
	Args:  cobra.ExactArgs(2),
	RunE:  runGitBranchSwitch,
}

var gitBranchRenameCmd = &cobra.Command{
	Use:   "git-branch-rename [group/repo] [new_name] | [group/repo] [old_name] [new_name]",
	Short: "Rename a branch (current branch, or from old to new).",
	Long:  "2 arguments: rename the current branch to new_name. 3 arguments: rename from old_name to new_name.",
	Args:  cobra.RangeArgs(2, 3),
	RunE:  runGitBranchRename,
}

var gitBranchDeleteCmd = &cobra.Command{
	Use:   "git-branch-delete [group/repo] [branch]",
	Short: "Delete a local branch (does not delete the remote).",
	Long:  "Use --force to delete even if the branch is not merged (-D).",
	Args:  cobra.ExactArgs(2),
	RunE:  runGitBranchDelete,
}

var gitBranchSetUpstreamCmd = &cobra.Command{
	Use:   "git-branch-set-upstream [group/repo] [upstream] [branch]",
	Short: "Set the upstream of a branch to remote/branch (e.g. origin/main).",
	Long:  "If branch is not specified, the current branch is used.",
	Args:  cobra.RangeArgs(2, 3),
	RunE:  runGitBranchSetUpstream,
}

func init() {
	gitBranchDeleteCmd.Flags().BoolVarP(&gitBranchForceDelete, "force", "f", false, "Force delete the branch even if it is not merged (-D).")
}

func runGitBranchCreate(cmd *cobra.Command, args []string) error {
	repoPath := args[0]
	branch := strings.TrimSpace(args[1])
	if branch == "" {
		return fmt.Errorf("branch name is required")
	}
	svc, err := blackvault.NewService()
	if err != nil {
		return err
	}
	ctx := context.Background()
	out, err := svc.GitCreateBranch(ctx, repoPath, branch, true)
	if err != nil {
		if strings.TrimSpace(out) != "" {
			fmt.Println(out)
		}
		return err
	}
	fmt.Print(out)
	return nil
}

func runGitBranchSwitch(cmd *cobra.Command, args []string) error {
	repoPath := args[0]
	branch := strings.TrimSpace(args[1])
	if branch == "" {
		return fmt.Errorf("branch name is required")
	}
	svc, err := blackvault.NewService()
	if err != nil {
		return err
	}
	ctx := context.Background()
	out, err := svc.GitCheckout(ctx, repoPath, branch)
	if err != nil {
		if strings.TrimSpace(out) != "" {
			fmt.Println(out)
		}
		return err
	}
	fmt.Print(out)
	return nil
}

func runGitBranchRename(cmd *cobra.Command, args []string) error {
	repoPath := args[0]
	var oldName, newName string
	if len(args) == 2 {
		newName = strings.TrimSpace(args[1])
	} else {
		oldName = strings.TrimSpace(args[1])
		newName = strings.TrimSpace(args[2])
	}
	if newName == "" {
		return fmt.Errorf("new branch name is required")
	}
	svc, err := blackvault.NewService()
	if err != nil {
		return err
	}
	ctx := context.Background()
	out, err := svc.GitBranchRename(ctx, repoPath, oldName, newName)
	if err != nil {
		if strings.TrimSpace(out) != "" {
			fmt.Println(out)
		}
		return err
	}
	fmt.Print(out)
	return nil
}

func runGitBranchDelete(cmd *cobra.Command, args []string) error {
	repoPath := args[0]
	branch := strings.TrimSpace(args[1])
	if branch == "" {
		return fmt.Errorf("branch name is required")
	}
	svc, err := blackvault.NewService()
	if err != nil {
		return err
	}
	ctx := context.Background()
	out, err := svc.GitBranchDelete(ctx, repoPath, branch, gitBranchForceDelete)
	if err != nil {
		if strings.TrimSpace(out) != "" {
			fmt.Println(out)
		}
		return err
	}
	fmt.Print(out)
	return nil
}

func runGitBranchSetUpstream(cmd *cobra.Command, args []string) error {
	repoPath := args[0]
	upstream := strings.TrimSpace(args[1])
	branch := ""
	if len(args) >= 3 {
		branch = strings.TrimSpace(args[2])
	}
	if upstream == "" {
		return fmt.Errorf("upstream is required (e.g. origin/main)")
	}
	svc, err := blackvault.NewService()
	if err != nil {
		return err
	}
	ctx := context.Background()
	out, err := svc.GitBranchSetUpstream(ctx, repoPath, branch, upstream)
	if err != nil {
		if strings.TrimSpace(out) != "" {
			fmt.Println(out)
		}
		return err
	}
	fmt.Print(out)
	return nil
}
