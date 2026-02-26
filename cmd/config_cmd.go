package cmd

import (
	"fmt"

	"github.com/darknamer/black-vault-lib"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show or set config (e.g. git_path)",
	Long:  "config [get [key]] | config set <key> <value>. Keys: git_path. Default: use system git if in PATH, else portable under ~/.blackvault/tools/git.",
}

var configGetCmd = &cobra.Command{
	Use:   "get [key]",
	Short: "Show config (all or one key)",
	RunE:  runConfigGet,
}

var configSetCmd = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "Set config key (e.g. set git_path /path/to/git)",
	RunE:  runConfigSet,
}

func init() {
	configCmd.AddCommand(configGetCmd)
	configCmd.AddCommand(configSetCmd)
}

func runConfigGet(cmd *cobra.Command, args []string) error {
	svc, err := blackvault.NewService()
	if err != nil {
		return err
	}
	resolved := svc.ResolvedGitPath()
	if len(args) == 0 {
		fmt.Printf("git_path (config): %q\n", svc.GetGitPath())
		fmt.Printf("git_path (resolved): %q\n", resolved)
		if resolved == "" {
			fmt.Println("(Using go-git instead — system or portable git not found.)")
		}
		return nil
	}
	switch args[0] {
	case "git_path":
		fmt.Println(svc.GetGitPath())
	case "git_path_resolved":
		fmt.Println(resolved)
	default:
		return fmt.Errorf("unknown key: %s", args[0])
	}
	return nil
}

func runConfigSet(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: blackvault config set <key> <value>")
	}
	key, value := args[0], args[1]
	svc, err := blackvault.NewService()
	if err != nil {
		return err
	}
	switch key {
	case "git_path":
		svc.SetGitPath(value)
		if err := svc.SaveConfig(); err != nil {
			return err
		}
		fmt.Printf("Set git_path = %q\n", value)
		fmt.Printf("Resolved: %q\n", svc.ResolvedGitPath())
	default:
		return fmt.Errorf("unknown key: %s", key)
	}
	return nil
}
