package models

// import (
// 	"backend/helpers"
// 	"log"

// 	"github.com/astaxie/beego/orm"
// )

// func init() {
// 	orm.RegisterModel(new(User))
// }

// type User struct {
// 	Firstname   string `json:"FirstName"`
// 	Lastname    string `json:"LastName"`
// 	Password    string `json:"Password"`
// 	Email       string `json:"Email"`
// 	Phonenumber string `json:"PhoneNumber"`
// 	Datecreated string `json:"DateCreated"`
// }

// type Reservation struct {
// 	ReservationID	string `json:"ReservationID"`
// 	Guestname       string `json:"GuestName"`
// 	Numberofpeople  string `json:"NumberOfPeople"`
// 	Phonenumber     string `json:"PhoneNumber"`
// 	Email           string `json:"Email"`
// 	Reservationdate string `json:"ReservationDate"`
// 	Reservationtime string `json:"ReservationTime"`
// 	Tablenumber     string `json:"TableNumber"`
// }

// //Register New User
// func AddUser(u User) (*User, error) {
// 	//ORM
// 	o := orm.NewOrm()

// 	//Validate Existing Username
// 	user := User{Email: u.Email}
// 	err := o.Read(&user, "email")
// 	if err == nil || err != orm.ErrNoRows {
// 		errMessage := helpers.ErrorMessage(helpers.UserExisted)
// 		return nil, errMessage
// 	}

// 	_ , err = o.Insert(&u)
// 	if err == nil {
// 		//Insert Successful
// 		return &u, nil
// 	} else {
// 		errMessage := helpers.ErrorMessage(helpers.Post)
// 		return nil, errMessage
// 	}
// }

// //Login function
// func Login(email string, pass string) (*User, error) {
// 	//ORM
// 	o := orm.NewOrm()

// 	//Compare Input Email and Password
// 	user := &User{Email: email, Password: pass}
// 	err := o.Read(&user, "email", "password")

// 	//If email or password is wrong
// 	if err != nil {
// 		if err == orm.ErrNoRows {
// 			errMessage := helpers.ErrorMessage(helpers.WrongPassword)
// 			return nil, errMessage
// 		}

// 		log.Print("[READ ERROR]: ", err)
// 		errMessage := helpers.ErrorMessage(helpers.Get)
// 		return nil, errMessage
// 	} else {
// 		return user, nil
// 	}
// }

// //Add New Reservation
// func AddReservation(res Reservation) (*Reservation, error) {
// 	//ORM
// 	o := orm.NewOrm()

// 	//Validate Existing Reservation
// 	resid := Reservation{ReservationID: res.ReservationID}
// 	err := o.Read(&resid, "reservationid")

// 	//If Reservation Data Existed
// 	if err == nil || err != orm.ErrNoRows {
// 		errMessage := helpers.ErrorMessage(helpers.ReservationExisted)
// 		return nil, errMessage
// 	}

// 	_ , err = o.Insert(&res)
// 	if err == nil {
// 		//Insertion Successful
// 		return &res, nil
// 	} else {
// 		errMessage := helpers.ErrorMessage(helpers.Post)
// 		return nil, errMessage
// 	}
// }

// //Delete Existing Reservation
// func DeleteReservation(resid string) error {
// 	//ORM
// 	o := orm.NewOrm()

// 	//Get Reservation by ReservationID
// 	//Using the function made earlier
// 	_, err := GetReservationByID(resid)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = o.Delete(&Reservation{ReservationID: resid}, "reservationid")

// 	if err != nil {
// 		log.Println("Reservation Deletion Failed")
// 		errMessage := helpers.ErrorMessage(helpers.Delete)
// 		return errMessage
// 	}
// 	return nil
// }

// func GetAllBerat() *BeratList {
// 	o := orm.NewOrm()
// 	list := &BeratList{}
// 	var dataBerat []*Berat
// 	sql := "SELECT * FROM berat ORDER BY tanggal DESC;"
// 	num, err := o.Raw(sql).QueryRows(&dataBerat)
// 	if err != nil {
// 		log.Print("error query: ", err)
// 		return nil
// 	}

// 	list.Data = dataBerat
// 	list.NumAcc = int(num)

// 	return list
// }

// func UpdateBerat(id int, uu *Berat) (u *Berat, err error) {
// 	o := orm.NewOrm()

// 	u, err = GetBerat(id)

// 	if err == nil {
// 		if uu.BeratMax != 0 {
// 			u.BeratMax = uu.BeratMax
// 		}
// 		if uu.BeratMin != 0 {
// 			u.BeratMin = uu.BeratMin
// 		}

// 		// ORM Update
// 		_, err1 := o.Update(u)

// 		if err1 == nil {
// 			//update successful
// 			return u, nil
// 		} else {
// 			return nil, helpers.ErrorMessage(helpers.Put)
// 		}
// 	} else {
// 		return nil, err
// 	}
// }

// func DeleteBerat(id int) error {
// 	o := orm.NewOrm()

// 	_, err := GetBerat(id)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = o.Delete(&Berat{Id: id}, "Id")

// 	if err != nil {
// 		log.Println("delete Berat failed")
// 		errMessage := helpers.ErrorMessage(helpers.Delete)
// 		return errMessage
// 	}
// 	return nil
// }
