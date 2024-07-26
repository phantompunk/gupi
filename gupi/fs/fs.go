package fs

import (
	"errors"
	"os"
	"path/filepath"
)

const appDir string = ".gupi"
const tmplDir string = "templates"

func GetHomeDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir, nil
}

func GetTemplateDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, appDir, tmplDir), nil
}

func GetTemplatePath(file string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, appDir, tmplDir, file), nil
}

func ReadTemplates() ([]os.FileInfo, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, errors.New("Failed to find home directory")
	}

	fileDir := filepath.Join(homeDir, appDir, tmplDir)
	err = createDirIfNotExists(fileDir)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(fileDir)
	if err != nil {
		return nil, errors.New("Template folder not found")
	}
	defer file.Close()

	filelist, err := file.Readdir(0)
	if err != nil {
		return nil, errors.New("Unable to read file")
	}
	return filelist, nil
}

func CreateFile(fileName string) error {
	dir, err := GetTemplateDir()
	if err != nil {
		return err
	}

	fullPath := filepath.Join(dir, fileName)
	err = createDirIfNotExists(dir)
	if err != nil {
		return err
	}

	_, err = os.Create(fullPath)
	if err != nil {
		return err
	}
	return nil
}

func createDirIfNotExists(pathDir string) error {
	_, err := os.Stat(pathDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(pathDir, os.ModePerm)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}
