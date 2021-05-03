package database

import (
	"log"
	"time"

	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

/*
	This is the detailed information of the table reservation_table,
	ID:{id,serial,primarykey}; reservationid:{string}; customerid:{string}; guestname:{string};
	numberofpeople:{string}; phonenumber:{string}; email:{string}; reservationdate:{string};
	reservationtime:{string}; tablenumber:{string}; rescreated:{timestampz}; status:{string}

	Status is a special field with only 3 string-type
	"HOLD", "CANCEL", "DONE"
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

//Create reservation table
func CreateReservationTable(db *pg.DB) error {
	opt := &orm.CreateTableOptions{
		IfNotExists: true,	//if table already exists, skip the creation process
	}
	err := db.CreateTable(&Reservation{}, opt)
	if err != nil {
		//If table creating is unsuccessful
		log.Println("Error creating reservation_table")
		log.Printf("because %v\n",err)
		log.Panic(err)
		return err
	} else {
		//If success or table already exist
		return nil
	}
}

//When customer creates a new reservation, reservation values will
//be inserted into the database with the following function
func (r *Reservation) Save(db *pg.DB) error {
	//Insert to database referencing reservation data struct
	err := db.Insert(r)

	//If error is produced by the function
	if err != nil {
		log.Println("error making new reservation")
		log.Printf("because %v\n",err)
		log.Panic(err)

		//then return the error
		return err
	} else {
		//If there is no error
		log.Printf("reservation %s was inserted successfully", r.ReservationID)
		
		//return error as nil
		return nil
	}
}

//Update reservation_table's status to CANCEL if a customer decided to cancel
//the reservation, instead of deleting the rows, this function allows the database to
//collect any reservation that has been made by each customer
func (r *Reservation) Cancel(db *pg.DB) (string, error) {
	//Comparing to postgres query
	//UPDATE reservation_table SET status = (?) WHERE reservationid = (?)
	_, err := db.Model(r).Set("status = ?status").
		Where("reservationid = ?reservationid").
		Update()

	//Temporary data holder to return response
	temp := r.ReservationID

	//If function returns error
	if err != nil {
		log.Printf("error canceling reservation for id %v\n", temp)
		log.Printf("because %v\n", err)
		log.Panic(err)
		
		//Returns emptry string and error
		return "", err
	} else {
		//If cancelation is successful
		log.Printf("the reservation with id %s has been canceled sucessfully\n", temp)
		return temp, nil
	}
}

/*
Get reservation with parameter customer id and reservation status. 
This function focuses on restricting any customer that has an active reservation 
with the code status = 'HOLD' to not be able to make any advance booking. The status 
must be updated with either 'CANCEL' when canceled or 'DONE' when conditions are met.

This function is also used to return the form data of every reservation 
respective to their reservation id. The data is then passed to the front 
end and then the value is returned to the HTML. If a customer has made 
a reservation and the reservation holds true, the data will be displayed in reservation-details.html 
*/
func (r *Reservation) GetResByCustomerID(db *pg.DB) (*Reservation, error) {
	//Comparing to postgres query
	//SELECT * FROM reservation_table WHERE customerid = (?) AND status = 'HOLD';
	err := db.Model(r).Where("customerid = ?customerid AND status = 'HOLD'").Select()
	
	//Temporary data holder for returning response
	temp := r.CustomerID

	//If function returns error
	if err != nil {
		log.Println("no booking was being holded for customer %v\n" + temp)

		//Returns empty object indication no reservation could be selected
		//or no reservation has been made and error
		return nil, err
	}else{
		//If a customer have a booking, that booking is returned and checked
		log.Printf("there is a booking on hold for customer %v\n", temp)

		//Returns the reservation object and no error
		return r, nil
	}
}

/*
	This code below is the function to control Real-time reservation
	This function accepts 3 parameters, database pointer, a string of 
	reservation date, and a string of table number.

	So how this function works is selecting booking with table number and 
	reservation date as conditions, this will return any booking that was made 
	within the same day and the table number that reservation was made. 

	This data is then passed into the front end and the front end will update 
	the options of time by disabling the output of this function which is the 
	time where a certain table with a certain number and date was booked.
*/
func (t *Reservation) GetBookedTable(db *pg.DB,resdate string, tnum string) ([]*Reservation,error) {
	//First to store an object of arrays we must initialize a variable
	//that is an array of the object pointer Reservation
	var result []*Reservation

	//Comparing to postgres query
	//SELECT * FROM reservation_table WHERE tablenumber = (?) AND reservationdate = (?) AND status = 'HOLD';
	_, err := db.Query(&result, `SELECT * FROM reservation_table WHERE reservationdate IN (?) AND tablenumber IN (?) AND status = 'HOLD'`, resdate, tnum)

	//If the function returns an error
	if err != nil {
		log.Println("there was an error getting booked time")
		log.Printf("because %v\n",err)
		log.Panic(err)

		//Returns empty object and error
		return nil, err
	} else {
		//If the function was executed perfectly
		log.Printf("get booked table successful")

		//Returns result which is an array of Reservation
		//or row(s) of reservation and no error
		return result, nil
	}
}