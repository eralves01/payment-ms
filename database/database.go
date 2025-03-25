package database

import (
	"database/sql"
	"os"
	"sync"

	"github.com/eralves01/payment-ms/configs"

	_ "github.com/lib/pq"
)

var instance *sql.DB
var once sync.Once

type Database struct{}

func NewDatadase() *Database {
	return &Database{}
}

func (db *Database) GetInstance() *sql.DB {
	log := configs.NewLogger("payment-ms")
	once.Do(func() {
		log.Info("Connecting to database...")
		var err error
		instance, err = connect()
		if err != nil {
			panic(err)
		}
	})
	return instance
}

func connect() (*sql.DB, error) {
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
