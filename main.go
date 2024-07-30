package main

import (
	"github.com/phantompunk/gupi/cmd"
	store "github.com/phantompunk/gupi/internal"
	"github.com/spf13/afero"
)

var editor *store.Editor

func main() {
	fileSystem := store.NewFileStore("$HOME/.gupi/templates", afero.NewOsFs())
	editor = store.NewEditor(fileSystem)
	cmd.Execute()
}
