package controllers

import (
	db "backend/database"
	help "backend/helpers"
	bm "backend/models"
	val "backend/validations"
	"encoding/json"

	"github.com/astaxie/beego"
)

//Instantiate a new controller focusing on contact
//functions handler
type ContactController struct {
	beego.Controller
}

// @Title Insert
// @Description New Contact Response
// @Success 200 {string} Response recorded by endpoint
// @Failure 400 {string} something went error
// @router / [post]
func (c *ContactController) Insert() {
	//Initiate connection to database and store the referenced 
	//database in pg_db variable
	pg_db := db.Connect()

	//Initiate a variable data of  
	//structure of Contact object
	var data db.Contact

	//Unpack and parse the JSON that was sended by the front end
	json.Unmarshal(c.Ctx.Input.RequestBody, &data)

	//Validate the JSON data with own built function ContactValidator
	validationsError := val.ContactValidator(&data)

	//Check if there is an error on the data
	if validationsError == nil {
		//Store error value from InsertData function 
		//for further inspection
		err := bm.InsertData(&data, pg_db)

		//If there is an error
		if err != nil {
			errCode := help.ErrorCode(err.Error())
			c.Ctx.ResponseWriter.WriteHeader(errCode)
			c.Data["json"] = err.Error()
		} else {
			//If there is no error
			c.Data["json"] = "Thank you, your response has been recorded!"
		}
	} else{
		//If there is a validation error
		errCode := help.ErrorCode(validationsError.Error())
		c.Ctx.ResponseWriter.WriteHeader(errCode)
		c.Data["json"] = "JSON Parsed with error"
	}
	c.ServeJSON()
}