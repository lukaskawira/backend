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
