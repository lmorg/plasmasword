package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

/*
const (
	sqlCreateTable = `CREATE TABLE IF NOT EXISTS access (
							id          INTEGER PRIMARY KEY,
							ip          TEXT,
							method      TEXT,
							proc        INTEGER,
							proto       TEXT,
							qs          TEXT,
							ref         TEXT,
							size        INTEGER,
							status      INTEGER,
							datetime    TEXT,
							uri         TEXT,
							ua          TEXT,
							uid         TEXT,
							file        TEXT
						);`

	sqlCreateStatuses = `CREATE TABLE IF NOT EXISTS status (
							status      INTEGER PRIMARY KEY,
							title       TEXT,
							desc        TEXT
						);`
)
*/
func init() {
	dbEngine["sqlite3"] = openSqlite3
}

func openSqlite3() {
	var err error
	//if db, err = sql.Open("sqlite3", "file:"+fDbFileName); err != nil {
	//	log.Fatalln(errCouldNotOpenDb, err)
	//}
	if db, err = sql.Open("sqlite3", "file:"+fDbConnStr); err != nil {
		log.Fatalln(errCouldNotOpenDb, err)
	}
}
