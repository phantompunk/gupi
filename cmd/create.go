package cmd

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createUsage = `Usage: gupi create [options...]
Examples:
  # Generate a report for the week containing Feb 2, 2021
	gupi create --date 02/17/2021

Options:
  --template	Path to custom template file for weekly report.
  --date	Date used to generate weekly report. Default is current date.
  --output 	Output directory for newly created report. Default is current directory.
`
var tmp string

var createFunc = func(cmd *cobra.Command, args []string) {
	if len(tmp) == 0 {
		errAndExit("Template required")
	}

	file_name := args[0]

	homeDir, err := os.UserHomeDir()
	if err != nil {
		errAndExit("Failed to return user's home directory")
	}

	path := filepath.Join(homeDir, ".gupi", "template", tmp)
	if _, err := os.Stat(path); err == nil {
		date := time.Now()
		data := getDates(date)

		t, err := template.ParseFiles(path)
		if err != nil {
			errAndExit("Failed to parse template")
		}

		f, err := os.Create(file_name)
		if err != nil {
			errAndExit("Failed to create template instance")
		}

		err = t.Execute(f, data)
		if err != nil {
			errAndExit("Failed to execute template")
		}

		currDir, _ := os.Getwd()
		fmt.Printf("Created '%s' in '%s'\n", file_name, filepath.Join(currDir, file_name))
	}
}

var createCmd = &cobra.Command{
	Use: "create",
	Short: "Create an instance of a template",
	Long: createUsage,
	Run: createFunc,
}


type weekYear struct {
	Week int
	Year int
	Mon  string
	Tue  string
	Wed  string
	Thu  string
	Fri  string
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

func getDates(start time.Time) *weekYear {
	year, week := start.ISOWeek()

	firstDayOfWeek := start.AddDate(0, 0, -days[int(start.Weekday())])
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

	return &weekYear{week, year, monday, tuesday, wednesday, thursday, friday}
}

