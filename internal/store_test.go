package store_test

import (
	"os"
	"testing"

	store "github.com/phantompunk/gupi/internal"
	"github.com/spf13/afero"
)

func TestNewFileStore(t *testing.T) {
	testFS := afero.NewMemMapFs()

	store.NewFileStore("src/templates", testFS)
	_, err := testFS.Stat("src/templates")
	if os.IsNotExist(err) {
		t.Errorf("base dir '%s' does not exist", "src/templates")
	}
}

func TestCreateTemplate(t *testing.T) {
	testFS := setupFS()

	filestore := store.NewFileStore("src/templates", testFS)
	filestore.CreateTemplate("test")

	_, err := testFS.Stat("src/templates/test")
	if os.IsNotExist(err) {
		t.Errorf("file '%s' does not exist", "test")
	}
}

func TestCreateTemplateFromTemplate(t *testing.T) {
	testFS := setupFS()
	afero.WriteFile(testFS, "src/test_template", []byte("Sample Template"), 0644)

	fileStore := store.NewFileStore("src/templates", testFS)
	fileStore.CreateTemplateFromFile("test", "src/test_template")

	name := "src/templates/test"
	if _, err := testFS.Stat(name); err != nil {
		t.Errorf("file '%s' does not exist", name)
	}

	test, err := afero.ReadFile(testFS, name)
	if err != nil {
	}

	actual := string(test)
	if actual != "Sample Template" {
		t.Errorf("Actual '%s', expected '%s'", actual, "Sample Template")
	}
}

func TestDeleteTemplate(t *testing.T) {
	testFS := setupFS()
	afero.WriteFile(testFS, "src/templates/test", []byte("Test Template"), 0644)

	fileStore := store.NewFileStore("src/templates", testFS)
	err := fileStore.DeleteTemplate("test")
	if err != nil {
		t.Fatal(err)
	}

	name := "src/templates/test"
	if _, err := testFS.Stat(name); err == nil {
		t.Errorf("file \"%s\" still exists", name)
	}
}

func TestListTemplates(t *testing.T) {
	testFS := afero.NewMemMapFs()
	testFS.MkdirAll("templates", 0755)
	afero.WriteFile(testFS, "templates/test", []byte("Test Template"), 0644)

	fileStore := store.NewFileStore("", testFS)
	filelist, err := fileStore.ListTemplates()
	if err != nil {
		t.Fatal(err)
	}

	r, _ := testFS.Open("templates")
	rfiles, _ := r.Readdir(0)

	actual := len(filelist)
	expected := len(rfiles)
	if actual != expected {
		t.Errorf("actual %d expected %d", actual, expected)
	}
}

func setupFS() afero.Fs {
	testFS := afero.NewMemMapFs()
	testFS.MkdirAll("src/templates", 0755)
	return testFS
}
