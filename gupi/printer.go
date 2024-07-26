package gupi

import (
	"errors"
	"fmt"

	"github.com/phantompunk/gupi/gupi/fs"
)

func Display() error {
	filelist, err := fs.ReadTemplates()
	if err != nil {
		return errors.New("Unable to read templates")
	}

	fmt.Printf("NAME\t\tSIZE\t\tMODIFIED")
	for _, files := range filelist {
		fmt.Printf("\n%-15s %-15v %v\n", files.Name(), files.Size(), files.ModTime().Format("2006-01-02 15:04:05"))
	}

	return nil
}
