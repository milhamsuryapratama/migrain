package migrain

import (
	"database/sql"
)

type Migration struct {
	Id   string
	Up   []string
	Down []string
}

type MigrationDirection int

const (
	Up MigrationDirection = iota
	Down
)

func Exec(db *sql.DB, migrations []*Migration, direction MigrationDirection) error {
	for _, migration := range migrations {

		var migrationData []string
		if direction == Up {
			migrationData = migration.Up
		} else {
			migrationData = migration.Down
		}

		for _, migrate := range migrationData {
			_, err := db.Exec(migrate)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
