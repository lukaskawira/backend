package database

import (
	"log"
	"time"

	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

/*
	This is the detailed information of the table customer_table,
	customerid:{string,primarykey}; firstname:{string}; lastname:{string}; password:{string};
	email:{string}; phonenumber:{string}; datecreated:{time.Time(time data in utc format)};
	islogin:{bool}
*/
type Customer struct {
	tableName 		struct{} 	`sql:"customer_table"`
	CustomerID		string 		`sql:"customerid,pk"`
	Firstname       string 		`sql:"firstname" json:"Firstname"`
	Lastname  		string 		`sql:"lastname" json:"Lastname"`
	Password	    string 		`sql:"password" json:"Password"`
	Email           string 		`sql:"email" json:"Email"`
	Phonenumber		string 		`sql:"phonenumber" json:"Phonenumber"`
	Datecreated		time.Time	`sql:"datecreated"`
	IsLogin			bool		`sql:"islogin"`
}

//Create customer table
func CreateCustomerTable(db *pg.DB) error {
	opt := &orm.CreateTableOptions{
		IfNotExists: true,	//if table existed then, skip the process
	}
	err := db.CreateTable(&Customer{}, opt)
	if err != nil {
		//If creating table is unsuccessful
		log.Println("error creating customer_table")
		log.Printf("because %v\n",err)
		log.Panic(err)
		return err
	} else{
		//If success or table already exist
		return nil
	}
}

//When a new user completes the user registration form, form values
//will be submitted into the database by the following function Save()
func (c *Customer) Save(db *pg.DB) error {
	//Insert to database with reference Customer data struct
	err := db.Insert(c)

	//If error was produced by the function
	if err != nil {
		log.Println("error registering new customer data")
		log.Printf("because %v\n",err)
		log.Panic(err)
		//return the error
		return err
	} else {
		//If no error was procuced by the function, return error is nil
		log.Printf("sucessfully registered a new customer, with reference id %v\n" + c.CustomerID)
		return nil
	}
}

//Delete customer data referencing to it's customer id value
//this function returns a string containing the customer id that was recently deleted,
//and an error if the function doesn't executed correctly
func (c *Customer) Delete(db *pg.DB) (string, error) {
	//Comparing to postgres query
	//DELETE FROM customer_table WHERE customerid = (?);
	_, err := db.Model(c).Where("customerid = ?customerid").Delete()
	
	//Temporary data holder for returning response
	temp := c.CustomerID

	//If an error was produced, return empty string and the error
	if err != nil {
		log.Println("error deleting customer data for customer " + temp)
		log.Printf("because %v\n",err)
		log.Panic(err)

		//Returns emptry string and error
		return "", err
	} else {
		//If deletion is successful
		log.Printf("customer %v has been deleted successfully" + c.CustomerID)
		return "Deletion Successful", nil
	}
}

//Fetch customer data by parameter customer id, this parameter is passed from the url
//this function will return an object of pointer Customer and an error parameter
//--GetCust is an abbreviation of GetCustomer--
func (r *Customer) GetCust(db *pg.DB) (*Customer, error) {
	
	//Select from database with reference to customer id that will be passed
	//in the models function
	err := db.Select(r)

	//If error is detected in data selection
	//return value of nil and pass the error
	if err != nil {
		log.Println("error getting data from the database")
		log.Printf("because %v\n",err)
		log.Panic(err)
		return nil, err
	}else{
		//If no error is produced
		log.Printf("data successfully fetched for customer %v\n" + r.CustomerID)
		
		//Returns customer data and no error
		return r, nil
	}
}

/*
	The next 2 function below is used for updating customer active status,
	by passing the first function (login) with customer id and password as the parameter
	this will udpate the current customer with the respective id if the password is correct the status of login equals to true
	same thing with the second function (logout), by passing the customer id as the parameter
	indicating that the current customer wants to end his/her session, this will update the status of login equals to false
*/

//Update customer login status to true
//where the customer id and the password matches on what the database stored
func (c *Customer) Login(db *pg.DB) error {
	
	//Update the values in the database by parameters, in postgres query:
	//UPDATE customer_table SET islogin = (?) WHERE customerid = (?) AND password = (?);
	_ , err := db.Model(c).Set("islogin = ?islogin").Where("customerid = ?customerid AND password = ?password").Update()
	
	//If update returns an error
	if err != nil {
		log.Printf("error updating customer %v login status to the database\n" + c.CustomerID)
		log.Printf("because %v\n",err)
		log.Panic(err)
		return err
	}else{
		//If update succeeded
		log.Printf("login status updated for customer %v\n", c.CustomerID)

		//Returns nil as counter
		return nil
	}
}

//Update customer Login status to false
//Where the customer id is passed to the database
func (c *Customer) Logout(db *pg.DB) error {

	//Update the values in the database by parameters, in postgres query:
	//UPDATE customer_table SET islogin = (?) WHERE customerid = (?);
	_ , err := db.Model(c).Set("islogin = ?islogin").Where("customerid = ?customerid").Update()

	//If update returns an error
	if err != nil {
		log.Printf("error updating customer %v logout status to the database\n" + c.CustomerID)
		log.Printf("because %v\n",err)
		log.Panic(err)
		return err
	}else{
		//If update successful
		log.Printf("customer %v has successfully logged out\n", c.CustomerID)

		//Returns nil as counter
		return nil
	}
}

//This is an extra function and is unused
func (r *Customer) SaveAndReturn(db *pg.DB) (*Customer, error) {
	result, err := db.Model(r).Returning("*").Insert()
	if err != nil {
		log.Printf("error inserting new data into database, because : %v\n", err)
		return nil, err
	} else {
		log.Println("registration successful")
		log.Printf("customer details: %v\n", result)
		return r, nil
	}
}





