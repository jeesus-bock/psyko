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

func GetAllTrans() (ret []Trans, err error) {
	ret = make([]Trans, 0)
	rows, err := DB.Query("SELECT * FROM trans;")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var trans Trans
		if err := rows.Scan(&trans.Id, &trans.Entity, &trans.Ts); err != nil {
			return ret, err
		}
		ret = append(ret, trans)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}
