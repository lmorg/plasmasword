//// +build ignore

package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	sqlCreateAccess = `CREATE TABLE IF NOT EXISTS access (
							id          INTEGER PRIMARY KEY,
							ip          VARCHAR[100],
							method      VARCHAR[5],
							proc        INTEGER,
							proto       VARCHAR[5],
							qs          VARCHAR[500],
							ref         VARCHAR[500],
							size        INTEGER,
							status      INTEGER,
							datetime    TEXT,
							uri         VARCHAR[500],
							ua          VARCHAR[500],
							uid         VARCHAR[100],
							file        VARCHAR[100]
						);`

	sqlCreateError = `CREATE TABLE IF NOT EXISTS error (
							id          INTEGER PRIMARY KEY,
							datetime    TEXT,
							timestamped CHAR,
							scope1      VARCHAR[100],
							scope2      VARCHAR[100],
							scope3      VARCHAR[100],
							scope4      VARCHAR[100],
							scope5      VARCHAR[100],
							scope6      VARCHAR[100],
							scope7      VARCHAR[100],
							scope8      VARCHAR[100],
							scope9      VARCHAR[100],
							scope10     VARCHAR[100],
							message     VARCHAR[2000],
							file        VARCHAR[100]
						);`

	sqlCreateStatuses = `CREATE TABLE IF NOT EXISTS status (
							status      INTEGER PRIMARY KEY,
							title       VARCHAR[100],
							desc        VARCHAR[1000]
						);`
)

func init() {
	dbEngine["mysql"] = openMySQL
}

func openMySQL() {
	/*
		var err error
		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4,utf8",
			fDbUserName, fDbPassword, fDbHostName, fDbPortNumber, fDbSchema))
		if err != nil {
			log.Fatalln(errCouldNotOpenDb, err)
		}
	*/

	var err error
	db, err = sql.Open("mysql", fDbConnStr+"charset=utf8mb4,utf8")
	if err != nil {
		log.Fatalln(errCouldNotOpenDb, err)
	}

}
