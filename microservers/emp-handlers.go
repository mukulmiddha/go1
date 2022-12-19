package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	file, _ := os.OpenFile("log1.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(file)
	w.Header().Set("Content-Type", "application/json")

	var emp Employee
	log.SetFlags(log.Ldate | log.Lshortfile)
	log.Println("connection to database")
	json.NewDecoder(r.Body).Decode(&emp)
	log.Println("create to database")
	Database.Create(&emp)
	json.NewEncoder(w).Encode(emp)
	file.Close()
}
func GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var empl []Employee
	log.Println("connection to database")
	log.Println("fetching  database")
	Database.Find(&empl)
	json.NewEncoder(w).Encode(empl)

}
func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var empl Employee
	log.Println("connection to database")
	log.Println("fetching  database by id")
	Database.First(&empl, mux.Vars(r)["eid"])
	json.NewEncoder(w).Encode(empl)
}
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var empl1 Employee
	log.Println("connection to database")
	Database.First(&empl1, mux.Vars(r)["eid"])
	log.Println("updating  database")
	json.NewDecoder(r.Body).Decode(&empl1)
	Database.Save(&empl1)
	json.NewEncoder(w).Encode(empl1)

}
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var empl1 Employee
	Database.Delete(&empl1, mux.Vars(r)["eid"])
	json.NewEncoder(w).Encode("employee is deleted")

}
