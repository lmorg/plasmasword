package main

import (
	"github.com/lmorg/apachelogs"
	"log"
)

const Version = "0.4 BETA"

var completed float32

func main() {
	log.Println("Version", Version)

	Flags()
	OpenDB()
	ReadLogs()
	CloseDB()
}

func ReadLogs() {
	log.Println("Adding access table")
	for i := range fLogAccess {
		completed = ((float32(i) + 1) / float32(len(fLogAccess))) * 100
		BeginTransaction()
		apachelogs.ReadAccessLog(fLogAccess[i], InsertAccess, Error)
		CommitTransaction()
		log.Printf("%5.0f%% Loaded %s (%d records total)", completed, fLogAccess[i], accessId)
	}

	log.Println("Adding error table")
	for i := range fLogError {
		completed = ((float32(i) + 1) / float32(len(fLogError))) * 100
		BeginTransaction()
		apachelogs.ReadErrorLog(fLogError[i], InsertError, Error)
		CommitTransaction()
		log.Printf("%5.0f%% Loaded %s", completed, fLogError[i])
	}
}

func Error(err error) {
	log.Println(err)
}
