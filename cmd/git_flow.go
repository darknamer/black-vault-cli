// Package cmd — คำสั่ง CLI แบบ git-flow (feature branch เป็นต้น)
package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/darknamer/black-vault-lib"
	"github.com/spf13/cobra"
)

// ตัวเลือกสำหรับ git-flow
var (
	gitFlowInitMainBranch    string
	gitFlowInitDevelopBranch string

	gitFlowReleaseFromBranch    string
	gitFlowReleaseMainBranch    string
	gitFlowReleaseDevelopBranch string
	gitFlowReleaseTag           bool

	gitFlowHotfixFromBranch    string
	gitFlowHotfixMainBranch    string
	gitFlowHotfixDevelopBranch string
	gitFlowHotfixTag           bool
)

// gitFlowCmd — helper command group in git-flow style
var gitFlowCmd = &cobra.Command{
	Use:   "git-flow",
	Short: "Helper commands in git-flow style (feature/release/hotfix).",
}

var gitFlowInitCmd = &cobra.Command{
	Use:   "init [group/repo]",
	Short: "Initialize a simple git-flow layout (main/master + develop).",
	Long:  "Create/prepare a main (or master) and develop branch for the feature/release/hotfix flow.",
	Args:  cobra.ExactArgs(1),
	RunE:  runGitFlowInit,
}

var gitFlowFeatureStartCmd = &cobra.Command{
	Use:   "feature-start [group/repo] [name]",
	Short: "Start a feature branch: create branch feature/<name> and check it out.",
	Args:  cobra.ExactArgs(2),
	RunE:  runGitFlowFeatureStart,
}

var gitFlowReleaseStartCmd = &cobra.Command{
	Use:   "release-start [group/repo] [name]",
	Short: "Start a release branch: create branch release/<name> from develop (or the specified branch).",
	Args:  cobra.ExactArgs(2),
	RunE:  runGitFlowReleaseStart,
}

var gitFlowReleaseFinishCmd = &cobra.Command{
	Use:   "release-finish [group/repo] [name]",
	Short: "Finish a release: merge release/<name> into main and develop (optionally create a tag).",
	Args:  cobra.ExactArgs(2),
	RunE:  runGitFlowReleaseFinish,
}

var gitFlowHotfixStartCmd = &cobra.Command{
	Use:   "hotfix-start [group/repo] [name]",
	Short: "Start a hotfix branch: create branch hotfix/<name> from main (or the specified branch).",
	Args:  cobra.ExactArgs(2),
	RunE:  runGitFlowHotfixStart,
}

var gitFlowHotfixFinishCmd = &cobra.Command{
	Use:   "hotfix-finish [group/repo] [name]",
	Short: "Finish a hotfix: merge hotfix/<name> into main and develop (optionally create a tag).",
	Args:  cobra.ExactArgs(2),
	RunE:  runGitFlowHotfixFinish,
}

func init() {
	gitFlowInitCmd.Flags().StringVar(&gitFlowInitMainBranch, "main", "", "Name of the main branch (default main; if not found, will try master).")
	gitFlowInitCmd.Flags().StringVar(&gitFlowInitDevelopBranch, "develop", "", "Name of the development branch (default develop).")

	gitFlowReleaseStartCmd.Flags().StringVar(&gitFlowReleaseFromBranch, "from", "", "Create the release from this branch (default develop).")
	gitFlowReleaseFinishCmd.Flags().StringVar(&gitFlowReleaseMainBranch, "main", "", "Main branch (default main).")
	gitFlowReleaseFinishCmd.Flags().StringVar(&gitFlowReleaseDevelopBranch, "develop", "", "Development branch (default develop).")
	gitFlowReleaseFinishCmd.Flags().BoolVar(&gitFlowReleaseTag, "tag", false, "Create a tag with the release name after merge.")

	gitFlowHotfixStartCmd.Flags().StringVar(&gitFlowHotfixFromBranch, "from", "", "Create the hotfix from this branch (default main).")
	gitFlowHotfixFinishCmd.Flags().StringVar(&gitFlowHotfixMainBranch, "main", "", "Main branch (default main).")
	gitFlowHotfixFinishCmd.Flags().StringVar(&gitFlowHotfixDevelopBranch, "develop", "", "Development branch (default develop).")
	gitFlowHotfixFinishCmd.Flags().BoolVar(&gitFlowHotfixTag, "tag", false, "Create a tag with the hotfix name after merge.")

	gitFlowCmd.AddCommand(gitFlowInitCmd)
	gitFlowCmd.AddCommand(gitFlowFeatureStartCmd)
	gitFlowCmd.AddCommand(gitFlowReleaseStartCmd)
	gitFlowCmd.AddCommand(gitFlowReleaseFinishCmd)
	gitFlowCmd.AddCommand(gitFlowHotfixStartCmd)
	gitFlowCmd.AddCommand(gitFlowHotfixFinishCmd)
}

func runGitFlowInit(cmd *cobra.Command, args []string) error {
	repoPath := args[0]
	svc, err := blackvault.NewService()
	if err != nil {
		return err
	}
	ctx := context.Background()
	out, err := svc.GitFlowInit(ctx, repoPath, gitFlowInitMainBranch, gitFlowInitDevelopBranch)
	if err != nil {
		if strings.TrimSpace(out) != "" {
			fmt.Println(out)
		}
		return err
	}
	fmt.Print(out)
	return nil
}

func runGitFlowFeatureStart(cmd *cobra.Command, args []string) error {
	repoPath := args[0]
	name := strings.TrimSpace(args[1])
	if name == "" {
		return fmt.Errorf("You must specify the feature name.")
	}
	svc, err := blackvault.NewService()
	if err != nil {
		return err
	}
	ctx := context.Background()
	out, err := svc.GitFlowFeatureStart(ctx, repoPath, name)
	if err != nil {
		if strings.TrimSpace(out) != "" {
			fmt.Println(out)
		}
		return err
	}
	fmt.Print(out)
	return nil
}

func runGitFlowReleaseStart(cmd *cobra.Command, args []string) error {
	repoPath := args[0]
	name := strings.TrimSpace(args[1])
	if name == "" {
		return fmt.Errorf("You must specify the release name.")
	}
	svc, err := blackvault.NewService()
	if err != nil {
		return err
	}
	ctx := context.Background()
	out, err := svc.GitFlowReleaseStart(ctx, repoPath, name, gitFlowReleaseFromBranch)
	if err != nil {
		if strings.TrimSpace(out) != "" {
			fmt.Println(out)
		}
		return err
	}
	fmt.Print(out)
	return nil
}

func runGitFlowReleaseFinish(cmd *cobra.Command, args []string) error {
	repoPath := args[0]
	name := strings.TrimSpace(args[1])
	if name == "" {
		return fmt.Errorf("You must specify the release name.")
	}
	svc, err := blackvault.NewService()
	if err != nil {
		return err
	}
	ctx := context.Background()
	out, err := svc.GitFlowReleaseFinish(ctx, repoPath, name, gitFlowReleaseMainBranch, gitFlowReleaseDevelopBranch, gitFlowReleaseTag)
	if err != nil {
		if strings.TrimSpace(out) != "" {
			fmt.Println(out)
		}
		return err
	}
	fmt.Print(out)
	return nil
}

func runGitFlowHotfixStart(cmd *cobra.Command, args []string) error {
	repoPath := args[0]
	name := strings.TrimSpace(args[1])
	if name == "" {
		return fmt.Errorf("You must specify the hotfix name.")
	}
	svc, err := blackvault.NewService()
	if err != nil {
		return err
	}
	ctx := context.Background()
	out, err := svc.GitFlowHotfixStart(ctx, repoPath, name, gitFlowHotfixFromBranch)
	if err != nil {
		if strings.TrimSpace(out) != "" {
			fmt.Println(out)
		}
		return err
	}
	fmt.Print(out)
	return nil
}

func runGitFlowHotfixFinish(cmd *cobra.Command, args []string) error {
	repoPath := args[0]
	name := strings.TrimSpace(args[1])
	if name == "" {
		return fmt.Errorf("You must specify the hotfix name.")
	}
	svc, err := blackvault.NewService()
	if err != nil {
		return err
	}
	ctx := context.Background()
	out, err := svc.GitFlowHotfixFinish(ctx, repoPath, name, gitFlowHotfixMainBranch, gitFlowHotfixDevelopBranch, gitFlowHotfixTag)
	if err != nil {
		if strings.TrimSpace(out) != "" {
			fmt.Println(out)
		}
		return err
	}
	fmt.Print(out)
	return nil
}

