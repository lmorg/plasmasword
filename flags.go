package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	fLogAccess FlagStrings
	fLogError  FlagStrings
)

type FlagStrings []string

func (fs *FlagStrings) String() string         { return fmt.Sprint(*fs) }
func (fs *FlagStrings) Set(value string) error { *fs = append(*fs, value); return nil }

func Flags() {
	flag.Usage = Usage

	//flag.StringVar(&fDbFileName, "db", "plasmasword.db", "")
	flag.StringVar(&fDbDriver, "d", "sqlite3", "")
	flag.StringVar(&fDbConnStr, "connect", "plasmasword.db", "")
	flag.Var(&fLogAccess, "a", "")
	flag.Var(&fLogError, "e", "")

	flag.Parse()
	//fLogAccess = append(fLogAccess, flag.Args()...)
	for _, f := range flag.Args() {
		if strings.Contains(f, "err") {
			fLogError = append(fLogError, f)
		} else {
			fLogAccess = append(fLogAccess, f)
		}

	}

	if len(fLogAccess) == 0 && len(fLogError) == 0 {
		fmt.Println("No logs selected for parsing.")
		flag.Usage()
		os.Exit(1)
	}
}

func Usage() {
	fmt.Print(`
Usage: plasmasword [-d sqlite3 | mysql ] [--connect string ] [-a | -e] filename ...

    -d        Database driver (Defaults to sqlite3)

    --connect Database connection string (Defaults to plasmasword.db):
              sqlite3: filename.db
              mysql:   username:password@tcp(server:portnumber)/schema

    -a        Force loading filename as an access log
    -e        Force loading filename as an error log

Files without an -a nor -e flag will be assumed an access log unless the filename contains the string "err"
`)
}
