package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {

}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
}
