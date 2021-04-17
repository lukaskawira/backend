package validations

import (
	db "backend/database"
	"errors"
)

/*
	DATA VALIDATOR for Reservation struct
	Validation for a data must be inserted
	Guestname       string `json:"GuestName"`
	Numberofpeople  string `json:"NumberOfPeople"`
	Phonenumber     string `json:"PhoneNumber"`
	Email           string `json:"Email"`
	Reservationdate string `json:"ReservationDate"`
	Reservationtime string `json:"ReservationTime"`
	Tablenumber		string `json:"TableNumber"`
*/
func ReservationValidator(r *db.Reservation) error {
	if r.Guestname == "" {
		return errors.New("guestname is required")
	}
	if r.Numberofpeople == "" {
		return errors.New("numberofpeople is required")
	}
	if r.Phonenumber == "" {
		return errors.New("phonenumber is required")
	}
	if r.Email == "" {
		return errors.New("email is required")
	}
	if r.Reservationdate == "" {
		return errors.New("reservationdate is required")
	}
	if r.Reservationtime == "" {
		return errors.New("reservationtime is required")
	}
	if r.Tablenumber == "" {
		return errors.New("tablenumber is required")
	}
	return nil
}

