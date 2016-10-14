package main

import (
	"github.com/lmorg/apachelogs"
	"log"
	"sync"
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
	var wg sync.WaitGroup

	for i := range fLogAccess {
		wg.Add(1)
		completed = ((float32(i) + 1) / float32(len(fLogAccess))) * 100
		ReadAccess(fLogAccess[i], &wg)
	}

	wg.Wait()

}

func ReadAccess(filename string, wg *sync.WaitGroup) {
	apachelogs.ReadAccessLog(filename, InsertAccess, Error)
	log.Printf("%3.0f%% Loaded %s", completed, filename)
	wg.Done()
}

func Error(err error) {
	log.Println(err)
}
