package store

import (
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

type Store interface {
	CreateTemplate(templateName, pathToTemplate string) error
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
	file, err := fstore.fileSystem.Create(fileName)
	if err != nil {
	}
	return file, nil
}

func (fstore *FileStore) CreateTemplate(templateName, pathToTemplate string) error {
	templatePath := fstore.GetPathToTemplate(templateName)
	file, err := fstore.fileSystem.Create(templatePath)
	if err != nil {
		return err
	}

	if len(pathToTemplate) > 0 {
		data, err := afero.ReadFile(fstore.fileSystem, pathToTemplate)
		if err != nil {
			return err
		}

		file.WriteString(string(data))
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
