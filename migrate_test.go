package migrain

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestMigrate(t *testing.T) {
	db, err := sql.Open("mysql", "root:@/migrain")
	if err != nil {
		panic(err)
	}

	migrainInstance := New()

	err = migrainInstance.ReadFile("testdata/articles.sql")
	if err != nil {
		panic(err)
	}

	err = migrainInstance.Exec(db, Up)
	if err != nil {
		panic(err)
	}

	err = migrainInstance.Exec(db, Down)
	if err != nil {
		panic(err)
	}
}
