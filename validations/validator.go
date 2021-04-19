package validations

import (
	db "backend/database"
	"backend/models"
	"errors"
)

//A function to check whether the data that was passed
//contains empty data -> Reservation JSON
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

	//If all test(s) are passed, return no error
	return nil
}

//A function to check whether the data that was passed
//contains empty data -> RealTime JSON
func ValidateData(r *db.Reservation) error {
	if r.Reservationdate == "" {
		return errors.New("reservation date parameter is required")
	}
	if r.Tablenumber == "" {
		return errors.New("table number parameter is required")
	}

	//If all test(s) are passed, return no error
	return nil
}

//A function to check whether the data that was passed
//contains empty data -> Customer JSON
func CustomerValidator(r *db.Customer) error {
	if r.Firstname == "" {
		return errors.New("firstname is required")
	}
	if r.Lastname == "" {
		return errors.New("lastname is required")
	}
	if r.Password == "" {
		return errors.New("password is required")
	}
	if r.Email == "" {
		return errors.New("email is required")
	}
	if r.Phonenumber == "" {
		return errors.New("phonenumber is required")
	}

	//If all test(s) are passed, return no error
	return nil
}

//A function to check whether the data that was passed
//contains empty data -> Login JSON
func LoginValidator(r *models.CustomerLogin) error {
	if r.CustomerID == "" {
		return errors.New("email is required")
	}
	if r.Password == "" {
		return errors.New("password is required")
	}

	//If all test(s) are passed, return no error
	return nil
}

//A function to check whether the data that was passed
//contains empty data -> Contact JSON
func ContactValidator(c *db.Contact) error {
	if c.Firstname == "" {
		return errors.New("firstname is required")
	}
	if c.Lastname == "" {
		return errors.New("lastname is required")
	}
	if c.Email == "" {
		return errors.New("email is required")
	}
	if c.Message == "" {
		return errors.New("message is required")
	}
	if c.Phonenumber == "" {
		return errors.New("phonenumber is required")
	}

	//If all test(s) are passed, return no error
	return nil
}

