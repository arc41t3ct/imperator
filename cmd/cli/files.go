package main

import (
	"embed"
	"errors"
	"fmt"
	"os"
)

//go:embed templates/*
var templateFS embed.FS

func copyFileFromTemplate(templatePath, targetFile string) error {
	if fileExists(targetFile) {
		return errors.New(fmt.Sprintf("target file %s already exists", targetFile))
	}

	data, err := templateFS.ReadFile(templatePath)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to read template file with error: %s", err))
	}

	if err := copyDataToFile(data, targetFile); err != nil {
		return errors.New(fmt.Sprintf("failed to copy template data to file with error: %s", err))
	}
	return nil
}

func copyDataToFile(templateData []byte, targetFile string) error {
	if fileExists(targetFile) {
		return errors.New(fmt.Sprintf("target file %s already exists", targetFile))
	}
	if err := os.WriteFile(targetFile, templateData, 0644); err != nil {
		return err
	}
	return nil
}

func replaceDataInFile(templateData []byte, targetFile string) error {
	if err := os.WriteFile(targetFile, templateData, 0644); err != nil {
		return err
	}
	return nil
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
