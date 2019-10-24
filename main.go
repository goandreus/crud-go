package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id.omitempty"`
	FirstName string   `json:"firstname.omitempty"`
	LastName  string   `json:"lastname.omitempty"`
	Address   *Address `json:"address.omitempty"`
}
type Address struct {
	City  string `json:"city.omitempty"`
	State string `json:"state.omitempty"`
}

var people []Person

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}
func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {

}
func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {

}
func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

func main() {
	router := mux.NewRouter()

	people = append(people, Person{ID: "1", FirstName: "Andres", LastName: "Chavez", Address: &Address{City: "San Juan de Miraflores", State: "Lima"}})
	people = append(people, Person{ID: "2", FirstName: "Carlos", LastName: "Chavez"})

	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}
