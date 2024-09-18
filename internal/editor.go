package store

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/olekukonko/tablewriter"
)

type Editor struct {
	store Store
}

func NewEditor(s Store) *Editor {
	return &Editor{store: s}
}

func (e *Editor) New(fileName, filePath, templateName string, useNextWeek bool) error {
	funcMap := initTplFunctions()
	data := getTemplateData(useNextWeek)
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

	newFilePath := filepath.Join(filePath, fileName)
	absFilePath, err := filepath.Abs(newFilePath)
	if err != nil {
		return err
	}
	fmt.Printf("Created '%s' in '%s'\n", fileName, absFilePath)
	return nil
}

func (e *Editor) Create(templateName, pathToTemplate string, useSampleTemplate bool) error {
	if pathToTemplate == "" && !useSampleTemplate {
		return e.Edit(templateName)
	}

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

	fileList := [][]string{}
	for _, file := range files {
		fileSize := humanize.Bytes(uint64(file.Size()))
		updatedAt := humanize.Time(file.ModTime())
		fileList = append(fileList, []string{file.Name(), fileSize, updatedAt})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Size", "Updated"})

	for _, v := range fileList {
		table.Append(v)
	}
	table.Render() // Send output

	return nil
}

func (e *Editor) openWithEditor(templateName string) error {
	templatePath := e.store.GetPathToTemplate(templateName)
	textEditor := findTextEditor()

	command := exec.Command(textEditor, templatePath)
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		return err
	}
	return nil
}

func findTextEditor() string {
	if isCommandAvailable("nvim") {
		return "nvim"
	} else if isCommandAvailable("vim") {
		return "vim"
	} else if isCommandAvailable("nano") {
		return "nano"
	} else if isCommandAvailable("editor") {
		return "editor"
	} else {
		return "vi"
	}
}

func isCommandAvailable(name string) bool {
	cmd := exec.Command("command", "-v", name)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
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

func getTemplateData(useNextWeek bool) any {
	now := time.Now()
	if useNextWeek {
		now = now.AddDate(0, 0, 7)
	}

	year, week := now.ISOWeek()

	date := now.Local().Format("Y.m.d")

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
		Year, Week                        int
		Date                              string
		Mon, Tue, Wed, Thu, Fri, Sat, Sun string
	}{
		Year: year,
		Week: week,
		Date: date,
		Mon:  monday,
		Tue:  tuesday,
		Wed:  wednesday,
		Thu:  thursday,
		Fri:  friday,
		Sat:  saturday,
		Sun:  sunday,
	}
}
