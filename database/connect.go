package database

import (
	"log"
	"os"

	pg "github.com/go-pg/pg"
)

func Connect() *pg.DB {
	opt := &pg.Options{
		User:		"postgres",
		Password: 	"root",
		Addr:		"localhost:5432",
	}

	var db *pg.DB = pg.Connect(opt)
	if db == nil {
		log.Println("failed to connect to database")
		os.Exit(100)
	}
	log.Println("connection to database successful")	//connected to database
	// CreateReservationTable(db)
	return db
}