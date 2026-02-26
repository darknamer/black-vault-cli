package main

import (
	"os"

	"github.com/darknamer/black-vault-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
