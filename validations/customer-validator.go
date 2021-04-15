package validations

import (
	"backend/models"
	"errors"
)

/*
	DATA VALIDATOR for Reservation struct
	Validation for a data must be inserted
	FirstName   string 		`json:"Guestname"`
	Lastname  	string 		`json:"Numberofpeople"`
	Password	string		`json:"Password"`
	Email     	string 		`json:"Phonenumber"`
	Phonenumber string 		`json:"Email"`
*/
func CustomerValidator(r *models.Customer) error {
	if r.FirstName == "" {
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
	return nil
}

func LoginValidator(r *models.CustomerLogin) error {
	if r.CustomerID == "" {
		return errors.New("email is required")
	}
	if r.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

