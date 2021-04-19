package models

import (
	db "backend/database"
	"time"

	pg "github.com/go-pg/pg"
)

//A struct to temporary hold customer login credentials
type CustomerLogin struct {
	CustomerID 	string		`json:"CustomerID"`
	Password 	string		`json:"Password"`
}

//Customer registration function
func InsertCustomer(r *db.Customer, ref * pg.DB) (string, error) {
	
	//Assign each value corresponds to the Customer object that was passed
	//by the function as parameter
	res := &db.Customer{
		CustomerID: r.Email,
		Firstname: r.Firstname,
		Lastname: r.Lastname,
		Password: r.Password,
		Email: r.Email,
		Phonenumber: r.Phonenumber,
		Datecreated: time.Now(),
		IsLogin: false,	//Always initate the status of a new user as false
	}

	//Check if the function returns an error
	err := res.Save(ref)

	//If error is found
	if err != nil {
		s := "Unfortunately, an error has occurred."
		return s, err
	} else {
		//Return successful statement as string
		s := "Sucessfully registered a new customer, new customer with id " + r.CustomerID 
		return s, nil
	}
}

//Delete a customer by customer id parameter
func DeleteCustomer(id string, ref * pg.DB) (string, error) {
	
	//Assign customer id value by passed string id
	res := &db.Customer{
		CustomerID: id,
	}

	//Assign str string value with the output from calling function delete
	str, err := res.Delete(ref)
	
	//Check if there is an error
	if err != nil {
		//If error occured
		s := "Unfortunately, an error has occurred."

		//Return string and error
		return s, err
	} else {
		//Else if success return str string and return no error
		return str, nil
	}
}

//Find existing customer by parameter customer id
func GetCustomerByID(id string, ref * pg.DB) (*db.Customer, error) {

	//Assign customer id with passed id string
	res := &db.Customer{
		CustomerID: id,
	}

	//Assign obj customer object and var error with output from function GetCust from database
	obj, err := res.GetCust(ref)

	//Check if there is an error
	if err != nil {
		//If error exists, return no object and the error
		return nil, err
	}else{
		//If no error existed, return object and no error
		return obj, nil
	}
}

//Customer login function, accepting a Login JSON object that will be 
//parsed to a CustomerLogin object
func Login(r * CustomerLogin, ref * pg.DB) error {

	//Assign values by parsing r CustomerLogin object 
	res := &db.Customer{
		CustomerID: r.CustomerID,
		Password: r.Password,
		IsLogin: true,	//Assing islogin status as true when customer are logged in
	}

	//Check if there is an error with the function output
	err := res.Login(ref)

	//If error was produced
	if err != nil {
		//Returns the error output
		return err
	} else {
		//Return no error
		return nil
	}
}

//Customer logout function, by passing the reference id with string
//the function will then parsed the reference id to the input
//of the logout inner function
func Logout(id string, ref * pg.DB) error {

	//Assign values obtained from the parameter input
	//and update the following fields
	res := &db.Customer{
		CustomerID: id,
		IsLogin: false, //Assign islogin status as false to indicate that the user is currently logged out
	}

	//Check if there is an error output returned from the function
	err := res.Logout(ref)

	//If there is an error
	if err != nil {
		//Returns the error
		return err
	} else {
		//Returns no error
		return nil
	}
}