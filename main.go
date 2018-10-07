package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//User defines model for storing account details in database
type User struct {
	Username  string
	Password  string `json:"-"`
	IsAdmin   bool
	CreatedAt time.Time
}

type SortRequest struct {
	BoardCards []string    `json:"board_cards"`
	AllHands   [][2]string `json:"all_hands"`
}

type SortResponse struct {
	AllHands    [][2]string `json:"all_hands"`
	HandClasses []string    `json:"hand_classes"`
}

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/healthcheck", healthCheck).Methods("GET")
	router.HandleFunc("/message", handleQryMessage).Methods("GET")
	router.HandleFunc("/m/{msg}", handleUrlMessage).Methods("GET")

	router.HandleFunc("/echo", echoHandler).Methods("POST")
	router.HandleFunc("/sortcards", sortHandler).Methods("POST", "OPTIONS")

	headersOk := handlers.AllowedHeaders([]string{"Authorization", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":8001", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func handleQryMessage(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	message := vars.Get("msg")

	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func handleUrlMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	message := vars["msg"]

	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Still alive!")
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	user := User{} //initialize empty user

	//Parse json request body and use it to set fields on user
	//Note that user is passed as a pointer variable so that it's fields can be modified
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	//Set CreatedAt field on user to current local time
	user.CreatedAt = time.Now().Local()

	//Marshal or convert user object back to json and write to response
	userJson, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	//Set Content-Type header so that clients will know how to read response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//Write json response back to response
	w.Write(userJson)
}

func sortHandler(w http.ResponseWriter, r *http.Request) {
	sortRequest := SortRequest{
		BoardCards: make([]string, 0),
		AllHands:   make([][2]string, 0),
	}
	err := json.NewDecoder(r.Body).Decode(&sortRequest)
	if err != nil {
		panic(err)
	}
	sortResponse := SortResponse{
		AllHands:    make([][2]string, 0),
		HandClasses: make([]string, 0),
	}

	for _, hand := range sortRequest.AllHands {
		sortResponse.AllHands = append(sortResponse.AllHands, hand)
		sortResponse.HandClasses = append(sortResponse.HandClasses, "Pair")
	}

	responseJSON, err := json.Marshal(sortResponse)
	if err != nil {
		panic(err)
	}

	//Set Content-Type header so that clients will know how to read response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//Write json response back to response
	w.Write(responseJSON)
}

// {
// 	"board_cards": ["3d", "4d", "Th"],
// 	"all_hands": [
// 		["Ad", "As"],
// 		["Ah", "As"],
// 		["Ah", "Ad"],
// 		["Ac", "As"],
// 		["Ac", "Ad"],
// 		["Ac", "Ah"]
// 	]
// }

// {
// 	"all_hands": [
// 		["Ah", "As"],
// 		["Ah", "Ad"],
// 		["7d", "6d"],
// 		["Ac", "As"],
// 		["Ac", "Ad"],
// 		["Ac", "Ah"]
// 	],
// 	"hand_classes": ["Pair", "Pair", "High Card", "Pair", "Pair", "Pair"]
// }
