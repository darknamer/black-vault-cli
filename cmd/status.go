package cmd

import (
	"fmt"

	"github.com/darknamer/black-vault-lib"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show workspace status (active / closed)",
	RunE:  runStatus,
}

func runStatus(cmd *cobra.Command, args []string) error {
	svc, err := blackvault.NewService()
	if err != nil {
		return err
	}
	entries, err := svc.Status()
	if err != nil {
		return err
	}
	fmt.Println("ACTIVE:")
	for _, e := range entries {
		if e.State == "active" {
			dirty := ""
			if e.Dirty {
				dirty = " (dirty)"
			}
			fmt.Printf("  - %s%s\n", e.RepoPath, dirty)
		}
	}
	if len(entries) == 0 {
		fmt.Println("  (none)")
	}
	fmt.Println("CLOSED:")
	// แสดงเฉพาะ active จาก disk; "closed" คือที่ไม่อยู่ในรายการ
	fmt.Println("  (repos not in ACTIVE)")
	return nil
}
