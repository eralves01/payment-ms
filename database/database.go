package database

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

type Database struct{}

var instance *sql.DB
var once sync.Once

func GetInstance() *sql.DB {
	once.Do(func() {
		fmt.Println("Connecting to database...")
		var err error
		instance, err = Connect()
		if err != nil {
			panic(err)
		}
	})
	return instance
}

func Connect() (*sql.DB, error) {
	return sql.Open("postgres", createDatabaseURL())
}

func createDatabaseURL() string {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	return "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"
}
