package migrain

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func connectToDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@/migrain")
	if err != nil {
		panic(err)
	}

	return db
}

func TestReadFileMigration(t *testing.T) {
	db := connectToDB()

	migrain := New()

	err := migrain.ReadFile("testdata/01_articles.sql")
	if err != nil {
		panic(err)
	}

	err = migrain.Run(db, Up)
	if err != nil {
		panic(err)
	}

	err = migrain.Run(db, Down)
	if err != nil {
		panic(err)
	}
}

func TestReadDirMigration(t *testing.T) {
	db := connectToDB()

	migrain := New()

	err := migrain.ReadDir("testdata")
	if err != nil {
		panic(err)
	}

	err = migrain.Run(db, Up)
	if err != nil {
		panic(err)
	}

	err = migrain.Run(db, Down)
	if err != nil {
		panic(err)
	}
}
