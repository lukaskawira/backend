package validations

import (
	db "backend/database"
	"errors"
)

func ValidateData(r *db.Reservation) error {
	if r.Reservationdate == "" {
		return errors.New("reservation date parameter is required")
	}
	if r.Tablenumber == "" {
		return errors.New("table number parameter is required")
	}
	return nil
}