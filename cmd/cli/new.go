package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	git "github.com/go-git/go-git/v5"
)

func newApp(appName string, appType string) error {
	// sanitize the application
	var err error
	appName, err = sanitizeAppName(appName)
	if err != nil {
		return err
	}
	imp.InfoLog.Println("App Name:", appName)

	if err := createEnvFile(); err != nil {
		return err
	}

	// git clone skelelton
	var repoName string
	color.Green("[+] Cloning...")
	switch appType {
	case "experimental":
		repoName = "arc41t3ct/imperator_app"
	case "landing":
		repoName = "arc41t3ct/imperator-landing"
	case "blog":
		repoName = "arc41t3ct/imperator-blog"
	case "portal":
		repoName = "arc41t3ct/imperator-portal"
	case "shop":
		repoName = "arc41t3ct/imperator-shop"
	default:
	}

	if err := cloneAppRepo(appName, repoName); err != nil {
		return err
	}

	// remove .git directory

	// create ready .env file which we need to create and not store in github

	// create make files

	// update the go.mod file

	// update existing go files with correct name imports
	// copy app files
	return nil
}

func sanitizeAppName(appName string) (string, error) {
	appName = strings.ToLower(appName)
	// convert url to single word
	if strings.Contains(appName, "/") {
		splitted := strings.SplitAfter(appName, "/")
		appName = splitted[(len(splitted) - 1)]
	}
	return appName, nil
}

func createEnvFile() error {
	if err := copyFileFromTemplate(
		"templates/rootlevel/dot.env.txt",
		imp.RootPath+"/.env"); err != nil {
		return err
	}
	return nil
}

func cloneAppRepo(appName, repo string) error {
	_, err := os.Stat(fmt.Sprintf("./%s", appName))
	if err != nil {
		return errors.New(fmt.Sprintf("app directory already exists"))
	}
	if _, err := git.PlainClone("./"+appName, false, &git.CloneOptions{
		URL:      "git@github.com/" + repo,
		Progress: os.Stdout,
		Depth:    1,
	}); err != nil {
		return err
	}
	return nil
}
