/*
A simple program to create a database in rethinkdb and read a table
from it.
*/
package main

import (
	"flag"
	"log"
	"os"

	r "gopkg.in/dancannon/gorethink.v2"
)

var (
	serverAddr     = flag.String("server", "localhost:28015", "RethinkDB URL")
	databaseName   = flag.String("database", "landoop", "database name")
	createDatabase = flag.Bool("create-db", false, "create database")
	tableName      = flag.String("table", "connect_test", "table name")
	createTable    = flag.Bool("create-table", false, "create table")
	readTable      = flag.Bool("read-table", false, "select * from table")
	logFilename    = flag.String("log", "", "file to write output (and logs), stdout if left empty")
)

var session *r.Session

// Just sets up the logger in case we need to log to a file
func init() {
	flag.Parse()

	if len(*logFilename) > 0 {
		logFile, err := os.OpenFile(*logFilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Error opening log file: %v\n", err)
		}
		defer logFile.Close()
		log.SetOutput(logFile)
	} else {
		// We use the app to verify some results, thus we want the logs
		// to go to stdout
		log.SetOutput(os.Stdout)
	}

}

func main() {
	var err error

	// connect
	session, err = r.Connect(r.ConnectOpts{Address: *serverAddr})
	if err != nil {
		log.Fatalf("Error connecting to server: %v\n", err)
	}
	defer session.Close()

	// create database if asked
	if *createDatabase {
		err = DBCreate()
		if err != nil {
			log.Fatalf("Error while creating database: %v\n", err)
		}
	}

	// create table if asked
	if *createTable {
		err = TableCreate()
		if err != nil {
			log.Fatalf("Error while creating table: %v\n", err)
		}
	}

	// read table if asked
	if *readTable {
		err = SelectAll()
		if err != nil {
			log.Fatalf("Error while creating table: %v\n", err)
		}
	}
}

func DBCreate() error {
	resp, err := r.DBCreate(*databaseName).RunWrite(session)
	if err != nil {
		return err
	}
	log.Printf("%d DB created", resp.DBsCreated)
	return nil
}

func TableCreate() error {
	resp, err := r.DB(*databaseName).TableCreate(*tableName).RunWrite(session)
	if err != nil {
		return err
	}
	log.Printf("%d table created\n", resp.TablesCreated)
	return nil
}

func SelectAll() error {
	resp, err := r.DB(*databaseName).Table(*tableName).Filter(func(a interface{}) interface{} { return a }).Run(session)
	if err != nil {
		return err
	}
	defer resp.Close()

	if resp.IsNil() {
		log.Println("database didn't return any rows")
		return nil
	}

	log.Printf("%q\n", resp)
	return nil
}
