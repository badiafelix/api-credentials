package models

import (
	//"encoding/json"
	"fmt"
	//"io/ioutil"
	"api-credentials/config"
	//"net/http"

	"gorm.io/gorm" //menggunakan gorm versi baru
)

type Users struct {
	gorm.Model
	Usr_cif       int    `json:"usr_cif"`
	Usr_name      string `json:"usr_name" validate:"required,min=3,max=50"`
	Usr_email     string `json:"usr_email"`
	Usr_phone     string `json:"usr_phone"`
	Usr_username  string `json:"usr_username"`
	Usr_password  string `json:"usr_password"`
	Usr_role      int    `json:"usr_role"`
	Usr_is_active int    `json:"usr_is_active"`
}

type Data_check struct {
	//gorm.Model
	Usr_cif          int    `json:"usr_cif"`
	Usr_name         string `json:"usr_name"`
	Usr_email        string `json:"usr_email"`
	Usr_phone        string `json:"usr_phone"`
	Usr_username     string `json:"usr_username"`
	Usr_role         int    `json:"usr_role"`
	Usr_is_active    int    `json:"usr_is_active"`
	Usr_created_date string `json:"usr_created_date"`
}

type Output struct {
	Status  string
	Message string
	Data    Data_check
}

func InsertNewUsers(paramInput Users) (*Users, bool, *gorm.DB) { // return 3 type data *superhero, isValid dan *gorm.db
	var InsertUsers Users
	var isValid bool
	config.ConnectDb()
	err := config.Db.Raw("INSERT INTO application_users(usr_cif,usr_name,usr_email,usr_phone,usr_username,usr_password,usr_role,usr_is_active)VALUES(?, ?, ?, ?, ?, ?, ?, ?);", paramInput.Usr_cif, paramInput.Usr_name, paramInput.Usr_email, paramInput.Usr_phone, paramInput.Usr_username, paramInput.Usr_password, paramInput.Usr_role, paramInput.Usr_is_active).Scan(&InsertUsers)
	if err.Error != nil {
		isValid = false
		fmt.Println("ada error")
		fmt.Printf("pesan errornya adalah : %v", err.Error.Error())
	} else {
		isValid = true
		fmt.Println("tidak ada error")
		//GetSuperheroById(paramINput.Name)
	}
	return &InsertUsers, isValid, err
}

func GetUsersById(paramName string) (*Data_check, bool, *gorm.DB) {
	var selById Data_check
	var isValid bool
	config.ConnectDb()
	err := config.Db.Raw("SELECT usr_cif, usr_name, usr_email, usr_phone, usr_username, usr_role, usr_is_active, usr_created_date FROM application_users WHERE usr_email = ?", paramName).Scan(&selById)
	if err.Error != nil {
		isValid = false
		fmt.Println("ada error")
		fmt.Printf("pesan errornya adalah : %v", err.Error.Error())
	} else {
		isValid = true
		fmt.Println("tidak ada error")
	}
	return &selById, isValid, err
}
