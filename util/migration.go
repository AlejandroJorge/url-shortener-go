package util

import (
	"database/sql"
	"os"
)

func RunMigration(db *sql.DB) error {
	migrationSQLBytes, err := os.ReadFile("sql/migration.sql")
	if err != nil {
		return err
	}
	migrationSQL := string(migrationSQLBytes)

	_, err = db.Exec(migrationSQL)
	if err != nil {
		return err
	}

	return nil
}
