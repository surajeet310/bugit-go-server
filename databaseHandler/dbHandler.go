package databaseHandler

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var database *sql.DB
var lock sync.Mutex
var once sync.Once

const (
	host     = "127.0.0.1"
	port     = "5432"
	user     = "bugit_user"
	password = "random123"
	dbname   = "bugit_test_db"
)

func OpenDbLocal() *sql.DB {
	if database == nil {
		dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
		db, err := sql.Open("postgres", dbInfo)
		if err != nil {
			log.Fatalf("Connection Error %s", err)
		}
		err = db.Ping()
		if err != nil {
			log.Fatalf("Ping Error %s", err)
		}
		log.Println("Connection Established.")
		database = db
	}

	return database
}

func OpenDbConnection() {
	lock.Lock()
	defer lock.Unlock()
	if database == nil {
		once.Do(func() {
			dbInfo := os.Getenv("DATABASE_URL")
			db, err := sql.Open("postgres", dbInfo)
			if err != nil {
				log.Fatalf("Connection Error %s", err)
			}
			err = db.Ping()
			if err != nil {
				log.Fatalf("Ping Error %s", err)
			}
			log.Println("Connection Established.")
			database = db
		})
	}
}

func OpenDbConnectionLocal() *sql.DB {
	return database
}
