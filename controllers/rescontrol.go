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
type ReservationController struct {
	beego.Controller
}

// @Title Post
// @Description New Reservation
// @router / [post]
func (c *ReservationController) Post() {
	//Connect to database
	pg_db := db.Connect()

	var data db.Reservation 	//init a variable data with struct bm.reservation
	json.Unmarshal(c.Ctx.Input.RequestBody, &data) 	//&data is a json object
	validationsError := validations.ReservationValidator(&data)
	if validationsError == nil {
		r, err := bm.InsertReservation(&data, pg_db)
		if err != nil {
			errCode := helpers.ErrorCode(err.Error())
			c.Ctx.ResponseWriter.WriteHeader(errCode)
			c.Data["json"] = err.Error()
		} else {
			c.Data["json"] = "Reservation has been added with id " + r
		}
	}else{
		errCode := helpers.ErrorCode(validationsError.Error())
		c.Ctx.ResponseWriter.WriteHeader(errCode)
		c.Data["json"] = "Reservation Parse Error"
	}
	c.ServeJSON()
}

// @Title Delete
// @Description Delete Reservation
// @Param	rid	path	string true	"ReservationID to be deleted"
// @router /:rid [delete]
func (c *ReservationController) Delete() {
	//Connect to database
	pg_db := db.Connect()

	resid := c.GetString(":rid")
	if resid != "" {
		//NOT EMPTY STRING
		r, err := bm.CancelReservation(resid, pg_db)
		if err != nil {
			errCode := helpers.ErrorCode(err.Error())
			c.Ctx.ResponseWriter.WriteHeader(errCode)
			c.Data["json"] = err.Error()
		}else{
			c.Data["json"] = "Reservation has been deleted for reservation id " + r
		}
	}
	c.ServeJSON()
}

// @Title Get
// @Description get reservation data by reservationid
// @Param	rid path string true "ReservationID as parameter"
// @router /:rid [get]
func (c *ReservationController) Get() {
	//Connect to database
	pg_db := db.Connect()

	resid := c.GetString(":rid")
	if resid != "" {
		//NOT EMPTY STRING
		r, err := bm.GetReservationByID(resid,pg_db)
		if err != nil {
			errCode := helpers.ErrorCode(err.Error())
			c.Ctx.ResponseWriter.WriteHeader(errCode)
			c.Data["json"] = err.Error()
		}else{
			c.Data["json"] = r
		}
	} else{
		c.Data["json"] = "Fatal Error"
	}
	c.ServeJSON()
}

// @Title GetResByCustomerID
// @Description get reservation data by customerid
// @Param	cid path string true "CustomerID as parameter"
// @Success {string} Data obtained successfully
// @Failure something went wrong
// @router /gcust/:cid [post]
func (c *ReservationController) GetResByCustomerID() {
	//Connect to database
	pg_db := db.Connect()

	cusid := c.GetString(":cid")
	if cusid != "" {
		//NOT EMPTY STRING
		r, err := bm.GetReservationByCustomerID(cusid,pg_db)
		if err != nil {
			errCode := helpers.ErrorCode(err.Error())
			c.Ctx.ResponseWriter.WriteHeader(errCode)
			c.Data["json"] = err.Error()
		}else{
			c.Data["json"] = r
		}
	} else {
		c.Data["json"] = "Something Wrong"
	}
	c.ServeJSON()
}

// @Title RealtimeGet
// @Desc Get reservation data by reservation date, tablenumber, and status parameter that was sent as JSON object
// @Success 200 {object} []*db.Reservation
// @Failure 400 {string} error getting data with requested parameter
// @router /realt/ [post]
func (r *ReservationController) RealtimeGet(){
	pg_db := db.Connect()

	var data db.Reservation
	json.Unmarshal(r.Ctx.Input.RequestBody, &data)
	validationsError := validations.ValidateData(&data)

	if validationsError == nil {
		container, err := bm.GetBookedTable(&data, pg_db)
		if err != nil {
			errCode := helpers.ErrorCode(err.Error())
			r.Ctx.ResponseWriter.WriteHeader(errCode)
			r.Data["json"] = err.Error()
		} else {
			r.Data["json"] = container
		}
	} else {
		r.Data["json"] = "Data Invalid"
	}
	r.ServeJSON()
}

/* // @Title GetRess
// @Description get reservation data by customerid for multiple rows
// @Param	cid path string true "CustomerID as parameter"
// @router /gcusts/:cid [get]
func (c *ReservationController) GetRess() {
	//Connect to database
	pg_db := db.Connect()

	cusid := c.GetString(":cid")
	if cusid != "" {
		//NOT EMPTY STRING
		r, err := bm.GetRess(cusid,pg_db)
		if err != nil {
			errCode := helpers.ErrorCode(err.Error())
			c.Ctx.ResponseWriter.WriteHeader(errCode)
			c.Data["json"] = err.Error()
		}else{
			c.Data["json"] = r
		}
	}
	c.ServeJSON()
} */