package fibnumdb

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "root"
// 	password = "mysecretpassword"
// 	dbname   = "root"
// )

var (
	Client *sql.DB
)

type ServiceConfig struct {
	DBDriver string
	DBSource string
}

func init() {
	var err error
	config := ServiceConfig{
		DBDriver: "postgres",
		DBSource: "postgresql://root:secret@localhost:5432/fibonacci?sslmode=disable",
	}
	if v := os.Getenv("DB_SOURCE"); v != "" {
		config.DBSource = v

	}

	fmt.Println("DBDriver", config.DBDriver)
	fmt.Println("DBSource", config.DBSource)

	Client, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}
