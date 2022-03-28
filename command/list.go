package command

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var listUsage = `List all currently avaible templates.

Usage: brief list
Options:
`

var listFunc = func(cmd *Command, args []string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		errAndExit("Failed to read home directory")
	}

	fileDir := filepath.Join(homeDir, ".gupi", "template")
	file, err := os.Open(fileDir)
	if err != nil {
		errAndExit("Template folder not found")
	}
	defer file.Close()

	filelist, err := file.Readdir(0)
	if err != nil {
		errAndExit("Unable to read file")
	}

	fmt.Printf("NAME\t\tSIZE\t\tMODIFIED")
	for _, files := range filelist {
		fmt.Printf("\n%-15s %-15v %v", files.Name(), files.Size(), files.ModTime().Format("2006-01-02 15:04:05"))
	}
}

func NewListCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("list", flag.ExitOnError),
		Execute: listFunc,
	}

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, listUsage)
	}

	return cmd
}
