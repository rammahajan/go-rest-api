package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var people []Person

func main() {

	people = append(people, Person{ID: "1", Firstname: "Ram2"})
	people = append(people, Person{ID: "2", Firstname: "Saurabh1"})
	log.Println("Starting")
	router := mux.NewRouter()
	router.HandleFunc("/", HealthCheck).Methods("GET")
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

//GetPeople mbnj
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

//HealthCheck ok
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{Status:ok}")
}

//GetPerson hh
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

//Person data structure
type Person struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
}
