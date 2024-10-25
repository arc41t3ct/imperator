package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ettle/strcase"
	"github.com/fatih/color"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

var (
	appName       string
	appDirName    string
	appModuleName string
)

func newApp(appType string, givenAppName string) error {
	color.Red("Creating new app: %s with type: %s", givenAppName, appType)
	// convert Example App Name to example-app-name
	appName = givenAppName
	appDirName = strcase.ToKebab(givenAppName)
	appModuleName = convertToModuleName(givenAppName)
	if err := checkAppPathExistsAndCreate(appDirName); err != nil {
		return err
	}

	// git clone skelelton
	var repoName string
	color.Green(fmt.Sprintf("[+] Creating %s using app type %s...", givenAppName, appType))
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
		color.Red("NOT YET IMPLEMENTED !!!")
		// repoName = "arc41t3ct/imperator-shop"
	default:
		return errors.New("imperator requires an <type> to pull a template from github")
	}

	// todo(andre): make these run in parallel

	color.Green("Cloning template from application repository: %s", repoName)
	if err := cloneAppRepo(appDirName, repoName); err != nil {
		return err
	}

	color.Green("Creating .env file in: %s", appDirName)
	if err := createEnvFile(appDirName, appName); err != nil {
		return err
	}

	color.Green("Deleting .git folder in: %s", appDirName)
	if err := deleteGitDir(appDirName); err != nil {
		return err
	}

	color.Green("Updating make files in: %s with module name: %s", appDirName, appModuleName)
	if err := updateMakeFiles(appDirName, appModuleName, givenAppName); err != nil {
		return err
	}

	color.Green("Updating the go.mod file .git folder in: %s", appDirName)
	if err := updateGoModFile(appDirName, appModuleName); err != nil {
		return err
	}

	color.Green("Updating all go file *.go imports in: %s", appDirName)
	if err := recursiveWalkFilesAndUpdate(appDirName); err != nil {
		return err
	}

	color.Green("Entering app dir: %s", appDirName)
	if err := enterNewAppDirectory(appDirName); err != nil {
		return err
	}

	color.Green("Running go mod tidy")
	if err := goRunModTidy(); err != nil {
		return err
	}

	return nil
}

func goRunModTidy() error {
	cmd := exec.Command("go", "mod", "tidy")
	if err := cmd.Start(); err != nil {
		return err
	}
	return nil
}

func updateGoFileImports(path string, fileInfo os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		return nil
	}

	matched, err := filepath.Match("*.go", fileInfo.Name())
	if err != nil {
		return err
	}

	if matched {
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		newContents := strings.ReplaceAll(string(content), "imperatorapp", appModuleName)
		if err := os.WriteFile(path, []byte(newContents), fileInfo.Mode()); err != nil {
			return err
		}
	}
	return nil
}

func recursiveWalkFilesAndUpdate(dir string) error {
	if err := filepath.Walk(dir, updateGoFileImports); err != nil {
		return err
	}
	return nil
}

func convertToModuleName(appName string) string {
	moduleName := strcase.ToSnake(appName)
	moduleName = strings.ReplaceAll(moduleName, "_", "")
	return moduleName
}

func updateGoModFile(appDirName, appModuleName string) error {
	path, err := getCurrentPath()
	if err != nil {
		return err
	}
	goModFile := fmt.Sprintf("%s/%s/go.mod", path, appDirName)
	goModData, err := os.ReadFile(goModFile)
	if err != nil {
		return err
	}
	goModContent := string(goModData)
	goModContent = strings.ReplaceAll(goModContent, "imperatorapp", appModuleName)
	if err := replaceDataInFile([]byte(goModContent), goModFile); err != nil {
		return err
	}
	return nil
}

func updateMakeFiles(appDirName, appExecutableName, appName string) error {
	path, err := getCurrentPath()
	if err != nil {
		return err
	}
	makeFile := fmt.Sprintf("%s/%s/Makefile", path, appDirName)
	makeFileData, err := os.ReadFile(makeFile)
	if err != nil {
		return err
	}
	makeFileContent := string(makeFileData)
	makeFileContent = strings.ReplaceAll(makeFileContent, "imperatorapp", appExecutableName)
	makeFileContent = strings.ReplaceAll(makeFileContent, "Imperator App", appName)
	if err := replaceDataInFile([]byte(makeFileContent), makeFile); err != nil {
		return err
	}
	return nil
}

func deleteGitDir(appDirName string) error {
	path, err := getCurrentPath()
	if err != nil {
		return err
	}
	if err := os.RemoveAll(fmt.Sprintf("%s/%s/.git", path, appDirName)); err != nil {
		return err
	}
	return nil
}

func getCurrentPath() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return path, nil
}

// checkAppPathExistsAndCreate checks if the path exists and throws and
// error if so. If it does not exist it gets created.
func checkAppPathExistsAndCreate(appDirName string) error {
	path, err := getCurrentPath()
	if err != nil {
		return err
	}
	appPath := fmt.Sprintf("%s/%s", path, appDirName)
	const mode = 0755
	if _, err := os.Stat(appPath); os.IsNotExist(err) {
		err := os.Mkdir(appPath, mode)
		if err != nil {
			return err
		}
	}
	return nil
}

// enterNewAppDirectory enters the new app directory.
func enterNewAppDirectory(appDirName string) error {
	path, err := getCurrentPath()
	if err != nil {
		return err
	}
	appPath := fmt.Sprintf("%s/%s", path, appDirName)
	if err := os.Chdir(appPath); err != nil {
		return err
	}
	return nil
}

// createEnvFile creates a new .emv file amd replaces app name and
// creates a new encryption key
func createEnvFile(appDirName, appName string) error {
	path, err := getCurrentPath()
	if err != nil {
		return err
	}
	fullPath := fmt.Sprintf("%s/%s/.env", path, appDirName)
	tmpl, err := templateFS.ReadFile("templates/rootlevel/dot.env.txt")
	if err != nil {
		return err
	}

	randString := imp.CreateRadomString(32)
	env := string(tmpl)
	env = strings.ReplaceAll(env, "{{APP_NAME}}", appName)
	env = strings.ReplaceAll(env, "{{ENCRYPTION_KEY}}", randString)
	// replace placeholders in .env
	if err := copyDataToFile([]byte(env), fullPath); err != nil {
		return err
	}
	return nil
}

func cloneAppRepo(appDirName, repo string) error {
	sshKeyPath := filepath.Join(os.Getenv("HOME"), ".ssh", "gh_deploy_rsa")
	publicKey, err := ssh.NewPublicKeysFromFile("git", sshKeyPath, "")
	if err != nil {
		return err
	}
	if _, err := git.PlainClone("./"+appDirName, false, &git.CloneOptions{
		Auth: publicKey,
		URL:  fmt.Sprintf("git@github.com:%s.git", repo),
		// Progress:     os.Stdout,
		SingleBranch: true,
		Depth:        1,
	}); err != nil {
		return err
	}
	return nil
}
