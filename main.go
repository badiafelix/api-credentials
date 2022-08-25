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
	r.HandleFunc("/signup", controllers.C_InsertNewUsers).Methods("POST")
	r.HandleFunc("/signin", controllers.C_GetUserAuth).Methods("POST")
	// r.HandleFunc("/pahlawan/{name}", controllers.C_GetSuperheroById).Methods("GET")
	// r.HandleFunc("/pahlawan", controllers.C_InsertSuperhero).Methods("POST")
	// r.HandleFunc("/pahlawan", controllers.C_UpdateSuperhero).Methods("PUT")
	// r.HandleFunc("/pahlawan/{name}", controllers.C_DeleteSuperheroById).Methods("DELETE")

	fmt.Printf("Starting server at port 8082\n")
	log.Fatal(http.ListenAndServe(":8082", r))

}
