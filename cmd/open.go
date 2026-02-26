package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/darknamer/black-vault-lib"
	"github.com/spf13/cobra"
)

var (
	openShallow  bool
	openBranch   string
	openDetached bool
)

var openCmd = &cobra.Command{
	Use:   "open [group/repo]",
	Short: "Open (clone) a workspace",
	Args:  cobra.ExactArgs(1),
	RunE:  runOpen,
}

func init() {
	openCmd.Flags().BoolVar(&openShallow, "shallow", false, "Shallow clone")
	openCmd.Flags().StringVar(&openBranch, "branch", "", "Branch to clone")
	openCmd.Flags().BoolVar(&openDetached, "detached", false, "Detached / review mode")
	log.SetFlags(0)
}

func runOpen(cmd *cobra.Command, args []string) error {
	repoPath := args[0]
	svc, err := blackvault.NewService()
	if err != nil {
		return err
	}
	ctx := context.Background()
	path, err := svc.Open(ctx, repoPath, blackvault.OpenOptions{
		Shallow:  openShallow,
		Branch:   openBranch,
		Detached: openDetached,
	})
	if err != nil {
		return err
	}
	fmt.Println(path)
	// ถ้าต้องการ: เปิดใน IDE (เช่น code .)
	if os.Getenv("BLACKVAULT_OPEN_IDE") == "1" {
		// TODO: รัน "code" หรือ "cursor" กับ path
		_ = path
	}
	return nil
}

