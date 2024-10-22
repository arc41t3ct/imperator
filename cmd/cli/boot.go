package main

import (
	"os"

	"github.com/arc41t3ct/imperator"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

// bootImperator setups the cli with things it needs
func bootImperator() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	// command line bootstraping imperator
	imp.RootPath = path
	imp.DB.DatabaseType = os.Getenv("DATABASE_TYPE")
	if imp.DB.DatabaseType == "mariadb" {
		imp.DB.DatabaseType = "mysql"
	}
	if imp.DB.DatabaseType == "postgressql" {
		imp.DB.DatabaseType = "postgres"
	}
	imp.Debug = false
	if os.Getenv("DEBUG") == "true" {
		imp.Debug = true
	}
	infoLog, errorLog := imp.StartLoggers()
	imp.InfoLog = infoLog
	imp.ErrorLog = errorLog
	imp.Validator = &imperator.Validation{
		Errors: make(map[string]string),
	}
	return nil
}

// showHelp show the help text when invalid arguments have been passed or help is requested.
func showHelp() {
	color.Yellow(`Imperator - easily build fullstack web apps
    Available commands:
        help                   - show the help command
        version                - show the imperator version
        new <appname> <type>   - creates a new imperator project application
                               - types: blog, portal, landing
        make                   - create scaffolds in your project with these sub commands:
            auth               - integrates authentication into the current application
            session            - integrates the session type into the current application
                                 session types: cookie, postgres, mysql, redis
            migration <name>   - creates new up and down migration templates
            mail <name>        - creates new mail templates for html and plain text mail
            key                - prints a 32 character length encryption key
            provider <name>    - creates new provider templates
            model <name>       - creates new model templates
            repos <name>       - creates new repository templates
            handler <name>     - creates new handler templates
            middleware <name>  - creates new middleware templates
            views <name>       - creates new jet view templates
            viewsgo <name>     - creates new go view templates
        test                   - create new tests for different items with these sub commands
            provider           - creates new provider test templates
            model              - creates new model test templates
        migrate                - migrate your database with these sub commands:
            up                 - migrate up all new migrations
            down               - migrate down one migration
            down all           - migrate all migrations down
            steps              - migrate steps (ex: -2 or 3) up or down
            force              - force migrations
            reset              - reset migrations by going down all and up again
    `)
}
