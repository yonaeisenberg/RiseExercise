package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"strconv"

	"github.com/gorilla/mux"
)

type Contact struct {
	Id          string `json:Id,omitempty`
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	PhoneNumber string `json:"PhoneNumber"`
}

var Contacts []Contact

const (
	MAXPERPAGE int = 10
)

//var nextId int = 0

func createNewContact(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var contact Contact
	json.Unmarshal(reqBody, &contact)
	contact.Id = strconv.Itoa(1)
	if len(Contacts) > 0 {
		lastId, _ := strconv.Atoi(Contacts[len(Contacts)-1].Id)
		contact.Id = strconv.Itoa(lastId + 1)
	}
	Contacts = append(Contacts, contact)
	fmt.Fprintf(w, "New contact successfully created.")
	//fmt.Fprintf(w, "%+v", string(reqBody))
}

// func returnSingleContact(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	key, err := strconv.Atoi(vars["id"])
// 	if err != nil{

// 	}
// 	for _, contact := range Contacts {
// 		if contact.Id == key {
// 			json.NewEncoder(w).Encode(contact)
// 		}
// 	}
// }

func getNextContacts(pageNum int) []Contact {
	start := (pageNum - 1) * MAXPERPAGE
	stop := start + MAXPERPAGE

	if start >= len(Contacts) {
		return nil
	}

	if stop > len(Contacts) {
		stop = len(Contacts)
	}

	return Contacts[start:stop]
}

func returnFirstContacts(w http.ResponseWriter, r *http.Request) {
	firstContacts := getNextContacts(1)
	if firstContacts == nil {
		fmt.Fprintf(w, "No contacts in the book")
		return
	}
	json.NewEncoder(w).Encode(firstContacts)
}

func returnContactsPerPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageNumStr := vars["pageNum"]
	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Error: Page must be a number")
		return
	}
	pageContacts := getNextContacts(pageNum)
	if pageContacts == nil {
		fmt.Fprintf(w, "Page does not exist")
		return
	}
	json.NewEncoder(w).Encode(pageContacts)

	// var pageContacts []Contacts
	// for i := 0; i < MAXPERPAGE; i++ {
	// 	pageContacts = append(pageContacts, Contacts[MAXPERPAGE*(pageNum-1)+i])
	// 	if outOfBoundError != nil {

	// 		fmt.Fprintf(w, "End of contact list reached")
	// 		break
	// 	}
	//fmt.Fprintf(w, "%+v", )
	//}
}

func deleteContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	// idNum, err := strconv.Atoi(idStr)
	// if err != nil {
	// 	fmt.Println(err)
	// 	fmt.Fprintf(w, "Error: Id must be a number")
	// 	return
	// }
	for i, value := range Contacts {
		if value.Id == idStr {
			Contacts = append(Contacts[:i], Contacts[i+1:]...)
			fmt.Fprintf(w, "Contact deleted successfully")
			return
		}
	}
	fmt.Fprintf(w, "Contact with passed id not found in the book")
}

func updateContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var editedContact Contact
	json.Unmarshal(reqBody, &editedContact)
	for i, value := range Contacts {
		if value.Id == idStr {
			editedContact.Id = idStr
			Contacts = append(append(Contacts[:i], editedContact), Contacts[i+1:]...)
			fmt.Fprintf(w, "Contact edited successfully")
			return
		}
	}
	fmt.Fprintf(w, "Contact with passed id not found in the book")
}

func search(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pattern := vars["pattern"]
	result := make([]Contact, 0, len(Contacts))
	for _, value := range Contacts {
		if strings.Contains(value.FirstName, pattern) || strings.Contains(value.LastName, pattern) || strings.Contains(value.PhoneNumber, pattern) {
			result = append(result, value)
		}
	}
	json.NewEncoder(w).Encode(result)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to phone contacts!")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/contacts", returnFirstContacts)
	myRouter.HandleFunc("/contacts/page/{pageNum}", returnContactsPerPage)
	myRouter.HandleFunc("/createContact", createNewContact).Methods("POST")
	//myRouter.HandleFunc("/contact/{id}", returnSingleContact)
	myRouter.HandleFunc("/deleteContact/{id}", deleteContact)
	myRouter.HandleFunc("/updateContact/{id}", updateContact).Methods("POST")
	myRouter.HandleFunc("/search/{pattern}", search)
	http.Handle("/", myRouter)
	//http.HandleFunc("/contacts", returnAllContacts)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	fmt.Println("Welcome to phone contacts")
	Contacts = []Contact{}
	handleRequests()
}
