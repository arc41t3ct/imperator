package main

import (
	"fmt"
	"time"
)

func implementSession() error {
	dbType := imp.DB.DatabaseType
	if err := createSessionTable(dbType); err != nil {
		return err
	}
	// run migration
	if err := migrate("up", ""); err != nil {
		return err
	}
	return nil
}

func createSessionTable(dbType string) error {
	fileName := fmt.Sprintf("%d_create_session_table", time.Now().UnixMicro())
	upFile := imp.RootPath + "/migrations/" + fileName + ".up.sql"
	downFile := imp.RootPath + "/migrations/" + fileName + ".down.sql"
	if err := copyFileFromTemplate(
		"templates/migrations/session_table."+dbType+".up.sql", upFile); err != nil {
		return err
	}
	if err := copyFileFromTemplate(
		"templates/migrations/session_table."+dbType+".down.sql", downFile); err != nil {
		return err
	}
	return nil
}
