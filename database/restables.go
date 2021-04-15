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
type ReservationTable struct {
	tableName 		struct{} 	`sql:"reservation_table"`
	ReservationID	string 		`sql:"reservationid,pk"`
	CustomerID		string		`sql:"customerid"`
	Guestname       string 		`sql:"guestname"`
	Numberofpeople  string 		`sql:"numberofpeople"`
	Phonenumber     string 		`sql:"phonenumber"`
	Email           string 		`sql:"email"`
	Reservationdate string 		`sql:"reservationdate"`
	Reservationtime string 		`sql:"reservationtime"`
	Tablenumber     string 		`sql:"tablenumber"`
	Rescreated		time.Time	`sql:"rescreated"`
}

//Create table
func CreateReservationTable(db *pg.DB) error {
	opt := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	err := db.CreateTable(&ReservationTable{}, opt)
	if err != nil {
		log.Printf("error in reservation table creation, because : %v\n", err)
		return err
	}
	return nil
}

//Insert into database
func (r *ReservationTable) Save(db *pg.DB) error {
	err := db.Insert(r)
	if err != nil {
		log.Printf("error inserting new data into database, becase : %v\n", err)
		return err
	} else {
		log.Printf("reservation %s was inserted successfully", r.ReservationID)
		return nil
	}
}

//Insert into database and get a return value
func (r *ReservationTable) SaveAndReturn(db *pg.DB) (*ReservationTable, error) {
	result, err := db.Model(r).Returning("*").Insert()
	if err != nil {
		log.Printf("error inserting new data into database, because : %v\n", err)
		return nil, err
	} else {
		log.Println("reservation added successful")
		log.Printf("reservation details: %v\n", result)
		return r, nil
	}
}

//Delete reservation by reservation id
func (r *ReservationTable) Delete(db *pg.DB) (string, error) {
	_, err := db.Model(r).Where("reservationid = ?reservationid").Delete()
	if err != nil {
		log.Printf("error deleting reservation, because : %v\n", err)
		return "", err
	} else {
		log.Printf("the reservation %s  has been deleted sucessfully\n", r.ReservationID)
		return r.ReservationID, nil
	}
}

//Get reservation by reservation id
func (r *ReservationTable) GetRes(db *pg.DB) (*ReservationTable, error) {
	err := db.Select(r)
	if err != nil {
		log.Printf("error getting reservation by id, because : %v\n", err)
		return nil, err
	}else{
		log.Printf("get reservation successful for %v\n", r.ReservationID)
		return r, nil
	}
}

//Get reservation by reservation id
func (r *ReservationTable) GetResByCustomerID(db *pg.DB) (*ReservationTable, error) {
	err := db.Model(r).Where("customerid = ?customerid").Select()
	if err != nil {
		log.Printf("error getting reservation by id, because : %v\n", err)
		return nil, err
	}else{
		log.Printf("get reservation successful for %v\n", r.ReservationID)
		return r, nil
	}
}