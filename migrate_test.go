package migrain

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestMigrate(t *testing.T) {
	_, err := sql.Open("mysql", "root:@/migrain")
	if err != nil {
		panic(err)
	}

	migrain := New()

	err = migrain.File("testdata/articles.sql")
	if err != nil {
		panic(err)
	}

	//err = migrain.Exec(db)
	//if err != nil {
	//	panic(err)
	//}
}
