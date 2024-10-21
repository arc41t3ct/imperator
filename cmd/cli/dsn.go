package main

import (
	"errors"
	"fmt"
	"os"
)

func dsn() (string, error) {
	dbType := imp.DB.DatabaseType
	switch dbType {
	case "postgres", "postgresql":
		if os.Getenv("DATABASE_PASSWORD") != "" {
			return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_PASSWORD"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE")), nil
		} else {
			return fmt.Sprintf("postgres://%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE")), nil
		}
	case "mysql":
		// todo(andre): this will not work and we need to build a valid connection string for
		// maria mysql
		return "mysql://" + imp.BuildDSN(), nil
	default:
	}
	return "", errors.New("DATABASE_TYPE needs to be set in your .env file")
}
