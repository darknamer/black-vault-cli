package cmd

import (
	"fmt"

	"github.com/darknamer/black-vault-lib"
	"github.com/spf13/cobra"
)

var closeForce bool

var closeCmd = &cobra.Command{
	Use:   "close [group/repo]",
	Short: "Close (delete) a workspace",
	Args:  cobra.ExactArgs(1),
	RunE:  runClose,
}

func init() {
	closeCmd.Flags().BoolVar(&closeForce, "force", false, "Skip dirty check and force delete")
}

func runClose(cmd *cobra.Command, args []string) error {
	repoPath := args[0]
	svc, err := blackvault.NewService()
	if err != nil {
		return err
	}
	// TODO: ถ้า !force ตรวจ dirty แล้วถามยืนยัน
	if err := svc.Close(repoPath, closeForce); err != nil {
		return err
	}
	fmt.Printf("Closed %s\n", repoPath)
	return nil
}
