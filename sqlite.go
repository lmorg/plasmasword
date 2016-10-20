package main

import (
	"database/sql"
	"fmt"
	"github.com/lmorg/apachelogs"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strings"
)

const (
	sqlCreateTable = `CREATE TABLE IF NOT EXISTS %s.access (
							id          integer PRIMARY KEY,
							ip          string,
							method      string,
							proc        integer,
							proto       string,
							qs          string,
							ref         string,
							size        integer,
							status      integer,
							datetime    datetime,
							uri         string,
							ua          string,
							uid         string,
							file        string
						);`

	sqlInsertAccess = `INSERT INTO mem.access (
							id,
							ip,
							method,
							proc,
							proto,
							qs,
							ref,
							size,
							status,
							datetime,
							uri,
							ua,
							uid,
							file
						) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

	sqlCreateStatuses = `CREATE TABLE IF NOT EXISTS %s.status (
							status      integer PRIMARY KEY,
							title       string,
							desc        string
						);`

	sqlInsertStatus = `INSERT INTO mem.status (
							status,
							title,
							desc
						) VALUES (?, ?, ?);`

	sqlSyncToDisk     = `INSERT INTO main.access SELECT * FROM mem.access;`
	sqlPurgeMemAccess = `DELETE FROM mem.access;`
	sqlSyncStatus     = `INSERT INTO main.status SELECT * FROM mem.status;`
	sqlPurgeMemStatus = `DROP TABLE mem.status;`
)

var (
	db  *sql.DB
	aid uint
)

func OpenDB() {
	var err error

	log.Println("Opening database")

	db, err = sql.Open("sqlite3", fmt.Sprintf("file:%s", fDbFileName))
	if err != nil {
		log.Fatalln("Could not open database:", err)
	}

	if _, err = db.Exec(fmt.Sprintf(sqlCreateTable, "main")); err != nil {
		log.Fatalln("Could not create main.access table:", err)
	}

	if _, err = db.Exec(`ATTACH DATABASE ':memory:' AS mem;`); err != nil {
		log.Fatalln("Could not create in memory database")
	}

	if _, err = db.Exec(fmt.Sprintf(sqlCreateTable, "mem")); err != nil {
		log.Fatalln("Could not create mem.access table:", err)
	}

	// statuses
	log.Println("Adding status table")
	if _, err = db.Exec(fmt.Sprintf(sqlCreateStatuses, "main")); err != nil {
		log.Fatalln("Could not create main.status table:", err)
	}
	if _, err = db.Exec(fmt.Sprintf(sqlCreateStatuses, "mem")); err != nil {
		log.Fatalln("Could not create mem.status table:", err)
	}

	for status := range statusTitle {
		if _, err := db.Exec(sqlInsertStatus,
			status,
			statusTitle[status],
			statusDescription[status],
		); err != nil {
			log.Println("Error inserting status record:", err)
		}
	}

	if _, err = db.Exec(sqlSyncStatus); err != nil {
		log.Println("Error syncing statuses to disk:", err)
	}
	db.Exec(sqlPurgeMemStatus)

	// views
	log.Println("Adding views")
	_, err = db.Exec(fmt.Sprint(
		viewAll,
		viewLatestNon200,
		viewLatestProc,
		viewLatest304,
		viewCountStatus,
		viewCount304,
		viewCountSize,
		viewListViews,
	))
	if err != nil && !strings.HasSuffix(err.Error(), "already exists") {
		log.Println("Could not create views:", err)
	}

}

func InsertAccess(access *apachelogs.AccessLog) {
	aid++
	_, err := db.Exec(sqlInsertAccess,
		aid,
		access.IP,
		access.Method,
		access.ProcTime,
		access.Protocol,
		access.QueryString,
		access.Referrer,
		access.Size,
		access.Status.I,
		access.DateTime,
		access.URI,
		access.UserAgent,
		access.UserID,
		access.FileName,
	)

	if err != nil {
		log.Println("Error inserting access log:", err)
	}

	return
}

func SyncDbToDisk() {
	if _, err := db.Exec(sqlSyncToDisk); err != nil {
		log.Println("Error syncing to disk:", err)
	}
	db.Exec(sqlPurgeMemAccess)
}

func CloseDB() {
	db.Close()
}
