package models

import (
	db "backend/database"
	"time"

	pg "github.com/go-pg/pg"
)

//Insert contact form data to database
func InsertData(c *db.Contact, ref *pg.DB) error {
	temp := &db.Contact{
		Firstname: c.Firstname,
		Lastname: c.Lastname,
		Email: c.Email,
		Phonenumber: c.Phonenumber,
		Message: c.Message,
		Datecreated: time.Now(),
	}
	err := temp.Save(ref)
	if err!=nil{
		return err
	} else{
		return nil
	}
	
}