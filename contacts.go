package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

type Contact struct {
	Id          string `json:"Id"`
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	PhoneNumber string `json:"PhoneNumber"`
}

var Contacts []Contact
var MAXPERPAGE int = 10

func createNewContact(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var contact Contact
	json.Unmarshal(reqBody, &contact)
	Contacts = append(Contacts, contact)
	fmt.Fprintf(w, "%+v", string(reqBody))
}

func returnSingleContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, contact := range Contacts {
		if contact.Id == key {
			json.NewEncoder(w).Encode(contact)
		}
	}
}

func returnContactsPerPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageNumStr := vars["pageNum"]
	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Page must be a number")
		return
	}
	json.NewEncoder(w).Encode(Contacts[MAXPERPAGE*(pageNum-1) : MAXPERPAGE*pageNum])
}

//func add(w http.ResponseWriter, r *http.Request)
func returnFirstContacts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Contacts[:MAXPERPAGE])
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"hello world!"}`))
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/contacts", returnFirstContacts)
	myRouter.HandleFunc("/contacts/page/{pageNum}", returnContactsPerPage)
	myRouter.HandleFunc("/contact", createNewContact).Methods("POST")
	myRouter.HandleFunc("/contact/{id}", returnSingleContact)
	http.Handle("/", myRouter)
	//http.HandleFunc("/contacts", returnAllContacts)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	fmt.Println("Welcome to phone contacts")
	Contacts = []Contact{
		{Id: "1", FirstName: "John", LastName: "Doe", PhoneNumber: "0544444444"},
		{Id: "2", FirstName: "Jane", LastName: "Doe", PhoneNumber: "0588442728"},
	}

	handleRequests()
}
