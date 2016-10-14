package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lmorg/apachelogs"
	_ "github.com/mattn/go-sqlite3"
	"strings"
	"sync"
)

const (
	sqlCreateTable = `CREATE TABLE IF NOT EXISTS %s.access (
							id 			integer PRIMARY KEY,
							ip          string,
							method      string,
							proc    	integer,
							proto		string,
							qs      	string,
							ref         string,
							size 		integer,
							status    	integer,
							datetime    datetime,
							uri         string,
							ua    		string,
							uid   		string,
							file    	string
						);`

	sqlInsertAccess = `INSERT INTO mem.access (
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
						) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

	sqlSyncToDisk = `INSERT INTO main.access SELECT * FROM mem.access;`
)

var (
	db    *sql.DB
	mutex = &sync.Mutex{}
)

const errAlreadyExists = "already exists"

func OpenDB() {
	var err error

	log.Println("Opening database")

	db, err = sql.Open("sqlite3", fmt.Sprintf("file:%s", fDbFileName))
	if err != nil {
		log.Fatalln("Could not open database:", err)
	}

	_, err = db.Exec(fmt.Sprintf(sqlCreateTable, "main"))
	if err != nil {
		log.Fatalln("Could not create table:", err)
	}

	_, err = db.Exec(`ATTACH DATABASE ':memory:' AS mem;`)
	if err != nil {
		log.Fatalln("Could not create in memory database")
	}

	_, err = db.Exec(fmt.Sprintf(sqlCreateTable, "mem"))
	if err != nil {
		log.Fatalln("Could not create memory table:", err)
	}

	// views
	view := func(sql string) {
		_, err = db.Exec(sql)
		if err != nil && !strings.HasSuffix(err.Error(), errAlreadyExists) {
			log.Println("Could not create view:", err)
		}
	}
	view(viewAll)
	view(viewLatestNon200)
	view(viewLatestProc)
	view(viewLatest304)
	view(viewCountStatus)
	view(viewCount304)
	view(viewCountSize)
	view(viewListViews)

}

func InsertAccess(access *apachelogs.AccessLog) {
	mutex.Lock()
	_, err := db.Exec(sqlInsertAccess,
		access.IP,
		access.Method,
		access.ProcTime,
		access.Protocol,
		access.QueryString,
		access.Referrer,
		access.Size,
		access.Status.I,
		//access.Status.Title(),       // will put this in separate table
		//access.Status.Description(), // will put this in separate table
		access.DateTime,
		access.URI,
		access.UserAgent,
		access.UserID,
		access.FileName,
	)
	mutex.Unlock()

	if err != nil {
		log.Println("Error inserting access log:", err)
	}

	return
}

func SyncDbToDisk() (err error) {
	log.Println("Syncing memory to", fDbFileName)
	_, err = db.Exec(sqlSyncToDisk)
	return
}

func CloseDB() {
	db.Close()
}
