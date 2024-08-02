package store

import (
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

type Store interface {
	CreateTemplate(templateName, pathToTemplate string, useSample ...bool) error
	DeleteTemplate(templateName string) error
	GetPathToTemplate(templateName string) string
	ListTemplates() ([]os.FileInfo, error)
	CreateFile(fileName, filePath string) (afero.File, error)
}

type FileStore struct {
	basePath   string
	fileSystem afero.Fs
}

func NewFileStore(basePath string, fileSystem afero.Fs) *FileStore {
	if _, err := fileSystem.Stat(basePath); os.IsNotExist(err) {
		fileSystem.MkdirAll(basePath, os.ModePerm)
	}
	return &FileStore{basePath: basePath, fileSystem: fileSystem}
}

func (fstore *FileStore) CreateFile(fileName, filePath string) (afero.File, error) {
	outPath := filepath.Join(filePath, fileName)
	file, err := fstore.fileSystem.Create(outPath)
	if err != nil {
	}
	return file, nil
}

func (fstore *FileStore) CreateTemplate(templateName, pathToTemplate string, useSample ...bool) error {
	// Create an empty template
	templatePath := fstore.GetPathToTemplate(templateName)
	file, err := fstore.fileSystem.Create(templatePath)
	if err != nil {
		return err
	}

	// Create template from url
	if isURLPath(pathToTemplate) {
		res, err := http.Get(pathToTemplate)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		body, err := afero.ReadAll(res.Body)
		if err != nil {
			return err
		}

		_, err = file.WriteString(string(body))
		if err != nil {
			return err
		}
		return nil
	}

	// Create template from file
	if isFilePath(pathToTemplate) {
		data, err := afero.ReadFile(fstore.fileSystem, pathToTemplate)
		if err != nil {
			return err
		}

		_, err = file.WriteString(string(data))
		if err != nil {
			return err
		}
	}

	// Create sample template
	if useSampleTemplate(useSample) {
		if len(useSample) > 0 && useSample[0] {
			err := fstore.createSampleTemplate(file)
			if err != nil {
				return err
			}
			return nil
		}
	}

	return nil
}

func (fstore *FileStore) DeleteTemplate(templateName string) error {
	templatePath := fstore.GetPathToTemplate(templateName)
	err := fstore.fileSystem.Remove(templatePath)
	if err != nil {
		return err
	}
	return nil
}

func (fstore *FileStore) EditTemplate(templateName string) error {
	return nil
}

func (fstore *FileStore) ListTemplates() ([]os.FileInfo, error) {
	fsys, err := fstore.fileSystem.Open(fstore.basePath)
	if err != nil {
		return nil, err
	}

	filelist, err := fsys.Readdir(0)
	if err != nil {
		return nil, err
	}

	return filelist, nil
}

func (fstore *FileStore) GetPathToTemplate(templateName string) string {
	return filepath.Join(fstore.basePath, templateName)
}

func (fstore *FileStore) createSampleTemplate(file afero.File) error {
	const sampleTemplate string = `# 🗓 8.1 Thur`
	_, err := file.WriteString(string(sampleTemplate))
	if err != nil {
		return err
	}
	return nil
}

func isURLPath(urlPath string) bool {
	_, err := url.ParseRequestURI(urlPath)
	return err == nil
}

func isFilePath(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	}
	return false
}

func useSampleTemplate(bools []bool) bool {
	anyTrue := false
	for _, b := range bools {
		anyTrue = anyTrue && b
	}
	return anyTrue
}
