package migrain

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type MigrationDirection int

const (
	Up MigrationDirection = iota
	Down
)

type Migrain struct {
	UpQueries   []Migration
	DownQueries []Migration
}

type Migration struct {
	Query    string
	FileName string
	Batch    int
}

func New() *Migrain {
	return &Migrain{}
}

func (m *Migrain) ReadFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	ls := strings.Split(string(file), "\n")

	var query string
	var isUpQuery, isDownQuery bool
	for _, l := range ls {
		if l == "-- UP --" {
			isUpQuery = true
			isDownQuery = false
			continue
		}

		if l == "-- DOWN --" {
			isUpQuery = false
			isDownQuery = true
			continue
		}

		if isUpQuery || isDownQuery {
			query += l
			if strings.Contains(l, ";") {
				// TODO: need to trim space
				if isUpQuery {
					m.UpQueries = append(m.UpQueries, Migration{
						FileName: path,
						Query:    query,
					})
				}

				if isDownQuery {
					m.DownQueries = append(m.DownQueries, Migration{
						FileName: path,
						Query:    query,
					})
				}

				query = ""
			}
		}
	}

	return nil
}

func (m *Migrain) ReadDir(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		err = m.ReadFile(fmt.Sprintf("%s/%s", dir, file.Name()))
		if err != nil {
			panic(err)
		}
	}

	return nil
}

func (m *Migrain) Run(db *sql.DB, migrationDirection MigrationDirection) error {
	if migrationDirection == Up && len(m.UpQueries) == 0 {
		return errors.New("no migration file found")
	}

	if migrationDirection == Down && len(m.DownQueries) == 0 {
		return errors.New("no migration file found")
	}

	var queries = m.UpQueries
	if migrationDirection == Down {
		queries = m.DownQueries
	}

	var tableName string
	row := db.QueryRow("SELECT TABLE_NAME FROM information_schema.tables" +
		" WHERE table_name = 'migrations'")

	err := row.Scan(&tableName)
	if err != nil {
		panic(err)
	}

	batch := 0
	var migratedFile []Migration

	if tableName == "" {
		// init migrations table
		_, err = db.Exec("CREATE TABLE IF NOT EXISTS migrations(" +
			"id INT NOT NULL AUTO_INCREMENT, " +
			"migration VARCHAR(255) NOT NULL DEFAULT '', " +
			"batch INT NOT NULL DEFAULT 0, " +
			"PRIMARY KEY (id))")
		if err != nil {
			panic(err)
		}
	} else {
		rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s", tableName))
		if err != nil {
			panic(err)
		}

		defer rows.Close()

		for rows.Next() {
			var migration Migration
			err := rows.Scan(&migration.Query, &migration.FileName, &migration.Batch)
			if err != nil {
				panic(err)
			}

			migratedFile = append(migratedFile, migration)
		}
	}

	var migratedFileMap = make(map[string]bool)
	if len(migratedFile) > 0 {
		for _, migration := range migratedFile {
			migratedFileMap[migration.FileName] = true
			batch = migration.Batch
		}
	}

	batch++

	for _, query := range queries {
		if len(migratedFileMap) > 0 && migratedFileMap[query.FileName] {
			continue
		}

		if migrationDirection == Up {
			_, err := db.Exec("INSERT INTO migrations(migration, batch) VALUES(?, ?)", query.FileName, batch)
			if err != nil {
				panic(err)
			}
		}

		if migrationDirection == Down {
			_, err := db.Exec(fmt.Sprintf("DELETE FROM migrations WHERE migration = '%s'", query.FileName))
			if err != nil {
				panic(err)
			}
		}

		_, err = db.Exec(query.Query)
		if err != nil {
			return err
		}

		log.Printf("success run %s migration", query.FileName)
	}

	log.Println("success exec migration")

	return nil
}
