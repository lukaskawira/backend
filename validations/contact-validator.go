package validations

import (
	db "backend/database"
	"errors"
)

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
	return nil
}

