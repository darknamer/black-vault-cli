// Package cmd — คำสั่ง CLI สำหรับ git add ใน workspace
package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/darknamer/black-vault-lib"
	"github.com/spf13/cobra"
)

// gitAddCmd — รัน git add ภายใน workspace ของ repo ที่ระบุ
// ถ้าไม่ระบุ path จะใช้ git add -A (เพิ่มทุกไฟล์ที่เปลี่ยนแปลง)
var gitAddCmd = &cobra.Command{
	Use:   "git-add [group/repo] [paths...]",
	Short: "Run git add in workspace (no paths = add all).",
	Long:  "Run git add <paths...> inside the workspace of the specified repo; if no paths are provided, git add -A is used to add all changes.",
	Args:  cobra.MinimumNArgs(1),
	RunE:  runGitAdd,
}

func runGitAdd(cmd *cobra.Command, args []string) error {
	repoPath := args[0]
	var paths []string
	if len(args) > 1 {
		for _, p := range args[1:] {
			p = strings.TrimSpace(p)
			if p != "" {
				paths = append(paths, p)
			}
		}
	}
	svc, err := blackvault.NewService()
	if err != nil {
		return err
	}
	ctx := context.Background()
	out, err := svc.GitAdd(ctx, repoPath, paths...)
	if err != nil {
		if strings.TrimSpace(out) != "" {
			fmt.Println(out)
		}
		return err
	}
	fmt.Print(out)
	return nil
}

