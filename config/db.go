package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB
var err error

func ConnDB() {
	cstring := os.Getenv("cstring")

	DB, err = sql.Open("postgres", cstring)
	if err != nil {
		panic(err)
	}

	fmt.Println("database connected successfully")
}
