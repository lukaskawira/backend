package database

import (
	"log"
	"time"

	pg "github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Contact struct {
	tableName   struct{} 	`sql:"contact_table"`
	Firstname   string   	`sql:"firstname" json:"Firstname"`
	Lastname    string   	`sql:"lastname" json:"Lastname"`
	Email       string   	`sql:"email" json:"Email"`
	Phonenumber string   	`sql:"phonenumber" json:"Phonenumber"`
	Message     string   	`sql:"message" json:"Message"`
	Datecreated time.Time	`sql:"datecreated"`
}

//Create contact table
func CreateContactTable(db *pg.DB) error {
	opt := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	err := db.CreateTable(&Contact{}, opt)
	if err != nil {
		//If unsuccessful
		log.Printf("error creating contact table because : %v\n", err)
		return err
	} else{
		//If success or table already exist
		return nil
	}
}

//On contact form submission, insert values into database
func (c *Contact) Save(db *pg.DB) error {
	err := db.Insert(c)
	if err != nil {
		log.Panic(err)
		return err
	} else {
		return nil
	}
}