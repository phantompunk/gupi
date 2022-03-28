package command

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var deleteUsage = `Removes a specific templates from the saved directory.

Usage: brief delete TEMPLATE

Options:
`

func NewDeleteCommand() *Command {
	cmd := &Command{
		flags: flag.NewFlagSet("delete", flag.ExitOnError),
		Execute: func(cmd *Command, args []string) {
			if len(args) == 0 {
				errAndExit("Template name is required")
			}
			file_name := args[0]

			homeDir, err := os.UserHomeDir()
			if err != nil {
				errAndExit("Failed to read home directory")
			}

			filePath := filepath.Join(homeDir, ".gupi", "template", file_name)
			if _, err := os.Stat(filePath); err == nil {
				os.Remove(filePath)
				fmt.Printf("gupi: Template '%s' was deleted\n", file_name)
			}
		},
	}

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, deleteUsage)
	}

	return cmd
}
