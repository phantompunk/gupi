package gupi

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

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

func EditTemplate(tmplName string) error {
	err := openWithEditor(tmplName)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTemplate(tmplName string) error {
	err := fs.DeleteFile(tmplName)
	if err != nil {
		return err
	}
	fmt.Printf("gupi: Template '%s' was deleted\n", tmplName)
	return nil
}

func RenderTemplate(fileName, tmplName string) error {
	path, err := fs.GetTemplatePath(tmplName)
	if err != nil {
		return err
	}

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return err
	}

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	err = tmpl.Execute(f, "")
	if err != nil {
		return err
	}
	currDir, _ := os.Getwd()
	fmt.Printf("Created '%s' in '%s'\n", fileName, filepath.Join(currDir, fileName))
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
