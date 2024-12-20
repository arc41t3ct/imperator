package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/arc41t3ct/imperator"
	"github.com/fatih/color"
)

const version = "v0.0.1"

var imp imperator.Imperator

func main() {
	arg1, arg2, arg3, err := validateInput()
	if err != nil {
		exitGracefully(err)
	}
	if err := run(arg1, arg2, arg3); err != nil {
		exitGracefully(err)
	}
	exitGracefully(nil, "ran successfully...")
}

// run uses the arguments to figure out what to run. Since
// some operations require running before bootImperator runs
// we need to decide at a case by case level
func run(arg1, arg2, arg3 string) error {
	switch arg1 {
	case "help":
		showHelp()

	case "version":
		color.Yellow("imperator version: " + version)

	case "migrate":
		if arg2 == "" {
			arg2 = "up"
		}
		if err := bootImperator(); err != nil {
			return err
		}
		if err := migrate(arg2, arg3); err != nil {
			return errors.New(fmt.Sprintf("migrate failed with error: %s", err))
		}

	case "make":
		if arg2 == "" {
			showHelp()
			return errors.New("make requires a subcommand")
		}
		if err := bootImperator(); err != nil {
			return err
		}
		if err := makeCommand(arg2, arg3); err != nil {
			return errors.New(fmt.Sprintf("could not execute make with error: %s", err))
		}

	case "new":
		if arg2 == "" {
			showHelp()
			return errors.New("new requires an application type")
		}
		if arg3 == "" {
			showHelp()
			return errors.New("new requires an application name in form \"App Name\"")
		}
		if err := newApp(arg2, arg3); err != nil {
			return errors.New(fmt.Sprintf("could not create new app with error: %s", err))
		}
		if err := bootImperator(); err != nil {
			return err
		}
	default:
		showHelp()
	}
	return nil
}

func debugError(msg string) {
	if imp.Debug {
		color.Red(msg)
	}
}

func debugInfo(msg string) {
	if imp.Debug {
		color.Blue(msg)
	}
}

// validateInput validates the arguments being passed to Imperator
func validateInput() (string, string, string, error) {
	var arg1, arg2, arg3 string
	if len(os.Args) > 1 {
		arg1 = os.Args[1]
		if len(os.Args) >= 3 {
			arg2 = os.Args[2]
		}
		if len(os.Args) >= 4 {
			arg3 = os.Args[3]
		}
	} else {
		showHelp()
		return "", "", "", errors.New("command required")
	}

	return arg1, arg2, arg3, nil
}

func exitGracefully(err error, msg ...string) {
	message := ""
	if len(msg) > 0 {
		message = msg[0]
	}
	if err != nil {
		color.Red("Error: %v\n", err)
	}

	if len(message) > 0 {
		color.Yellow(message)
	} else {
		color.Green("imperator finished...")
	}
	os.Exit(0)
}
