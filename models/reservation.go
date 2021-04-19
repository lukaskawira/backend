package models

import (
	db "backend/database"
	"time"

	pg "github.com/go-pg/pg"
)

//Add new reservation
func InsertReservation(r *db.Reservation, ref * pg.DB) (string, error) {
	res := &db.Reservation{
		ReservationID: r.Reservationdate + r.Guestname + r.Tablenumber,
		CustomerID: r.CustomerID,
		Guestname: r.Guestname,
		Numberofpeople: r.Numberofpeople,
		Phonenumber: r.Phonenumber,
		Email: r.Email,
		Reservationdate: r.Reservationdate,
		Reservationtime: r.Reservationtime,
		Tablenumber: r.Tablenumber,
		Rescreated: time.Now(),
		Status: "HOLD",
	}
	err := res.Save(ref)
	if err != nil {
		return "", err
	} else {
		return res.ReservationID, nil
	}
}

//Delete existing reservation
func CancelReservation(resid string, ref * pg.DB) (string, error) {
	
	res := &db.Reservation{
		ReservationID: resid,
		Status: "CANCEL",
	}

	s, err := res.Cancel(ref)
	if err != nil {
		return "", err
	} else {
		return s, nil
	}
}

//Find an Existing Reservation
func GetReservationByID(resid string, ref * pg.DB) (*db.Reservation, error) {

	//Get a reservation by ReservationID
	res := &db.Reservation{
		ReservationID: resid,
	}

	r, err := res.GetRes(ref)
	if err != nil {
		return nil, err
	}else{
		return r, nil
	}
}

//Find an Existing Reservation by CustomerID
func GetReservationByCustomerID(cusid string, ref * pg.DB) (*db.Reservation, error) {

	//Get a reservation by ReservationID
	res := &db.Reservation{
		CustomerID: cusid,
	}

	r, err := res.GetResByCustomerID(ref)
	if err != nil {
		return nil, err
	}else{
		return r, nil
	}
}

func GetBookedTable(t *db.Reservation, ref *pg.DB) ([]*db.Reservation, error) {

	result, err := t.GetBookedTable(ref,t.Reservationdate, t.Tablenumber)
	if err != nil {
		return nil, err
	} else {
		return result, err
	}
}