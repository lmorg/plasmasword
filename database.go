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

	sqlInsertError = `INSERT INTO error (
							id,
							datetime,
							timestamped,
							scope1,
							scope2,
							scope3,
							scope4,
							scope5,
							scope6,
							scope7,
							scope8,
							scope9,
							scope10,
							message,
							file
						) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

	sqlInsertStatus = `INSERT INTO status (
							status,
							title,
							desc
						) VALUES (?, ?, ?);`

	errCouldNotOpenDb = "Could not open database:"
)

var (
	db         *sql.DB
	tx         *sql.Tx
	accessId   uint
	errorId    uint
	dbEngine   map[string]func() = make(map[string]func())
	fDbDriver  string
	fDbConnStr string
	//fDbFileName   string
	//fDbUserName   string
	//fDbPassword   string
	//fDbHostName   string
	//fDbPortNumber string
	//fDbSchema     string
)

var emptyScopes []string = make([]string, 10)

func OpenDB() {
	var err error

	log.Println("Opening database")

	if dbEngine[fDbDriver] != nil {
		dbEngine[fDbDriver]()
	} else {
		log.Fatalln("Database driver", fDbDriver, "does not exist.")
	}

	if tx, err = db.Begin(); err != nil {
		log.Fatalln("Could not open transaction:", err)
	}

	if _, err = tx.Exec(sqlCreateAccess); err != nil {
		log.Fatalln("Could not create access table:", err)
	}

	if _, err = tx.Exec(sqlCreateError); err != nil {
		log.Fatalln("Could not create error table:", err)
	}

	// statuses
	log.Println("Adding status table")
	if _, err = tx.Exec(sqlCreateStatuses); err != nil {
		log.Fatalln("Could not create main.status table:", err)
	}

	for status := range apachelogs.StatusTitle {
		if _, err := tx.Exec(sqlInsertStatus,
			status,
			apachelogs.StatusTitle[status],
			apachelogs.StatusDescription[status],
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

func InsertAccess(access *apachelogs.AccessLine) {
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

func InsertError(error *apachelogs.ErrorLine) {
	var (
		scope    []string
		lenScope int = len(error.Scope)
	)

	switch {
	case lenScope < 10:
		scope = append(error.Scope, emptyScopes[:10-lenScope]...)

	case lenScope == 10:
		scope = error.Scope

	case lenScope > 10:
		scope = error.Scope[:9]
	}

	errorId++
	_, err := tx.Exec(sqlInsertError,
		errorId,
		error.DateTime,
		boolToYN(error.HasTimestamp),
		scope[0],
		scope[1],
		scope[2],
		scope[3],
		scope[4],
		scope[5],
		scope[6],
		scope[7],
		scope[8],
		scope[9],
		error.Message,
		error.FileName,
	)

	if err != nil {
		log.Println("Error inserting error log:", err)
	}

	return
}

// Convert a bool type to Y/N
func boolToYN(tf bool) string {
	if tf {
		return "Y"
	}
	return "N"
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
