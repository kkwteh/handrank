package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/kkwteh/handrank/internal/app/sorter"
)

type SortRequest struct {
	BoardCards []string           `json:"board_cards"`
	AllHands   []sorter.HoleCards `json:"all_hands"`
}

type SortResponse struct {
	AllHands    []sorter.HoleCards `json:"all_hands"`
	HandClasses []string           `json:"hand_classes"`
}

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/healthcheck", healthCheck).Methods("GET")
	router.HandleFunc("/sortcards", sortHandler).Methods("POST", "OPTIONS")

	headersOk := handlers.AllowedHeaders([]string{"Authorization", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":8001", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Still alive!")
}

func sortHandler(w http.ResponseWriter, r *http.Request) {
	sortRequest := SortRequest{
		BoardCards: make([]string, 0),
		AllHands:   make([]sorter.HoleCards, 0),
	}
	err := json.NewDecoder(r.Body).Decode(&sortRequest)
	if err != nil {
		panic(err)
	}
	sortResponse := SortResponse{}

	sortResponse.AllHands = sorter.SortRange(sortRequest.AllHands, sortRequest.BoardCards)
	sortResponse.HandClasses = sorter.ClassifyHands(sortResponse.AllHands, sortRequest.BoardCards)

	responseJSON, err := json.Marshal(sortResponse)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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
