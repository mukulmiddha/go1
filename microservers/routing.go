package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandlerRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/employee", GetEmployee).Methods("Get")
	r.HandleFunc("/employee/{eid}", GetEmployeeByID).Methods("Get")
	r.HandleFunc("/employee", CreateEmployee).Methods("Post")
	r.HandleFunc("/employee/{eid}", UpdateEmployee).Methods("Put")
	r.HandleFunc("/employee/{eid}", DeleteEmployee).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8080", r))

}
