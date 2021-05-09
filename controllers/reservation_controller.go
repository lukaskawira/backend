package controllers

import (
	db "backend/database"
	help "backend/helpers"
	bm "backend/models"
	val "backend/validations"
	"encoding/json"

	"github.com/astaxie/beego"
)

//Instantiate a new controller focusing on handling reservation
type ReservationController struct {
	beego.Controller
}

// @Title Post
// @Description Insert new reservation
// @Success 200 {string} New reservation has been added to the database!
// @Failure 400 {string} Reservation failed, something is wrong!
// @router / [post]
func (c *ReservationController) Post() {
	//Initiate connection to database and store the referenced 
	//database in pg_db variable
	pg_db := db.Connect()

	//Initiate a variable data of structure of Reservation object
	var data db.Reservation
	
	//Unpack and parse the JSON that was sended by the front end
	json.Unmarshal(c.Ctx.Input.RequestBody, &data) 	//&data is a json object
	
	//Validate the JSON data with own built function ReservationValidator
	validationsError := val.ReservationValidator(&data)
	
	//Check if there is an error on the data
	if validationsError == nil {
		//If no error was found, 
		//insert new data into reservation_table. 
		//Store the response from function InsertReservation in r and err variable 
		//where r is the response and err is error container
		r, err := bm.InsertReservation(&data, pg_db)

		//If there is an error
		if err != nil {
			errCode := help.ErrorCode(err.Error())
			c.Ctx.ResponseWriter.WriteHeader(errCode)
			c.Data["json"] = err.Error()
		} else {
			//If there is no error
			c.Data["json"] = "Resevation successful for id " + r
		}
	}else{
		//If there is an error when validating the data
		errCode := help.ErrorCode(validationsError.Error())
		c.Ctx.ResponseWriter.WriteHeader(errCode)
		c.Data["json"] = "JSON Parsed with error"
	}
	c.ServeJSON()
}

// @Title Cancel()
// @Description Cancel Reservation by ReservationID
// @Success 200 {string} Reservation has been successfully canceled!
// @Failure 400 {string} No reservation ID was present
// @Param	rid	path string true "ReservationID to be canceled"
// @router /:rid [post]
func (c *ReservationController) Cancel() {
	//Initiate connection to database and store the referenced 
	//database in pg_db variable
	pg_db := db.Connect()

	//Store the customer id from the url path referencing :rid
	resid := c.GetString(":rid")

	//If reservation id is not empty
	if resid != "" {
		//Store output of CancelReservation function in r and err variable
		//where r is string and err is error container
		r, err := bm.CancelReservation(resid, pg_db)

		//If there is an error
		if err != nil {
			errCode := help.ErrorCode(err.Error())
			c.Ctx.ResponseWriter.WriteHeader(errCode)
			c.Data["json"] = err.Error()
		}else{
			//If there is no error
			c.Data["json"] = "Reservation has been canceled by reservation id " + r
		}
	} else {
		c.Data["json"] = "No reservation ID was passed on the url path"
	}
	c.ServeJSON()
}

// @Title GetResByCustomerID
// @Description Get reservation by customer id
// @Param	cid path string true "CustomerID reference to get"
// @Success {object} db.Reservation
// @Failure {string} No ID was present
// @router /gcust/:cid [post]
func (c *ReservationController) GetResByCustomerID() {
	//Initiate connection to database and store the referenced 
	//database in pg_db variable
	pg_db := db.Connect()

	//Store the customer id from the url path referencing :cid
	cusid := c.GetString(":cid")

	//If customer id is not empty
	if cusid != "" {
		//Store output of GetReservationByCustomerID function in r and err variable
		//where r is Customer object and err is error container
		r, err := bm.GetReservationByCustomerID(cusid,pg_db)

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
		c.Data["json"] = "No customer ID was passed on the url path"
	}
	c.ServeJSON()
}

// @Title RealtimeGet
// @Desc Get reservation data by table number and reservation date
// @Success 200 {object} []*db.Reservation
// @Failure 400 {string} error getting data with requested parameter
// @router /realt/ [post]
func (r *ReservationController) RealtimeGet(){
	//Initiate connection to database and store the referenced 
	//database in pg_db variable
	pg_db := db.Connect()

	//Initiate a variable data of structure of Reservation object
	var data db.Reservation

	//Unpack and parse the JSON that was sended by the front end
	json.Unmarshal(r.Ctx.Input.RequestBody, &data)

	//Validate the JSON data with own built function ValidateData
	validationsError := val.ValidateData(&data)

	//Check if there is an error on the data
	if validationsError == nil {
		//If no error found, run function GetBookedTable with reference 
		//to the data from parsed JSON and store the output in variable 
		//container and err where container is an array of reservations 
		//and err is error container
		container, err := bm.GetBookedTable(&data, pg_db)

		//If there is an error
		if err != nil {
			errCode := help.ErrorCode(err.Error())
			r.Ctx.ResponseWriter.WriteHeader(errCode)
			r.Data["json"] = err.Error()
		} else {
			//If there is no error
			r.Data["json"] = container
		}
	} else {
		r.Data["json"] = "JSON parsed with error"
	}
	r.ServeJSON()
}