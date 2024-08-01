package store_test

import (
	"os"
	"path/filepath"
	"testing"

	store "github.com/phantompunk/gupi/internal"
	"github.com/spf13/afero"
)

var (
	testFS    afero.Fs
	testStore store.Store
	baseDir   string = "templates"
)

func init() {
	testFS = afero.NewMemMapFs()
	testStore = store.NewFileStore(baseDir, testFS)
}

func TestNewFileStore(t *testing.T) {
	_, err := testFS.Stat(baseDir)
	if os.IsNotExist(err) {
		t.Errorf("base dir '%s' does not exist", baseDir)
	}
}

func TestCreateSampleTemplate(t *testing.T) {
	templateName := "sample"
	templatePath := filepath.Join(baseDir, templateName)
	err := testStore.CreateTemplate(templateName, "", true)
	if err != nil {
		t.Errorf("failed to create sample template '%s'", templateName)
	}

	_, err = testFS.Stat(templatePath)
	if err != nil {
		t.Error("sample template not found")
	}
}

func TestCreateEmptyTemplate(t *testing.T) {
	templateName := "test"
	templatePath := filepath.Join(baseDir, templateName)
	err := testStore.CreateTemplate(templateName, "")
	if err != nil {
		t.Errorf("failed to create template '%s'", templateName)
	}

	_, err = testFS.Stat(templatePath)
	if os.IsNotExist(err) {
		t.Errorf("file '%s' does not exist", templateName)
	}
}

func TestCreateTemplateFromTemplate(t *testing.T) {
	templateName := "test_template"
	templatePath := filepath.Join(baseDir, templateName)
	createTestFile(t, templatePath, []byte("Sample Template"))

	err := testStore.CreateTemplate("test", templatePath)
	if err != nil {
		t.Errorf("failed to create template '%s'", templateName)
	}

	if _, err := testFS.Stat(templatePath); err != nil {
		t.Errorf("file '%s' does not exist", templateName)
	}

	test, err := afero.ReadFile(testFS, templatePath)
	if err != nil {
		t.Errorf("failed reading file '%s'", templatePath)
	}

	actual := string(test)
	if actual != "Sample Template" {
		t.Errorf("Actual '%s', expected '%s'", actual, "Sample Template")
	}
}

func TestDeleteTemplate(t *testing.T) {
	templateName := "test"
	templatePath := filepath.Join(baseDir, templateName)
	createTestFile(t, templatePath, []byte("Test Template"))

	err := testStore.DeleteTemplate(templateName)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := testFS.Stat(templatePath); err == nil {
		t.Errorf("file \"%s\" still exists", templateName)
	}
}

func TestListTemplates(t *testing.T) {
	templateName := "test"
	templatePath := filepath.Join(baseDir, templateName)
	createTestFile(t, templatePath, []byte("Test Template"))

	filelist, err := testStore.ListTemplates()
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

func createTestFile(t *testing.T, templatePath string, data []byte) {
	t.Helper()
	err := afero.WriteFile(testFS, templatePath, data, 0644)
	if err != nil {
		t.Errorf("error creating test file '%s'", err.Error())
	}
}
