package database

import (
	"log"
	"time"

	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

/*
	This is the detailed information of the table contact_table,
	a table to handle message that was sent by the customer from the front end [contact.html].

	Respective fields and data type:
	firstname:{string}; lastname:{string}; email:{string}; phonenumber:{string};
	email:{string}; datecreated:{timestampz}
*/
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
		IfNotExists: true,	//if table existed then, skip the process
	}
	err := db.CreateTable(&Contact{}, opt)
	if err != nil {
		//If creating table is unsuccessful
		log.Println("Error creating contact_table")
		log.Printf("because %v\n",err)
		log.Panic(err)
		return err
	} else{
		//If success or table already exist
		return nil
	}
}

//On Contact form submission, register the form values into database
//referencing the contact_table that was created earlier
func (c *Contact) Save(db *pg.DB) error {
	err := db.Insert(c)
	if err != nil {
		//If inserting is unsuccessful
		log.Println("Error in inserting form values into contact_table")
		log.Printf("because %v\n",err)
		log.Panic(err)
		return err
	} else {
		//If insertion is successful
		log.Println("Form value insertion is successful")
		return nil
	}
}