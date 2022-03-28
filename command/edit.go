package command

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var editUsage = `Edits an existing template.

Usage: brief edit TEMPLATE

Options:
`

func NewEditCommand() *Command {
	cmd := &Command{
		flags: flag.NewFlagSet("edit", flag.ExitOnError),
		Execute: func(cmd *Command, args []string) {
			if len(args) == 0 {
				errAndExit("Template name required")
			}

			homeDir, err := os.UserHomeDir()
			if err != nil {
				errAndExit("Failed to read home directory")
			}

			file_name := args[0]
			file_path := filepath.Join(homeDir, ".gupi", "template", file_name)

			if _, err := os.Stat(file_path); err == nil {
				command := exec.Command("vim", file_path)
				command.Stdout = os.Stdout
				command.Stdin = os.Stdin
				command.Stderr = os.Stderr
				err := command.Run()
				if err != nil {
					os.Exit(1)
				}
			}
			fmt.Printf("gupi: Template '%s' was edited\n", file_name)
		},
	}

	cmd.flags.Usage = func() {
		fmt.Fprint(os.Stderr, editUsage)
	}

	return cmd
}
