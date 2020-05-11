package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type goods struct {
	Item  string  `json:"Item"`
	Price float32 `json:"Price"`
}

type allEvents []goods

var events = allEvents{
	{
		Item:  "Rice",
		Price: 35,
	},
	{
		Item:  "Pen",
		Price: 5.5,
	},
	{
		Item:  "Surf",
		Price: 10.2,
	},
}

//  To get a particular Item
func get(w http.ResponseWriter, r *http.Request) {
	productID := mux.Vars(r)["Item"]

	for _, singleEvent := range events {
		if singleEvent.Item == productID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
	w.WriteHeader(http.StatusOK)
}

//  To add new items to the database
func post(w http.ResponseWriter, r *http.Request) {
	var newGood goods
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter the item name and the price ")
	}

	json.Unmarshal(reqBody, &newGood)
	events = append(events, newGood)
	w.WriteHeader(http.StatusAccepted)

	json.NewEncoder(w).Encode(newGood)
}

//  To update the price of a particular item
func put(w http.ResponseWriter, r *http.Request) {
	productID := mux.Vars(r)["Item"]
	var updatedGood goods

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(reqBody, &updatedGood)

	for i, singleEvent := range events {
		if singleEvent.Item == productID {
			singleEvent.Price = updatedGood.Price
			events = append(events[:i], singleEvent)
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
	w.WriteHeader(http.StatusAccepted)
}

//  To delete an Item
func delete(w http.ResponseWriter, r *http.Request) {
	productID := mux.Vars(r)["Item"]

	for i, singleGood := range events {
		if singleGood.Item == productID {
			events = append(events[:i], events[i+1:]...)
			fmt.Fprintf(w, "The Good with name %v has been deleted successfully", productID)
		}
	}
	w.WriteHeader(http.StatusAccepted)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/supermarket_name/{Item}", get).Methods("GET")
	router.HandleFunc("/supermarket_name", post).Methods("POST")
	router.HandleFunc("/supermarket_name", put).Methods("PUT")
	router.HandleFunc("/supermarket_name/{Item}", delete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
