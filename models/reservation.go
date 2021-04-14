package models

import (
	db "backend/database"
	"time"

	pg "github.com/go-pg/pg"
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
	ReservationID	string
	Guestname       string 		`json:"Guestname"`
	Numberofpeople  string 		`json:"Numberofpeople"`
	Phonenumber     string 		`json:"Phonenumber"`
	Email           string 		`json:"Email"`
	Reservationdate string 		`json:"Reservationdate"`
	Reservationtime string 		`json:"Reservationtime"`
	Tablenumber     string 		`json:"Tablenumber"`
}

//Add new reservation
func InsertReservation(r *Reservation, ref * pg.DB) (string, error) {
	res := &db.ReservationTable{
		ReservationID: r.Reservationdate + r.Guestname + r.Tablenumber,
		Guestname: r.Guestname,
		Numberofpeople: r.Numberofpeople,
		Phonenumber: r.Phonenumber,
		Email: r.Email,
		Reservationdate: r.Reservationdate,
		Reservationtime: r.Reservationtime,
		Tablenumber: r.Tablenumber,
		Rescreated: time.Now(),
	}
	result, err := res.SaveAndReturn(ref)
	if err != nil {
		return "", err
	} else {
		return result.ReservationID, nil
	}
}

//Delete existing reservation
func DeleteReservation(resid string, ref * pg.DB) (string, error) {
	
	res := &db.ReservationTable{
		ReservationID: resid,
	}

	s, err := res.Delete(ref)
	if err != nil {
		return "", err
	} else {
		return s, nil
	}
}

//Find an Existing Reservation
func GetReservationByID(resid string, ref * pg.DB) (*db.ReservationTable, error) {

	//Get a reservation by ReservationID
	res := &db.ReservationTable{
		ReservationID: resid,
	}

	r, err := res.GetRes(ref)
	if err != nil {
		return nil, err
	}else{
		return r, nil
	}
}