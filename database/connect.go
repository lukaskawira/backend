package database

import (
	"log"
	"os"

	pg "github.com/go-pg/pg"
)

//This function is created to initialize and create a table of each corresponding table
//that the front end will need to store some user's data
func createTable(p *pg.DB) error {

	//Create customer_table
	CreateCustomerTable(p)

	//Create reservation_table
	CreateReservationTable(p)

	//Create contact_table
	CreateContactTable(p)

	//Return error as nil, or no error
	return nil
}

func Connect() *pg.DB {
	//Initiate a connection to the database by following parameters, here I am using a GO library
	//that is specialized in connecting to a postgres database called go-pg, documentation of the library
	//is available on the import link or here github.com/go-pg/pg
	opt := &pg.Options{
		User:		"postgres",
		Password: 	"root",
		Addr:		"localhost:5432",	//a postgres port by default
	}

	//Initiate connection to database using passed parameters
	var db *pg.DB = pg.Connect(opt)
	if db == nil {
		log.Println("connection to database failed")
		os.Exit(100)
	}

	//If database is not empty then we can assume the connection is successful
	log.Println("connection to database successful")	//connected to database

	//Call a function create table
	createTable(db)

	//Return the database as pointer of *pg.DB
	return db
}