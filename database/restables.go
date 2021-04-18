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
type Reservation struct {
	tableName 		struct{} 	`sql:"reservation_table"`
	ID				int			`sql:"id,serial,pk"`
	ReservationID	string 		`sql:"reservationid"`
	CustomerID		string		`sql:"customerid" json:"CustomerID"`
	Guestname       string 		`sql:"guestname" json:"Guestname"`
	Numberofpeople  string 		`sql:"numberofpeople" json:"Numberofpeople"`
	Phonenumber     string 		`sql:"phonenumber" json:"Phonenumber"`
	Email           string 		`sql:"email" json:"Email"`
	Reservationdate string 		`sql:"reservationdate" json:"Reservationdate"`
	Reservationtime string 		`sql:"reservationtime" json:"Reservationtime"`
	Tablenumber     string 		`sql:"tablenumber" json:"Tablenumber"`
	Rescreated		time.Time	`sql:"rescreated"`
	Status			string		`sql:"status"`

}

//Create table
func CreateReservationTable(db *pg.DB) error {
	opt := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	err := db.CreateTable(&Reservation{}, opt)
	if err != nil {
		log.Printf("error in reservation table creation, because : %v\n", err)
		return err
	}
	return nil
}

//Insert into database
func (r *Reservation) Save(db *pg.DB) error {
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
func (r *Reservation) SaveAndReturn(db *pg.DB) (*Reservation, error) {
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

//Cancel reservation by reservation id
func (r *Reservation) Cancel(db *pg.DB) (string, error) {
	_, err := db.Model(r).Set("status = ?status").Where("reservationid = ?reservationid").Update()
	if err != nil {
		log.Printf("error cancelling reservation, because : %v\n", err)
		return "", err
	} else {
		log.Printf("the reservation %s  has been canceled sucessfully\n", r.ReservationID)
		return r.ReservationID, nil
	}
}

//Get reservation by reservation id
func (r *Reservation) GetRes(db *pg.DB) (*Reservation, error) {
	err := db.Select(r)
	if err != nil {
		log.Printf("error getting reservation by id, because : %v\n", err)
		return nil, err
	}else{
		log.Printf("get reservation successful for %v\n", r.ReservationID)
		return r, nil
	}
}

//Get reservation by customer id
func (r *Reservation) GetResByCustomerID(db *pg.DB) (*Reservation, error) {
	err := db.Model(r).Where("customerid = ?customerid AND status = 'HOLD'").Select()
	if err != nil {
		log.Printf("error getting reservation by customer id, because : %v\n", err)
		return nil, err
	}else{
		log.Printf("getting reservation successful for %v\n", r)
		return r, nil
	}
}

// //Get reservation by reservation id, multiple rows
// func (r *Reservation) GetRess(db *pg.DB) ([]*Reservation, error) {
// 	result := []*Reservation{}
// 	err := db.Model(r).
// 		Where("customerid = ?customerid").
// 		ForEach(func(t *Reservation) error {
// 			log.Println(t)
// 			result = append(result, t)
// 			return nil
// 		})
// 	if err != nil {
// 		log.Printf("error getting reservation by id, because : %v\n", err)
// 		return nil, err
// 	}else{
// 		log.Printf("get reservation successful for %v\n", r.ReservationID)
// 		return result, nil
// 	}
// }

//Get table reservation with parameter table number, reservation date, and where the reservations status is on HOLD
func (t *Reservation) GetBookedTable(db *pg.DB,resdate string, tnum string) ([]*Reservation,error) {

	var result []*Reservation
	_, err := db.Query(&result, `SELECT * FROM reservation_table WHERE reservationdate IN (?) AND tablenumber IN (?) AND status = 'HOLD'`, resdate, tnum)

	if err != nil {
		log.Printf("error getting booked table because %v\n",err)
		return nil, err
	} else {
		log.Printf("get booked table successful")
		return result ,nil
	}
}