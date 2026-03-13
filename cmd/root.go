package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "blackvault",
	Short: "Git BlackVault — Open what you need. Burn the rest.",
	Long:  "Ephemeral workspace manager for Git repos: open (clone), close (delete), status.",
}

func RootCmd() *cobra.Command {
	return rootCmd
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(openCmd)
	rootCmd.AddCommand(closeCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(installGitCmd)
	rootCmd.AddCommand(gitCommitCmd)
	rootCmd.AddCommand(gitAddCmd)
	rootCmd.AddCommand(gitFetchCmd)
	rootCmd.AddCommand(gitPullCmd)
	rootCmd.AddCommand(gitPushCmd)
	rootCmd.AddCommand(gitBranchCreateCmd)
	rootCmd.AddCommand(gitBranchSwitchCmd)
	rootCmd.AddCommand(gitBranchRenameCmd)
	rootCmd.AddCommand(gitBranchDeleteCmd)
	rootCmd.AddCommand(gitBranchSetUpstreamCmd)
	rootCmd.AddCommand(gitMergeCmd)
	rootCmd.AddCommand(gitRemoteCmd)
	rootCmd.AddCommand(gitFlowCmd)
}
