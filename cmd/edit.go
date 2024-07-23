package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(editCmd)
}

var editUsage = `Edits an existing template.

Usage: gupi edit TEMPLATE

Options:
`

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a template",
	Long:  editUsage,
	Run:   editFunc,
}

var editFunc = func(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		errAndExit("template name required")
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		errAndExit("failed to read home directory")
	}

	file_name := args[0]
	file_path := filepath.Join(homedir, ".gupi", "template", file_name)

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

	cmd.Flags().Usage = func() {
		fmt.Fprint(os.Stderr, editUsage)
	}
}

