package migrain

import (
	"database/sql"
	"errors"
	"io/ioutil"
	"strings"
)

type MigrationDirection int

const (
	Up MigrationDirection = iota
	Down
)

type Migrain struct {
	Queries []string
}

func New() *Migrain {
	return &Migrain{}
}

func (m *Migrain) File(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	ls := strings.Split(string(file), "\n")

	query := strings.Fields(strings.Join(ls, ""))

	m.Queries = append(m.Queries, strings.Join(query, " "))

	return nil
}

func (m *Migrain) Exec(db *sql.DB) error {
	if len(m.Queries) == 0 {
		return errors.New("no migration file found")
	}

	for _, query := range m.Queries {
		_, err := db.Exec(query)
		if err != nil {
			return err
		}
	}

	return nil
}
