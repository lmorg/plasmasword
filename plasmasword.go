package main

import (
	"github.com/lmorg/apachelogs"
	"log"
)

const Version = "0.1 ALPHA"

var completed float32

func main() {
	log.Println("Version", Version)

	Flags()

	OpenDB()

	ReadLogs()

	if err := SyncDbToDisk(); err != nil {
		log.Println(err)
	}

	CloseDB()
}

func ReadLogs() {
	for i := range fLogAccess {
		completed = ((float32(i) + 1) / float32(len(fLogAccess))) * 100
		apachelogs.ReadAccessLog(fLogAccess[i], InsertAccess, Error)
		log.Printf("%3.0f%% Loaded %s", completed, fLogAccess[i])
	}
}

func Error(err error) {
	log.Println(err)
}
