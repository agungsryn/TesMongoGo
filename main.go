package main

import (
	"log"
	"net/http"

	"be-go-mongo/db"
	"be-go-mongo/formulir"
	"be-go-mongo/user"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	route := mux.NewRouter()
	s := route.PathPrefix("/api").Subrouter() //Base Path
	_ = db.ConnectMongo()
	//Routes

	//user
	s.HandleFunc("/createUser", user.CreateUser).Methods("POST")
	s.HandleFunc("/getUsers", user.GetUsers).Methods("GET")
	s.HandleFunc("/updateUser/{id}", user.UpdateUser).Methods("PUT")
	s.HandleFunc("/deleteUser/{id}", user.DeleteUser).Methods("DELETE")
	s.HandleFunc("/getUserById/{id}", user.GetUserById).Methods("GET")

	//formulir
	s.HandleFunc("/createFormulir", formulir.CreateFormulir).Methods("POST")


	c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET","POST","PUT","HEAD","DELETE"},
        AllowCredentials: true,
    })

    handler := c.Handler(s)
	
	log.Fatal(http.ListenAndServe(":8000", handler)) // Run Server
}