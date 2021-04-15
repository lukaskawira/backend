package database

import (
	"log"
	"time"

	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

/*
	TABLE reservation_table
	FIELD			DATA
	reservationid	string
	guestname		string
	numberofpeople	string
	phonenumber		string
	email			string
	reservationdate	string
	reservationtime	string
	tablenumber		string
	rescreated		time.Time
*/
type CustomerTable struct {
	tableName 		struct{} 	`sql:"customer_table"`
	CustomerID		string 		`sql:"custid,pk"`
	Firstname       string 		`sql:"firstname"`
	Lastname  		string 		`sql:"lastname"`
	Password	    string 		`sql:"password"`
	Email           string 		`sql:"email"`
	Phonenumber		string 		`sql:"phonenumber"`
	Rescreated		time.Time	`sql:"datecreated"`
	IsLogin			bool		`sql:"islogin"`
}

//Create customer table
func CreateCustomerTable(db *pg.DB) error {
	opt := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	err := db.CreateTable(&CustomerTable{}, opt)
	if err != nil {
		log.Printf("error in customer table creation, because : %v\n", err)
		return err
	} else {
		log.Println("customer table has been created")
		return nil
	}
}

//Insert into database
func (r *CustomerTable) Save(db *pg.DB) error {
	err := db.Insert(r)
	if err != nil {
		log.Printf("error inserting new data into database, becase : %v\n", err)
		return err
	} else {
		log.Printf("customer %s has successfully registered", r.CustomerID)
		return nil
	}
}

//Insert into database and get a return value
func (r *CustomerTable) SaveAndReturn(db *pg.DB) (*CustomerTable, error) {
	result, err := db.Model(r).Returning("*").Insert()
	if err != nil {
		log.Printf("error inserting new data into database, because : %v\n", err)
		return nil, err
	} else {
		log.Println("registration successful")
		log.Printf("customer details: %v\n", result)
		return r, nil
	}
}

//Delete data by customer id
func (r *CustomerTable) Delete(db *pg.DB) (string, error) {
	_, err := db.Model(r).Where("customerid = ?customerid").Delete()
	if err != nil {
		log.Printf("error deleting data, because : %v\n", err)
		return "", err
	} else {
		log.Printf("the customer %s  has been deleted sucessfully\n", r.CustomerID)
		return r.CustomerID, nil
	}
}

//Get customer by customer id
func (r *CustomerTable) GetCust(db *pg.DB) (*CustomerTable, error) {
	err := db.Select(r)
	if err != nil {
		log.Printf("error getting customer by id, because : %v\n", err)
		return nil, err
	}else{
		log.Printf("get customer successful for %v\n", r.CustomerID)
		return r, nil
	}
}

//Update customer login status
func (r *CustomerTable) Login(db *pg.DB) (*CustomerTable, error) {
	_ , err := db.Model(r).Set("islogin = ?islogin").Where("customerid = ?customerid AND password = ?password").Update()
	if err != nil {
		log.Printf("error updating login status, because : %v\n", err)
		return nil, err
	}else{
		log.Printf("update login status successful for customer %v\n", r.CustomerID)
		return r , nil
	}
}

//Update customer login status
func (r *CustomerTable) Logout(db *pg.DB) (*CustomerTable, error) {
	_ , err := db.Model(r).Set("islogin = ?islogin").Where("customerid = ?customerid").Update()
	if err != nil {
		log.Printf("error updating logout status, because : %v\n", err)
		return nil, err
	}else{
		log.Printf("update logout status successful for customer %v\n", r.CustomerID)
		return r , nil
	}
}

