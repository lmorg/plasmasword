package main

import (
	"database/sql"
	"fmt"
	"github.com/lmorg/apachelogs"
	"log"
	"strings"
)

const (
	sqlInsertAccess = `INSERT INTO access (
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

	sqlInsertStatus = `INSERT INTO status (
							status,
							title,
							desc
						) VALUES (?, ?, ?);`
)

var (
	db       *sql.DB
	tx       *sql.Tx
	accessId uint
)

func OpenDB() {
	var err error

	log.Println("Opening database")

	openSqlite3()

	if tx, err = db.Begin(); err != nil {
		log.Fatalln("Could not open transaction:", err)
	}

	if _, err = tx.Exec(sqlCreateTable); err != nil {
		log.Fatalln("Could not create access table:", err)
	}

	// statuses
	log.Println("Adding status table")
	if _, err = tx.Exec(sqlCreateStatuses); err != nil {
		log.Fatalln("Could not create main.status table:", err)
	}

	for status := range statusTitle {
		if _, err := tx.Exec(sqlInsertStatus,
			status,
			statusTitle[status],
			statusDescription[status],
		); err != nil {
			log.Println("Error inserting status record:", err)
		}
	}

	// views
	log.Println("Adding views")
	_, err = tx.Exec(fmt.Sprint(
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

	if err = tx.Commit(); err != nil {
		log.Println("Could not commit transaction:", err)
	}

}

func InsertAccess(access *apachelogs.AccessLog) {
	accessId++
	_, err := tx.Exec(sqlInsertAccess,
		accessId,
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

func BeginTransaction() {
	var err error
	if tx, err = db.Begin(); err != nil {
		log.Fatalln("Could not open transaction:", err)
	}
}

func CommitTransaction() {
	if err := tx.Commit(); err != nil {
		log.Println("Error commiting access:", err)
	}
}

func CloseDB() {
	db.Close()
}
