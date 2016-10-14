package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	fDbFileName string

	fLogAccess FlagStrings
	fLogError  FlagStrings
)

type FlagStrings []string

func (fs *FlagStrings) String() string         { return fmt.Sprint(*fs) }
func (fs *FlagStrings) Set(value string) error { *fs = append(*fs, value); return nil }

func Flags() {
	flag.Usage = Usage

	flag.StringVar(&fDbFileName, "db", "plasmasword.db", "")
	flag.Var(&fLogAccess, "a", "")
	flag.Var(&fLogError, "e", "")

	flag.Parse()
	fLogAccess = append(fLogAccess, flag.Args()...)

	if len(fLogAccess) == 0 && len(fLogError) == 0 {
		fmt.Println("No logs selected for parsing.")
		flag.Usage()
		os.Exit(1)
	}
}

func Usage() {
	fmt.Print(`
Usage: plasmasword [--db filename] [-a | -e] filename ...

    --db      Sqlite3 database filename. Defaults to plasmasword.db
    -a        Force loading filename as an access log
    -e        Force loading filename as an error log
`)
}
