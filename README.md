# example

```
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

	migrainInstance := migrain.New()

	err = migrainInstance.ReadFile("testdata/articles.sql")
	if err != nil {
		panic(err)
	}

	err = migrainInstance.Exec(db, migrain.Down)
	if err != nil {
		panic(err)
	}

	err = migrainInstance.Exec(db, migrain.Down)
	if err != nil {
		panic(err)
	}
}
```
