package store

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
	"time"
)

type Editor struct {
	store Store
}

func NewEditor(s Store) *Editor {
	return &Editor{store: s}
}

func (e *Editor) New(fileName, filePath, templateName string) error {
	funcMap := initTplFunctions()
	data := getTemplateData()
	templatePath := e.store.GetPathToTemplate(templateName)
	fileTemplate, err := template.New(templateName).Funcs(funcMap).ParseFiles(templatePath)
	if err != nil {
		return err
	}

	file, err := e.store.CreateFile(fileName, filePath)
	if err != nil {
		return err
	}

	err = fileTemplate.Execute(file, data)
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

func initTplFunctions() template.FuncMap {
	return template.FuncMap{
		"Week": Week,
	}
}

func Week() int {
	now := time.Now()
	_, week := now.ISOWeek()
	return week
}

var days = map[int]int{
	0: -1,
	1: 0,
	2: 1,
	3: 2,
	4: 3,
	5: 4,
	6: -2,
}

func getTemplateData() any {
	now := time.Now()
	year, week := now.ISOWeek()

	firstDayOfWeek := now.AddDate(0, 0, -days[int(now.Weekday())])
	_, m, d := firstDayOfWeek.Date()
	monday := fmt.Sprintf("%d.%d", m, d)

	_, m, d = firstDayOfWeek.AddDate(0, 0, 1).Date()
	tuesday := fmt.Sprintf("%d.%d", m, d)

	_, m, d = firstDayOfWeek.AddDate(0, 0, 2).Date()
	wednesday := fmt.Sprintf("%d.%d", m, d)

	_, m, d = firstDayOfWeek.AddDate(0, 0, 3).Date()
	thursday := fmt.Sprintf("%d.%d", m, d)

	_, m, d = firstDayOfWeek.AddDate(0, 0, 4).Date()
	friday := fmt.Sprintf("%d.%d", m, d)

	_, m, d = firstDayOfWeek.AddDate(0, 0, 5).Date()
	saturday := fmt.Sprintf("%d.%d", m, d)

	_, m, d = firstDayOfWeek.AddDate(0, 0, 6).Date()
	sunday := fmt.Sprintf("%d.%d", m, d)

	return struct {
		Year, Week int
		Mon, Tue, Wed, Thu, Fri, Sat, Sun string
	}{
		Year: year,
		Week: week,
		Mon: monday,
		Tue: tuesday,
		Wed: wednesday,
		Thu: thursday,
		Fri: friday,
		Sat: saturday,
		Sun: sunday,
	}
}
