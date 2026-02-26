// Package cmd — คำสั่ง CLI สำหรับจัดการ remote (เพิ่ม/ลบ/ตั้ง URL/แสดงรายการ)
package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/darknamer/black-vault-lib"
	"github.com/spf13/cobra"
)

var gitRemoteCmd = &cobra.Command{
	Use:   "git-remote",
	Short: "Manage repo remotes (list / add / remove / set-url).",
	Long:  "List, add, remove or change remote URLs in the specified workspace.",
}

var gitRemoteListCmd = &cobra.Command{
	Use:   "list [group/repo]",
	Short: "List remotes (git remote -v).",
	Args:  cobra.ExactArgs(1),
	RunE:  runGitRemoteList,
}

var gitRemoteAddCmd = &cobra.Command{
	Use:   "add [group/repo] [name] [url]",
	Short: "Add a remote named <name> pointing to <url>.",
	Args:  cobra.ExactArgs(3),
	RunE:  runGitRemoteAdd,
}

var gitRemoteRemoveCmd = &cobra.Command{
	Use:   "remove [group/repo] [name]",
	Short: "Remove the specified remote.",
	Args:  cobra.ExactArgs(2),
	RunE:  runGitRemoteRemove,
}

var gitRemoteSetUrlCmd = &cobra.Command{
	Use:   "set-url [group/repo] [name] [url]",
	Short: "Change the URL of an existing remote.",
	Args:  cobra.ExactArgs(3),
	RunE:  runGitRemoteSetUrl,
}

func init() {
	gitRemoteCmd.AddCommand(gitRemoteListCmd)
	gitRemoteCmd.AddCommand(gitRemoteAddCmd)
	gitRemoteCmd.AddCommand(gitRemoteRemoveCmd)
	gitRemoteCmd.AddCommand(gitRemoteSetUrlCmd)
}

func runGitRemoteList(cmd *cobra.Command, args []string) error {
	repoPath := args[0]
	svc, err := blackvault.NewService()
	if err != nil {
		return err
	}
	ctx := context.Background()
	out, err := svc.GitRemoteList(ctx, repoPath)
	if err != nil {
		if strings.TrimSpace(out) != "" {
			fmt.Println(out)
		}
		return err
	}
	fmt.Print(out)
	return nil
}

func runGitRemoteAdd(cmd *cobra.Command, args []string) error {
	repoPath, name, url := args[0], strings.TrimSpace(args[1]), strings.TrimSpace(args[2])
	if name == "" || url == "" {
		return fmt.Errorf("remote name and url are required")
	}
	svc, err := blackvault.NewService()
	if err != nil {
		return err
	}
	ctx := context.Background()
	out, err := svc.GitRemoteAdd(ctx, repoPath, name, url)
	if err != nil {
		if strings.TrimSpace(out) != "" {
			fmt.Println(out)
		}
		return err
	}
	fmt.Print(out)
	return nil
}

func runGitRemoteRemove(cmd *cobra.Command, args []string) error {
	repoPath, name := args[0], strings.TrimSpace(args[1])
	if name == "" {
		return fmt.Errorf("remote name is required")
	}
	svc, err := blackvault.NewService()
	if err != nil {
		return err
	}
	ctx := context.Background()
	out, err := svc.GitRemoteRemove(ctx, repoPath, name)
	if err != nil {
		if strings.TrimSpace(out) != "" {
			fmt.Println(out)
		}
		return err
	}
	fmt.Print(out)
	return nil
}

func runGitRemoteSetUrl(cmd *cobra.Command, args []string) error {
	repoPath, name, url := args[0], strings.TrimSpace(args[1]), strings.TrimSpace(args[2])
	if name == "" || url == "" {
		return fmt.Errorf("remote name and url are required")
	}
	svc, err := blackvault.NewService()
	if err != nil {
		return err
	}
	ctx := context.Background()
	out, err := svc.GitRemoteSetUrl(ctx, repoPath, name, url)
	if err != nil {
		if strings.TrimSpace(out) != "" {
			fmt.Println(out)
		}
		return err
	}
	fmt.Print(out)
	return nil
}
