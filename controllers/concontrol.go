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
type ContactController struct {
	beego.Controller
}

// @Title Insert
// @Description New Contact Response
// @Success 200 {object} db.Contact
// @Failure 400 {string} something went error
// @router / [post]
func (c *ContactController) Insert() {
	//Connect to database
	pg_db := db.Connect()

	var data db.Contact	//init a variable data with struct db.Contact
	json.Unmarshal(c.Ctx.Input.RequestBody, &data) 	//&data is a json object
	validationsError := validations.ContactValidator(&data)
	if validationsError == nil {
		err := bm.InsertData(&data, pg_db)
		if err != nil {
			errCode := helpers.ErrorCode(err.Error())
			c.Ctx.ResponseWriter.WriteHeader(errCode)
			c.Data["json"] = err.Error()
		} else {
			c.Data["json"] = "Thank you, your response has been recorded!"
		}
	} else{
		errCode := helpers.ErrorCode(validationsError.Error())
		c.Ctx.ResponseWriter.WriteHeader(errCode)
		c.Data["json"] = "Registration JSON Parse Error"
	}
	c.ServeJSON()
}