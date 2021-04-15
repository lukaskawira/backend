package controllers

import (
	db "backend/database"
	"backend/helpers"
	bm "backend/models"
	"backend/validations"
	"encoding/json"

	"github.com/astaxie/beego"
)

//Reservation Operations
type CustomerController struct {
	beego.Controller
}

// @Title Post
// @Description New Customer Registration
// @Success 200 {object} bm.Customer
// @router / [post]
func (c *CustomerController) Post() {
	//Connect to database
	pg_db := db.Connect()

	var data bm.Customer	//init a variable data with struct bm.customer
	json.Unmarshal(c.Ctx.Input.RequestBody, &data) 	//&data is a json object
	validationsError := validations.CustomerValidator(&data)
	if validationsError == nil {
		r, err := bm.InsertCustomer(&data, pg_db)
		if err != nil {
			errCode := helpers.ErrorCode(err.Error())
			c.Ctx.ResponseWriter.WriteHeader(errCode)
			c.Data["json"] = err.Error()
		} else {
			c.Data["json"] = "Registration successful for id " + r
		}
	}else{
		errCode := helpers.ErrorCode(validationsError.Error())
		c.Ctx.ResponseWriter.WriteHeader(errCode)
		c.Data["json"] = "Registration JSON Parse Error"
	}
	c.ServeJSON()
}

// @Title Delete
// @Description Delete Customer
// @Param	cid	path	string true	"CustomerID to be deleted"
// @Success 200 {string} Customer deleted!
// @Failure 400 cid is empty
// @router /:cid [delete]
func (c *CustomerController) Delete() {
	//Connect to database
	pg_db := db.Connect()

	cusid := c.GetString(":cid")
	if cusid != "" {
		//NOT EMPTY STRING
		r, err := bm.DeleteCustomer(cusid, pg_db)
		if err != nil {
			errCode := helpers.ErrorCode(err.Error())
			c.Ctx.ResponseWriter.WriteHeader(errCode)
			c.Data["json"] = err.Error()
		}else{
			c.Data["json"] = "Data has been deleted for customer id " + r
		}
	}
	c.ServeJSON()
}

// @Title Get
// @Description get customer data by customerid
// @Param	cid path string true "CustomerID to get"
// @Success 200 {object} bm.Customer
// @Failure 400 cid is empty
// @router /:cid [get]
func (c *CustomerController) Get() {
	//Connect to database
	pg_db := db.Connect()

	cusid := c.GetString(":cid")
	if cusid != "" {
		//NOT EMPTY STRING
		r, err := bm.GetCustomerByID(cusid,pg_db)
		if err != nil {
			errCode := helpers.ErrorCode(err.Error())
			c.Ctx.ResponseWriter.WriteHeader(errCode)
			c.Data["json"] = err.Error()
		}else{
			c.Data["json"] = r
		}
	}
	c.ServeJSON()
}

// @Title Login
// @Description Get Customer Email and Password
// @Success 200 {string} Login Successful!
// @Failure 400 Login credential invalid
// @router /login [post]
func (c *CustomerController) Login(){
	//Connect to database
	pg_db := db.Connect()

	var data bm.CustomerLogin	//init a variable data with struct bm.customer
	json.Unmarshal(c.Ctx.Input.RequestBody, &data) 	//&data is a json object
	//validationsError := validations.LoginValidator(&data)
		r, err := bm.Login(&data, pg_db)
		if err != nil {
			errCode := helpers.ErrorCode(err.Error())
			c.Ctx.ResponseWriter.WriteHeader(errCode)
			c.Data["json"] = err.Error()
		} else {
			c.Data["json"] = r
		}
	
	c.ServeJSON()
}

// @Title Logout
// @Description Get CustomerID
// @Param	cid path string true "CustomerID to get"
// @Success 200 {string} Logout Successful!
// @Failure 400 no customer logged in
// @router /logout/:cid [post]
func (c *CustomerController) Logout(){
	//Connect to database
	pg_db := db.Connect()

	cusid := c.GetString(":cid")
	if cusid != "" {
		//NOT EMPTY STRING
		r, err := bm.Logout(cusid,pg_db)
		if err != nil {
			errCode := helpers.ErrorCode(err.Error())
			c.Ctx.ResponseWriter.WriteHeader(errCode)
			c.Data["json"] = err.Error()
		}else{
			c.Data["json"] = r
		}
	}
	c.ServeJSON()
}