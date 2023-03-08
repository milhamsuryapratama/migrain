package main

import (
	"database/sql"
	"github.com/milhamsuryapratama/migrain"

	_ "github.com/go-sql-driver/mysql"
)

var sqliteMigrations = []*migrain.Migration{
	{
		Up:   []string{"CREATE TABLE IF NOT EXISTS product(product_id int primary key auto_increment, product_name text, product_price int, created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP) "},
		Down: []string{"DROP TABLE product"},
	},
}

func main() {
	db, err := sql.Open("mysql", "root:@/migrain")
	if err != nil {
		panic(err)
	}

	err = migrain.Exec(db, sqliteMigrations, migrain.Up)
	if err != nil {
		panic(err)
	}

	err = migrain.Exec(db, sqliteMigrations, migrain.Down)
	if err != nil {
		panic(err)
	}
}
