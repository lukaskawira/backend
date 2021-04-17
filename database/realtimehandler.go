package database

import (
	"log"

	pg "github.com/go-pg/pg"
)

//Get table reservation by reservationd date
func (t *Reservation) GetBookedTable(db *pg.DB) (*Reservation,error) {
	err := db.Model(t).Where("reservationdate = ?reservationdate AND tablenumber = ?tablenumber").
		Select("reservationdate, reservationtime, tablenumber")
	if err != nil {
		log.Printf("error getting booked table because %v\n",err)
		return nil, err
	} else {
		return t ,nil
	}
}