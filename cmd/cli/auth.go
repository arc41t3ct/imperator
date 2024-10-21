package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func implementAuth() error {
	dbType := imp.DB.DatabaseType
	// create migration
	if err := createUsersTable(dbType); err != nil {
		return err
	}
	if err := createTokensTable(dbType); err != nil {
		return err
	}
	if err := createRememberTokensTable(dbType); err != nil {
		return err
	}
	// run migration
	if err := migrate("up", ""); err != nil {
		return err
	}
	// copy files over
	if err := copyAuthFiles(); err != nil {
		return err
	}

	color.Green("[+] Authentincation has been implemented...")
	color.Yellow(" - added migrations: users, tokens, remeber_tokens to /migrations and ran migrate up")
	color.Yellow(" - copied models: user.go, token.go to /models")
	color.Yellow(" - copied middleware: auth.go, auth_token.go remeber.go to /middleware")
	color.Yellow(" - copied handlers: login.go, logout.go password.go to /handlers")
	color.Yellow(" - copied mail: password_reset.html and password_reset.plain to /mail")
	color.Yellow(" - copied views: layout/base.jet, password_forgot.jet, password_reset.jet and login.jet to /views")
	fmt.Println("")
	color.Green("[!] Next manual steps...")
	color.Blue(" - remember to add User, RemeberToken and Token models to /models/models.go")
	color.Blue(" - remember to add Auth, Remember and AuthToken middleware to /routes.go")
	color.Blue(" - remember to add Login, Logout and Password handlers to /routes.go")

	return nil
}

func copyAuthFiles() error {
	// models
	if err := copyFileFromTemplate(
		"templates/models/user.go.txt",
		imp.RootPath+"/models/user.go"); err != nil {
		return err
	}
	if err := copyFileFromTemplate(
		"templates/models/token.go.txt",
		imp.RootPath+"/models/token.go"); err != nil {
		return err
	}
	if err := copyFileFromTemplate(
		"templates/models/token_token.go.txt",
		imp.RootPath+"/models/token_token.go"); err != nil {
		return err
	}
	// middleware
	if err := copyFileFromTemplate(
		"templates/middleware/auth.go.txt",
		imp.RootPath+"/middleware/auth.go"); err != nil {
		return err
	}
	if err := copyFileFromTemplate(
		"templates/middleware/auth_token.go.txt",
		imp.RootPath+"/middleware/auth_token.go"); err != nil {
		return err
	}
	if err := copyFileFromTemplate(
		"templates/middleware/remeber.go.txt",
		imp.RootPath+"/middleware/remember.go"); err != nil {
		return err
	}
	// handlers
	if err := copyFileFromTemplate(
		"templates/handlers/login.go.txt",
		imp.RootPath+"/handler/login.go"); err != nil {
		return err
	}
	if err := copyFileFromTemplate(
		"templates/handlers/logout.go.txt",
		imp.RootPath+"/handler/logout.go"); err != nil {
		return err
	}
	if err := copyFileFromTemplate(
		"templates/handlers/password.go.txt",
		imp.RootPath+"/handler/password.go"); err != nil {
		return err
	}
	// mail
	if err := copyFileFromTemplate(
		"templates/mail/password_reset.html.tmpl",
		imp.RootPath+"/mail/password_reset.html.tmpl"); err != nil {
		return err
	}
	if err := copyFileFromTemplate(
		"templates/mail/password_reset.plain.tmpl",
		imp.RootPath+"/mail/password_reset.plain.tmpl"); err != nil {
		return err
	}
	// views
	if err := copyFileFromTemplate(
		"templates/views/layouts/base.jet",
		imp.RootPath+"/views/layouts/base.jet"); err != nil {
		return err
	}
	if err := copyFileFromTemplate(
		"templates/views/password_forgot.jet",
		imp.RootPath+"/views/password_forgot.jet"); err != nil {
		return err
	}
	if err := copyFileFromTemplate(
		"templates/views/password_reset.jet",
		imp.RootPath+"/views/password_reset.jet"); err != nil {
		return err
	}
	if err := copyFileFromTemplate(
		"templates/views/login.jet",
		imp.RootPath+"/views/layouts/login.jet"); err != nil {
		return err
	}
	return nil
}

func createUsersTable(dbType string) error {
	fileName := fmt.Sprintf("%d_create_users_table_and_set_time_stamp", time.Now().UnixMicro())
	upFile := imp.RootPath + "/migrations/" + fileName + ".up.sql"
	downFile := imp.RootPath + "/migrations/" + fileName + ".down.sql"
	if err := copyFileFromTemplate(
		"templates/migrations/users_table_with_set_timestamp."+dbType+".up.sql", upFile); err != nil {
		return err
	}
	if err := copyFileFromTemplate(
		"templates/migrations/users_table_with_set_timestamp."+dbType+".down.sql", downFile); err != nil {
		return err
	}
	return nil
}

func createTokensTable(dbType string) error {
	fileName := fmt.Sprintf("%d_create_tokens_table", time.Now().UnixMicro())
	upFile := imp.RootPath + "/migrations/" + fileName + ".up.sql"
	downFile := imp.RootPath + "/migrations/" + fileName + ".down.sql"
	if err := copyFileFromTemplate(
		"templates/migrations/tokens_table."+dbType+".up.sql", upFile); err != nil {
		return err
	}
	if err := copyFileFromTemplate(
		"templates/migrations/tokens_table."+dbType+".down.sql", downFile); err != nil {
		return err
	}
	return nil
}

func createRememberTokensTable(dbType string) error {
	fileName := fmt.Sprintf("%d_create_remeber_tokens_table", time.Now().UnixMicro())
	upFile := imp.RootPath + "/migrations/" + fileName + ".up.sql"
	downFile := imp.RootPath + "/migrations/" + fileName + ".down.sql"
	if err := copyFileFromTemplate(
		"templates/migrations/remember_tokens_table."+dbType+".up.sql", upFile); err != nil {
		return err
	}
	if err := copyFileFromTemplate(
		"templates/migrations/remember_tokens_table."+dbType+".down.sql", downFile); err != nil {
		return err
	}
	return nil
}
