package models

import (
	//"encoding/json"
	"fmt"
	//"io/ioutil"
	"api-credentials/config"
	//"net/http"

	"gorm.io/gorm" //menggunakan gorm versi baru
)

type Login struct {
	gorm.Model
	Usr_username string `json:"usr_username"`
	Usr_password string `json:"usr_password"`
}

type UserData struct {
	//gorm.Model
	Usr_cif       int    `json:"usr_cif"`
	Usr_name      string `json:"usr_name"`
	Usr_email     string `json:"usr_email"`
	Usr_username  string `json:"usr_username"`
	Usr_role      int    `json:"usr_role"`
	Usr_is_active int    `json:"usr_is_active"`
}

type OutputContent struct {
	Status  string
	Message string
	Data    UserData
}

// type OutputError struct {
// 	Status  string
// 	Message string
// }

func GetPasswordData(usrname string) string {
	var get_pwd string
	config.ConnectDb()
	config.Db.Raw("SELECT usr_password FROM application_users WHERE usr_username = ?", usrname).Scan(&get_pwd)
	return get_pwd
}

func GetUserData(paramInput Login) (*UserData, bool, *gorm.DB) {
	var selById UserData
	var isValid bool
	config.ConnectDb()
	err := config.Db.Raw("SELECT usr_cif, usr_name, usr_email, usr_phone, usr_username, usr_password, usr_role, usr_is_active, usr_created_date FROM application_users WHERE usr_username = ?", paramInput.Usr_username).Scan(&selById)
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
