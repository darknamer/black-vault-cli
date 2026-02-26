package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/darknamer/black-vault-lib"
	"github.com/spf13/cobra"
)

var installGitCmd = &cobra.Command{
	Use:   "install-git",
	Short: "Prepare directory for portable Git and optionally set git_path",
	Long:  "Creates ~/.blackvault/tools/git. If you download portable Git there, blackvault will use it when system git is not in PATH.",
	RunE:  runInstallGit,
}

func runInstallGit(cmd *cobra.Command, args []string) error {
	dir := blackvault.PortableGitDir()
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	fmt.Printf("Created: %s\n", dir)
	fmt.Println()
	switch runtime.GOOS {
	case "windows":
		fmt.Println("Portable Git (Windows):")
		fmt.Println("  1. Download from https://git-scm.com/download/win (e.g. 64-bit Portable) or")
		fmt.Println("     https://github.com/git-for-windows/git/releases")
		fmt.Println("  2. Extract to the folder above so that git.exe is at:")
		fmt.Printf("     %s\n", filepath.Join(dir, "cmd", "git.exe"))
		fmt.Println("     or")
		fmt.Printf("     %s\n", filepath.Join(dir, "bin", "git.exe"))
	default:
		fmt.Println("Portable Git (Linux/macOS):")
		fmt.Println("  1. Install git via your package manager (e.g. apt install git, brew install git), or")
		fmt.Println("  2. Download a prebuilt binary and place it at:")
		fmt.Printf("     %s\n", filepath.Join(dir, "bin", "git"))
	}
	fmt.Println()
	fmt.Println("Then run: blackvault config set git_path " + dir)
	fmt.Println("Or leave git_path empty — blackvault will auto-detect system git or the portable one above.")
	return nil
}
