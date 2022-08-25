package controllers

import (
	"api-credentials/libs"
	"api-credentials/models"
	"encoding/json"
	"fmt"
	"net/http"
	//"github.com/gorilla/mux"
)

func C_InsertNewUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var InsertUsers models.Users
	var output models.Output
	var outputError libs.ErrorOutput
	_ = json.NewDecoder(r.Body).Decode(&InsertUsers)
	validasi := libs.CobaValidator(InsertUsers)
	//melakukan hashing dengan bcrypt
	hash_pwd, _ := libs.HashBcryptPassword(InsertUsers.Usr_password)
	//encode_hash := libs.Base64encode(hash_pwd)
	InsertUsers.Usr_password = hash_pwd //value password dirubah jadi bentuk hash
	if len(validasi) > 0 {
		outputError.Status = "error"
		outputError.Message = validasi
		json.NewEncoder(w).Encode(outputError) //menampilkan message error input
	} else {
		_, isValid, errorMessage := models.InsertNewUsers(InsertUsers)
		if isValid != true {
			output.Status = "error"
			output.Message = errorMessage.Error.Error()
			w.WriteHeader(409) // membuat custom response code
		} else {
			checkData, isValid, errorMessage := models.GetUsersById(InsertUsers.Usr_email)
			if isValid != true {
				fmt.Printf(errorMessage.Error.Error())
			} else {
				output.Status = "success"
				output.Message = ""
				output.Data = *checkData
				w.WriteHeader(201) // membuat custom response code
			}

		}
		json.NewEncoder(w).Encode(output)
	}

}
