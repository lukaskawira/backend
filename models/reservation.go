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
		Reservationtime: customTimeVal(r.Reservationtime),
		Tablenumber: r.Tablenumber,
		Rescreated: time.Now(),
		Status: "HOLD",
	}
	result, err := res.SaveAndReturn(ref)
	if err != nil {
		return "", err
	} else {
		return result.ReservationID, nil
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

//Find an Existing Reservation by CustomerID
func GetRess(cusid string, ref * pg.DB) ([]*db.Reservation, error) {
	//Get a reservation by ReservationID
	res := &db.Reservation{
		CustomerID: cusid,
	}

	result, err := res.GetRess(ref)
	if err != nil {
		return nil, err
	}else{
		return result, nil
	}
}

//get reservation by d {today's date} and t {table number}
func IsBooked(d string, t string, ref *pg.DB) ([]*db.Reservation, error) {
	return nil, nil
}

//CUSTOMER TIME VALUE INPUT
func customTimeVal(t string) (string) {
	if t == "10:00" {
		return "1"
	}
	if t == "11:00" {
		return "2"
	}
	if t == "12:00" {
		return "3"
	}
	if t == "13:00" {
		return "4"
	}
	if t == "14:00" {
		return "5"
	}
	if t == "15:00" {
		return "6"
	}
	if t == "16:00" {
		return "7"
	}
	if t == "17:00" {
		return "8"
	}
	if t == "18:00" {
		return "9"
	}
	if t == "19:00" {
		return "10"
	}
	return ""
}