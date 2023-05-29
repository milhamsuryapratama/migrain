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
	UpQueries   []string
	DownQueries []string
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

	//var ncls []string
	var query string
	var upQuery, downQuery []string
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
					upQuery = append(upQuery, query)
				}

				if isDownQuery {
					downQuery = append(downQuery, query)
				}

				query = ""
			}
		}
	}

	m.UpQueries = append(m.UpQueries, upQuery...)
	m.DownQueries = append(m.DownQueries, downQuery...)

	fmt.Println("m.UpQueries ", m.UpQueries)

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

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			return err
		}
	}

	log.Println("success exec migration")

	return nil
}
