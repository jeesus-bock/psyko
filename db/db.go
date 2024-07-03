package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func SetupDB() (err error) {
	log.Println("Setting up database.")
	DB, err = sql.Open("sqlite3", "psyko.db")
	if err != nil {
		return
	}

	_, err = DB.Exec(createTransSQL)
	if err != nil {
		return
	}
	_, err = DB.Exec(seedTransSQL)
	if err != nil {
		return
	}

	_, err = DB.Exec(createDataSQL)
	return
}

const createDataSQL string = `
  CREATE TABLE IF NOT EXISTS data (
  id INTEGER NOT NULL PRIMARY KEY,
  trans_id INTEGER NOT NULL,
  key VARCHAR(128) NOT NULL,
  val VARCHAR(1024) NOT NULL,
  FOREIGN KEY (trans_id) REFERENCES trans (id)
  );`

const createTransSQL string = `
  CREATE TABLE IF NOT EXISTS trans (
  id INTEGER NOT NULL PRIMARY KEY,
  entity VARCHAR(128) NOT NULL,
  ts DATETIME NOT NULL,
  desc VARCHAR(1024) NOT NULL
  );`

const seedTransSQL string = `INSERT INTO trans (entity, ts, desc) 
  VALUES ('system', CURRENT_TIMESTAMP, 'Initial transaction');`
