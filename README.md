# plasmasword
Command line tool for importing Apache logs into an sqlite3 database (currently error logs are unsupported)

## Flags:

	Usage: plasmasword [-d sqlite3 | mysql ] [--connect string ] [-a | -e] filename ...
	
	    -d        Database driver (Defaults to sqlite3)
	    
	    --connect Database connection string (Defaults to plasmasword.db):
	              sqlite3: filename.db
	              mysql:   username:password@tcp(server:portnumber)/schema
		      
	    -a        Force loading filename as an access log
	    -e        Force loading filename as an error log
	    
	Files without an -a nor -e flag will be assumed an access log unless the filename contains the string "err"

## Prerequisites:

If you haven't already, you will need the Go (golang) toolchain installed on your machine to compile this source code: https://golang.org/

## Install instructions:

### Linux, OS X, FreeBSD:

	go get github.com/mattn/go-sqlite3	# for sqlite3 support. See "disabling drivers" for details
	go get github.com/go-sql-driver/mysql	# for MySQL support. See "disabling drivers" for details
	go install github.com/lmorg/plasmasword


### Windows install notes:
In addition to the Go language, you will need gcc installed to run `go install` against sqlite3:
https://sourceforge.net/projects/mingw-w64/?source=typ_redirect

Also you will need git installed (if it isn't already):
https://git-scm.com/download/win

Then run:

	set PATH=%PATH%;c:\Program Files\mingw-w64\x86_64-6.2.0-posix-seh-rt_v5-rev1\mingw64\bin

(where the above path is the install destination of mingw-w64)

	go get github.com/mattn/go-sqlite3	# for sqlite3 support. See "disabling drivers" for details
	go get github.com/go-sql-driver/mysql	# for MySQL support. See "disabling drivers" for details
	go install github.com/lmorg/plasmasword

## Recompiling changes to _plasmasword_:

Simply run:

	go install github.com/lmorg/plasmasword

## Disabling drivers:

Drivers can be disabled from plasmasword by setting the build ignore flag on the relevent source files. For example, to disable sqlite3 support then add the following line to the top of `sqlite3.go`, slashes and all:

	// +build ignore

Please bare in mind the default flags for `--connect` and `-d` (these can be set in `database.go`)	
