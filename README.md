# plasmasword
Command line tool for importing Apache logs into an sqlite3 database

## Flags:

    Usage: plasmasword [--db filename] [-a | -e] filename ...
    
        --db      Sqlite3 database filename. Defaults to parkrun.db
        -a        Force loading filename as an access log
        -e        Force loading filename as an error log

## Prerequisites:

If you haven't already, you will need the Go (golang) toolchain installed on your machine to compile this source code: https://golang.org/

## Install instructions:

### Linux, OS X, FreeBSD:

	go get github.com/mattn/go-sqlite3
	go install github.com/mattn/go-sqlite3
	go install github.com/lmorg/plasmasword


### Windows install notes:
In addition to the Go language, you will need gcc installed to run `go install` against sqlite3:
https://sourceforge.net/projects/mingw-w64/?source=typ_redirect

Also you will need git installed (if it isn't already):
https://git-scm.com/download/win

Then run:

	set PATH=%PATH%;c:\Program Files\mingw-w64\x86_64-6.2.0-posix-seh-rt_v5-rev1\mingw64\bin

(where the above path is the install destination of mingw-w64)

	go get github.com/mattn/go-sqlite3
	go install github.com/mattn/go-sqlite3
	go install github.com/lmorg/plasmasword

## Recompiling changes to _plasmasword_:

Simply run:

	go install github.com/lmorg/plasmasword
