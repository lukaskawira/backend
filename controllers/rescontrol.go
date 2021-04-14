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

	var data bm.Reservation 	//init a variable data with struct bm.reservation
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
		r, err := bm.DeleteReservation(resid, pg_db)
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
// @Param	rid path string true "ReservationID to get"
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
	}
	c.ServeJSON()
}


// // @Title Post
// // @Description Insert Data
// // @Success 200 {object} models.Berat
// // @router / [post]
// func (u *UserController) Post() {
// 	var berat models.Berat
// 	json.Unmarshal(u.Ctx.Input.RequestBody, &berat)
// 	validationErr := validations.BeratValidation(&berat)
// 	if validationErr == nil {
// 		newAcc, err := models.AddBerat(berat)
// 		if err != nil {
// 			errCode := helpers.ErrorCode(err.Error())
// 			u.Ctx.ResponseWriter.WriteHeader(errCode)
// 			u.Data["json"] = err.Error()
// 		} else {
// 			u.Data["json"] = newAcc
// 		}
// 	} else {
// 		errCode := helpers.ErrorCode(validationErr.Error())
// 		u.Ctx.ResponseWriter.WriteHeader(errCode)
// 		u.Data["json"] = validationErr.Error()
// 	}

// 	u.ServeJSON()
// }

// // @Title GetAll
// // @Description get all Data
// // @Success 200 {object} models.BeratList
// // @router / [get]
// func (u *UserController) GetAll() {
// 	berat := models.GetAllBerat()
// 	u.Data["json"] = berat
// 	u.ServeJSON()
// }

// // @Title GetById
// // @Description get data by id
// // @Param	id		path 	int	true		"The key for staticblock"
// // @Success 200 {object} models.Berat
// // @Failure 400 :id is empty
// // @router /id/:id [get]
// func (u *UserController) GetById() {
// 	id, err := u.GetInt(":id")
// 	if err != nil {
// 		errCode := helpers.ErrorCode(err.Error())
// 		u.Ctx.ResponseWriter.WriteHeader(errCode)
// 		u.Data["json"] = err.Error()
// 	} else {
// 		if id != 0 {
// 			berat, err := models.GetBerat(id)
// 			if err != nil {
// 				errCode := helpers.ErrorCode(err.Error())
// 				u.Ctx.ResponseWriter.WriteHeader(errCode)
// 				u.Data["json"] = err.Error()
// 			} else {
// 				u.Data["json"] = berat
// 			}
// 		}
// 	}
// 	u.ServeJSON()
// }

// // @Title GetByTanggal
// // @Description get data by tanggal
// // @Param	tanggal		path 	string	true		"The key for staticblock"
// // @Success 200 {object} models.Berat
// // @Failure 400 :tanggal is empty
// // @router /:tanggal [get]
// func (u *UserController) GetByTanggal() {
// 	tanggal := u.GetString(":tanggal")
// 	if tanggal != "" {
// 		berat, err := models.GetBeratByTanggal(tanggal)
// 		if err != nil {
// 			errCode := helpers.ErrorCode(err.Error())
// 			u.Ctx.ResponseWriter.WriteHeader(errCode)
// 			u.Data["json"] = err.Error()
// 		} else {
// 			u.Data["json"] = berat
// 		}
// 	}
// 	u.ServeJSON()
// }

// // @Title Update
// // @Description update the data
// // @Param	id		path 	string	true		"The id you want to update"
// // @Param	body		body 	models.Berat	true		"body for user content"
// // @Success 200 {object} models.Berat
// // @Failure 400 :id is empty
// // @router /:id [put]
// func (u *UserController) Put() {

// 	id, err := u.GetInt(":id")
// 	if err != nil {
// 		errCode := helpers.ErrorCode(err.Error())
// 		u.Ctx.ResponseWriter.WriteHeader(errCode)
// 		u.Data["json"] = err.Error()
// 	} else {
// 		if id != 0 {
// 			var berat models.Berat
// 			json.Unmarshal(u.Ctx.Input.RequestBody, &berat)
// 			validationErr := validations.BeratValidation(&berat)
// 			if validationErr == nil {
// 				newBerat, err := models.UpdateBerat(id, &berat)
// 				if err != nil {
// 					errCode := helpers.ErrorCode(err.Error())
// 					u.Ctx.ResponseWriter.WriteHeader(errCode)
// 					u.Data["json"] = err.Error()
// 				} else {
// 					u.Data["json"] = newBerat
// 				}
// 			} else {
// 				errCode := helpers.ErrorCode(validationErr.Error())
// 				u.Ctx.ResponseWriter.WriteHeader(errCode)
// 				u.Data["json"] = validationErr.Error()
// 			}
// 		}
// 	}
// 	u.ServeJSON()
// }

// // @Title Delete
// // @Description delete the data
// // @Param	id		path 	int	true		"The id you want to delete"
// // @Success 200 {string} delete success!
// // @Failure 400 id is empty
// // @router /:id [delete]
// func (u *UserController) Delete() {

// 	id, err := u.GetInt(":id")
// 	if err != nil {
// 		errCode := helpers.ErrorCode(err.Error())
// 		u.Ctx.ResponseWriter.WriteHeader(errCode)
// 		u.Data["json"] = err.Error()
// 	} else {
// 		if id != 0 {
// 			err := models.DeleteBerat(id)
// 			if err != nil {
// 				errCode := helpers.ErrorCode(err.Error())
// 				u.Ctx.ResponseWriter.WriteHeader(errCode)
// 				u.Data["json"] = err.Error()
// 			} else {
// 				u.Data["json"] = "delete success!"
// 			}
// 		}
// 	}
// 	u.ServeJSON()
// }
