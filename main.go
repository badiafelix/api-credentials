package main

import (
	"api-credentials/controllers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	//menggunakan gorm versi baru
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/users/signup", controllers.C_InsertNewUsers).Methods("POST")
	r.HandleFunc("/users/login", controllers.C_GetUserAuth).Methods("POST")
	// r.HandleFunc("/pahlawan/{name}", controllers.C_Getsomething).Methods("GET")

	fmt.Printf("Starting server at port 8082\n")
	log.Fatal(http.ListenAndServe(":8082", r))

}
