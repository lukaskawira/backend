package controllers

import (
	db "backend/database"
	help "backend/helpers"
	bm "backend/models"
	val "backend/validations"
	"encoding/json"

	"github.com/astaxie/beego"
)

//Reservation Operations
type CustomerController struct {
	beego.Controller
}

// @Title Post
// @Description Insert new customer data
// @Success 200 {string} New customer registration Successful!
// @Failure 400 {string} Registration failed, something went wrong! 
// @router / [post]
func (c *CustomerController) Post() {
	//Initiate connection to database and store the referenced 
	//database in pg_db variable
	pg_db := db.Connect()

	//Initiate a variable data of structure of Customer object
	var data db.Customer

	//Unpack and parse the JSON that was sended by the front end
	json.Unmarshal(c.Ctx.Input.RequestBody, &data)

	//Validate the JSON data with own built function CustomerValidator
	validationsError := val.CustomerValidator(&data)

	//Check if there is an error on the data
	if validationsError == nil {
		//If no error was found, 
		//insert new data into customer_table. 
		//Store the response from function InsertCustomer in r and err variable 
		//where r is the response and err is error container
		r, err := bm.InsertCustomer(&data, pg_db)

		//If there is an error
		if err != nil {
			errCode := help.ErrorCode(err.Error())
			c.Ctx.ResponseWriter.WriteHeader(errCode)
			c.Data["json"] = err.Error()
		} else {
			//If there is no error
			c.Data["json"] = "Registration successful for id " + r
		}
	} else{
		//If there is an error when validating the data
		errCode := help.ErrorCode(validationsError.Error())
		c.Ctx.ResponseWriter.WriteHeader(errCode)
		c.Data["json"] = "JSON Parsed with error"
	}
	c.ServeJSON()
}

// @Title Delete
// @Description Delete Customer by CustomerID
// @Param	cid	path string true "CustomerID to be deleted"
// @Success 200 {string} Customer has been successfully deleted
// @Failure 400 {string} No ID was present
// @router /:cid [delete]
func (c *CustomerController) Delete() {
	//Initiate connection to database and store the referenced 
	//database in pg_db variable
	pg_db := db.Connect()

	//Store the customer id from the url path referencing :cid
	cusid := c.GetString(":cid")

	//If customer id is not empty
	if cusid != "" {
		//Store output of DeleteCustomer function in r and err variable
		//where r is string and err is error container
		r, err := bm.DeleteCustomer(cusid, pg_db)

		//If there is an error
		if err != nil {
			errCode := help.ErrorCode(err.Error())
			c.Ctx.ResponseWriter.WriteHeader(errCode)
			c.Data["json"] = err.Error()
		}else{
			//If there is no error
			c.Data["json"] = "Customer " + r + " has been successfully deleted."
		}
	} else {
		c.Data["json"] = "No ID was passed on the url path"
	}
	c.ServeJSON()
}

// @Title Get
// @Description Get customer data with CustomerID
// @Param	cid path string true "CustomerID reference to get"
// @Success 200 {object} bm.Customer
// @Failure 400 {string} No ID was present
// @router /:cid [get]
func (c *CustomerController) Get() {
	//Initiate connection to database and store the referenced 
	//database in pg_db variable
	pg_db := db.Connect()

	//Store the customer id from the url path referencing :cid
	cusid := c.GetString(":cid")

	//If customer id is not empty
	if cusid != "" {
		//Store output of GetCustomerByID function in r and err variable
		//where r is Customer object and err is error container
		r, err := bm.GetCustomerByID(cusid,pg_db)

		//If there is an error
		if err != nil {
			errCode := help.ErrorCode(err.Error())
			c.Ctx.ResponseWriter.WriteHeader(errCode)
			c.Data["json"] = err.Error()
		}else{
			//If there is no error
			c.Data["json"] = r
		}
	} else {
		c.Data["json"] = "No ID was passed on the url path"
	}
	c.ServeJSON()
}

// @Title Login
// @Description Customer login with email and password
// @Success 200 {string} Login Successful!
// @Failure 400 {string} Login credential are invalid!
// @router /login/ [post]
func (c *CustomerController) Login(){
	//Initiate connection to database and store the referenced 
	//database in pg_db variable
	pg_db := db.Connect()

	//Initiate a variable data of structure of CustomerLogin object
	var data bm.CustomerLogin

	//Unpack and parse the JSON that was sended by the front end
	json.Unmarshal(c.Ctx.Input.RequestBody, &data)

	//Validate the JSON data with own built function CustomerValidator
	validationsError := val.LoginValidator(&data)

	//Check if there is an error on the data
	if validationsError == nil {
		//Send data of CustomerLogin to function login 
		//where the function returns an error that is 
		//stored in err variable
		err := bm.Login(&data, pg_db)

		//If there is an error
		if err != nil {
			errCode := help.ErrorCode(err.Error())
			c.Ctx.ResponseWriter.WriteHeader(errCode)
			c.Data["json"] = err.Error()
		} else {
			//If there is no error
			c.Data["json"] = "Login Sucessful!"
		}
	} else {
		//If validations error 
		errCode := help.ErrorCode(validationsError.Error())
		c.Ctx.ResponseWriter.WriteHeader(errCode)
		c.Data["json"] = "Login credentials contains empty field"
	}
	c.ServeJSON()
}

// @Title Logout
// @Description Logout refering to customer's email
// @Param	cid path string true "Customer email reference"
// @Success 200 {string} Logout Successful!
// @Failure 400 {string} Fatal Error, cannot logout!
// @router /logout/:cid [post]
func (c *CustomerController) Logout(){
	//Initiate connection to database and store the referenced 
	//database in pg_db variable
	pg_db := db.Connect()

	//Store the customer id from the url path referencing :cid
	cusid := c.GetString(":cid")

	//If customer id is not empty
	if cusid != "" {
		//Store error output by Logout function
		err := bm.Logout(cusid,pg_db)

		//If there is an error
		if err != nil {
			errCode := help.ErrorCode(err.Error())
			c.Ctx.ResponseWriter.WriteHeader(errCode)
			c.Data["json"] = err.Error()
		}else{
			//If there is no error
			c.Data["json"] = "Logout Successful"
		}
	} else {
		//If id is empty
		c.Data["json"] = "ID reference for logout is empty"
	}
	c.ServeJSON()
}