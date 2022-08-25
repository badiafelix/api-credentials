package controllers

import (
	"api-credentials/libs"
	"api-credentials/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func C_GetUserAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var LoginUsers models.Login
	var output models.OutputContent
	var outputError libs.ErrorOutput
	_ = json.NewDecoder(r.Body).Decode(&LoginUsers)
	validasi := libs.ValidatorSignin(LoginUsers)

	if len(validasi) > 0 {
		outputError.Status = "error"
		outputError.Message = validasi
		w.WriteHeader(409)
		json.NewEncoder(w).Encode(outputError) //menampilkan message error input
	} else {
		getPwd := models.GetPasswordData(LoginUsers.Usr_username)
		// fmt.Println("password input:", LoginUsers.Usr_password)
		// fmt.Println("dapat password: ", getPwd)
		validUser := libs.CompareBcrypt([]byte(getPwd), []byte(LoginUsers.Usr_password))
		if validUser == true {
			userData, isValid, errorMessage := models.GetUserData(LoginUsers)
			if isValid != true {
				fmt.Printf(errorMessage.Error.Error())
			} else {
				output.Status = "success"
				output.Message = ""
				output.Data = *userData
				w.WriteHeader(200) // membuat custom response code
				json.NewEncoder(w).Encode(output)
			}

		} else {
			outputError.Status = "error"
			outputError.Message = "password salah"
			w.WriteHeader(409) // membuat custom response code
			json.NewEncoder(w).Encode(outputError)
		}

		//json.NewEncoder(w).Encode(output)
	}

}
