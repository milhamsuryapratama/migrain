package migrain

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var sqliteMigrations = []*Migration{
	{
		Up:   []string{"CREATE TABLE IF NOT EXISTS product(product_id int primary key auto_increment, product_name text, product_price int, created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP) "},
		Down: []string{"DROP TABLE product"},
	},
}

func TestMigrate(t *testing.T) {
	db, err := sql.Open("mysql", "root:@/migrain")
	if err != nil {
		panic(err)
	}

	err = Exec(db, sqliteMigrations, Up)
	if err != nil {
		panic(err)
	}

	err = Exec(db, sqliteMigrations, Down)
	if err != nil {
		panic(err)
	}
}
