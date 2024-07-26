package gupi

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/phantompunk/gupi/gupi/fs"
)

func AddTemplate(tmplName, pathToFile, urlToFile string) error {
	if len(tmplName) == 0 {
		return errors.New("Template name required")
	}

	if len(pathToFile) > 0 && len(urlToFile) > 0 {
		return errors.New("Use a filepath or url only, not both")
	}

	if len(pathToFile) == 0 && len(urlToFile) == 0 {
		err := createEmptyTemplate(tmplName)
		if err != nil {
			return err
		}
	}
	return nil
}

func createEmptyTemplate(tmplName string) error {
	// Create a file
	err := fs.CreateFile(tmplName)
	if err != nil {
		return err
	}
	// Open file with editor
	err = openWithEditor(tmplName)
	if err != nil {
		return err
	}
	return nil
}

func openWithEditor(file string) error {
	fullPath, err := fs.GetTemplatePath(file)
	if err != nil {
		return err
	}

	fmt.Print("Opening in vim")
	command := exec.Command("vim", fullPath)
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	err = command.Run()
	if err != nil {
		return err
	}
	return nil
}
