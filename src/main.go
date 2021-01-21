package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the pimento paradise 🌶")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    // New instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc(apiURL, homePage)
	myRouter.HandleFunc(fmt.Sprintf("%s/chillies", apiURL), returnAllChillies)
	myRouter.HandleFunc(fmt.Sprintf("%s/chilli", apiURL), createNewChili).Methods("POST")
    myRouter.HandleFunc(fmt.Sprintf("%s/chillies/{id}", apiURL), returnSingleChili)
    log.Fatal(http.ListenAndServe(":667", myRouter))
}

func main() {
	fmt.Println("Go Pimento API v1.0")
	Chillies = []Chilli{
		{ID:"1", Name:"Pimiento", Species:"Capsicum annuum", ScovilleScale:[2]int{100,500}},
		{ID:"2", Name:"Ghost pepper", Species:"Unknown", ScovilleScale:[2]int{100}},
	}
	handleRequests()
}

func returnAllChillies(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Chillies)
}

func returnSingleChili(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]

for _, chili := range Chillies {
	if chili.ID == key {
		fmt.Println("Endpoint Hit: Return", chili.ID)
		json.NewEncoder(w).Encode(chili)
	}
}
}

func createNewChili(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))
}

const apiURL string = "/api/v1"

// Chilli type 
type Chilli struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Species string `json:"species"`
	ScovilleScale [2]int `json:"scoville_scale"`
}

// Chillies array 
var Chillies []Chilli;