package main

import "strings"

func newApp(appName string) error {
	// sanitize the application
	var err error
	appName, err = sanitizeAppName(appName)
	if err != nil {
		return err
	}
	imp.InfoLog.Println("Creating app:", appName)

	// git clone skelelton

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
