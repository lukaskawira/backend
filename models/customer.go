package models

import (
	db "backend/database"
	"time"

	pg "github.com/go-pg/pg"
)

/*
	TABLE customer_table
	FIELD			DATA
	customerid		string	email
	firstname		string
	lastname		string
	password		string
	email			string
	phonenumber		string
	datecreated		time.Time
*/
type Customer struct {
	CustomerID	string		
	FirstName   string 		`json:"Guestname"`
	Lastname  	string 		`json:"Numberofpeople"`
	Password	string		`json:"Password"`
	Email     	string 		`json:"Phonenumber"`
	Phonenumber string 		`json:"Email"`
}

//Registration
func InsertCustomer(r *Customer, ref * pg.DB) (string, error) {
	res := &db.CustomerTable{
		CustomerID: r.Email,
		Firstname: r.FirstName,
		Lastname: r.Lastname,
		Password: r.Phonenumber,
		Email: r.Email,
		Phonenumber: r.Phonenumber,
		Rescreated: time.Now(),
	}
	result, err := res.SaveAndReturn(ref)
	if err != nil {
		return "", err
	} else {
		return result.CustomerID, nil
	}
}

//Delete existing reservation
func DeleteCustomer(cusid string, ref * pg.DB) (string, error) {
	
	res := &db.CustomerTable{
		CustomerID: cusid,
	}

	s, err := res.Delete(ref)
	if err != nil {
		return "", err
	} else {
		return s, nil
	}
}

//Find an Existing Reservation
func GetCustomerByID(cusid string, ref * pg.DB) (*db.CustomerTable, error) {

	//Get a reservation by ReservationID
	res := &db.CustomerTable{
		CustomerID: cusid,
	}

	r, err := res.GetCust(ref)
	if err != nil {
		return nil, err
	}else{
		return r, nil
	}
}