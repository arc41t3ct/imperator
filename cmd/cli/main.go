package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/arc41t3ct/imperator"
	"github.com/fatih/color"
)

const version = "1.0.0"

var imp imperator.Imperator

func main() {
	arg1, arg2, arg3, err := validateInput()
	if err != nil {
		exitGracefully(err)
	}
	if err := bootImperator(); err != nil {
		exitGracefully(err)
	}
	if err := run(arg1, arg2, arg3); err != nil {
		exitGracefully(err)
	}
	exitGracefully(nil, "ran successfully...")
}

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
		if err := migrate(arg2, arg3); err != nil {
			return errors.New(fmt.Sprintf("migrate failed with error: %s", err))
		}

	case "make":
		if arg2 == "" {
			showHelp()
			return errors.New("make requires a subcommand")
		}
		if err := makeCommand(arg2, arg3); err != nil {
			return errors.New(fmt.Sprintf("could not execute make with error: %s", err))
		}

	case "new":
		if arg2 == "" {
			showHelp()
			return errors.New("new requires an application name")
		}
		if err := newApp(arg2); err != nil {
			return errors.New(fmt.Sprintf("could not create new app with error: %s", err))
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
