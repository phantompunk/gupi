package command

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var filePath string

var addUsage = `Add a template from a file path or URL.

Usage: brief add [OPTIONS] TEMPLATE

Options:
	-f, --file	path to an existing template file
`

var addFunc = func(cmd *Command, args []string) {
	if len(filePath) == 0 {
		errAndExit("File path required")
	}

	if _, err := os.Stat(filePath); err != nil {
		errAndExit("File does not exist")
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		errAndExit("Failed to return user's home directory")
	}

	fileDir := filepath.Join(homeDir, ".gupi", "template")
	if _, err := os.Stat(fileDir); err != nil {
		fmt.Println("Creating template folder at", fileDir)
		os.Mkdir(fileDir, 0755)
	}

	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		errAndExit(fmt.Sprintf("failed to read from %s\n", filePath))
	}

	file_name := filepath.Base(filePath)
	fileOutPath := filepath.Join(fileDir, file_name)
	out, err := os.Create(fileOutPath)
	if err != nil {
		errAndExit("Failed to write file " + fileOutPath)
	}
	defer out.Close()

	out.WriteString(string(f))
	fmt.Printf("gupi: Template '%s' was added\n", file_name)
}

func NewAddCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("add", flag.ExitOnError),
		Execute: addFunc,
	}

	cmd.flags.StringVar(&filePath, "file", "", "")
	cmd.flags.StringVar(&filePath, "f", "", "")

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, addUsage)
	}

	return cmd
}
