// Package cmd — คำสั่ง CLI สำหรับแสดงเวอร์ชันของ black-vault-cli และ library
package cmd

import (
	"fmt"

	"github.com/darknamer/black-vault-lib"
	"github.com/spf13/cobra"
)

// cliVersion คือเวอร์ชันของ black-vault-cli (ใช้ semantic versioning)
// TODO: อัปเดตเลขเวอร์ชันเมื่อมีการ release ใหม่
const cliVersion = "0.1.0"

// versionCmd — show versions of the CLI and library
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show versions of black-vault-cli and black-vault-lib",
	RunE:  runVersion,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func runVersion(cmd *cobra.Command, args []string) error {
	fmt.Printf("black-vault-cli: %s\n", cliVersion)
	fmt.Printf("black-vault-lib: %s\n", blackvault.Version())
	return nil
}

