package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Variabel golbal koneksi database
const (
	host     = "localhost"
	port     = 5434
	user     = "postgres"
	password = "1234"
	dbname   = "enigmalaundry"
)

// connection string
var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func ConnectDb() *sql.DB {
	// Open Database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// Check DB

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Database Connected!")
	}

	return db
}
