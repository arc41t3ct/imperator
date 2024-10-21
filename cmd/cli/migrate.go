package main

import (
	"strconv"
)

func migrate(arg2, arg3 string) error {
	dsn, err := dsn()
	if err != nil {
		imp.ErrorLog.Println("failed to DSN:", err)
		return err
	}
	switch arg2 {
	case "up":
		if err := imp.MigrateUp(dsn); err != nil {
			imp.ErrorLog.Println("failed to run MigrateUp:", err)
			return err
		}
	case "down":
		if arg3 == "all" {
			if err := imp.MigrateDownAll(dsn); err != nil {
				imp.ErrorLog.Println("failed to run MigrateDownAll:", err)
				return err
			}
		} else {
			if err := imp.Steps(-1, dsn); err != nil {
				imp.ErrorLog.Println("failed to run Steps with -1:", err)
				return err
			}
		}
	case "reset":
		if err := imp.MigrateDownAll(dsn); err != nil {
			imp.ErrorLog.Println("failed to run MigrateDownAll:", err)
			return err
		}
		if err := imp.MigrateUp(dsn); err != nil {
			imp.ErrorLog.Println("failed to run MigrateUp:", err)
			return err
		}
	case "steps":
		arg3Int, err := strconv.Atoi(arg3)
		if err != nil {
			imp.ErrorLog.Println("failed to run convert string int:", err)
			return err
		}
		if err := imp.Steps(arg3Int, dsn); err != nil {
			imp.ErrorLog.Println("failed to run Steps:", err)
			return err
		}
	case "force":
		if err := imp.MigrateForce(dsn); err != nil {
			imp.ErrorLog.Println("failed to run MigrateForce:", err)
			return err
		}
	default:
		showHelp()
	}
	return nil
}
