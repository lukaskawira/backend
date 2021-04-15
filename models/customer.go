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
	FirstName   string 		`json:"Firstname"`
	Lastname  	string 		`json:"Lastname"`
	Password	string		`json:"Password"`
	Email     	string 		`json:"Email"`
	Phonenumber string 		`json:"Phonenumber"`
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
		IsLogin: false,
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

//Find an Existing Customer
func GetCustomerByID(cusid string, ref * pg.DB) (*db.CustomerTable, error) {

	//Get a customer by customerID
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

//Customer login
func Login(r * Customer, ref * pg.DB) (*db.CustomerTable, error) {

	//Update login status 
	res := &db.CustomerTable{
		CustomerID: r.CustomerID,
		Password: r.Password,
		IsLogin: true,
	}

	_ , err := res.Login(ref)
	if err != nil {
		return nil, err
	} else {
		return res, nil
	}
}

//Customer logout
func Logout(cusid string, ref * pg.DB) (*db.CustomerTable, error) {

	//Update logout status 
	res := &db.CustomerTable{
		CustomerID: cusid,
		IsLogin: false,
	}
	_ , err := res.Logout(ref)
	if err != nil {
		return nil, err
	} else {
		return res, nil
	}
}