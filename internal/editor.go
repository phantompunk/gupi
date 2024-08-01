package store

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

type Editor struct {
	store Store
}

func NewEditor(s Store) *Editor {
	return &Editor{store: s}
}

func (e *Editor) New(fileName, templateName string) error {
	templatePath := e.store.GetPathToTemplate(templateName)
	fileTemplate, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	file, err := e.store.CreateFile(fileName, ".")
	if err != nil {
		return err
	}

	err = fileTemplate.Execute(file, "")
	if err != nil {
		return err
	}

	currDir, _ := os.Getwd()
	fmt.Printf("Created '%s' in '%s'\n", fileName, filepath.Join(currDir, fileName))
	return nil
}

func (e *Editor) Create(templateName, pathToTemplate string, useSampleTemplate bool) error {
	err := e.store.CreateTemplate(templateName, pathToTemplate, useSampleTemplate)
	if err != nil {
		return err
	}
	return nil
}

func (e *Editor) Delete(templateName string) error {
	err := e.store.DeleteTemplate(templateName)
	if err != nil {
		return err
	}
	return nil
}

func (e *Editor) Edit(templateName string) error {
	// templatePath := e.store.GetTemplatePath(templateName)
	err := e.openWithEditor(templateName)
	if err != nil {
		return err
	}
	return nil
}

func (e *Editor) List() error {
	files, err := e.store.ListTemplates()
	if err != nil {
		return errors.New("Unable to read templates")
	}

	fmt.Printf("NAME\t\tSIZE\t\tMODIFIED")
	for _, files := range files {
		fmt.Printf("\n%-15s %-15v %v", files.Name(), files.Size(), files.ModTime().Format("2006-01-02 15:04:05"))
	}

	return nil
}

func (e *Editor) openWithEditor(templateName string) error {
	templatePath := e.store.GetPathToTemplate(templateName)

	fmt.Print("Opening in vim")
	command := exec.Command("vim", templatePath)
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		return err
	}
	return nil
}
