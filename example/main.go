package main

import (
	"database/sql"
	"github.com/milhamsuryapratama/migrain"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@/migrain")
	if err != nil {
		panic(err)
	}

	migrain := migrain.New()

	err = migrain.File("testdata/articles.sql")
	if err != nil {
		panic(err)
	}

	err = migrain.Exec(db)
	if err != nil {
		panic(err)
	}
}
