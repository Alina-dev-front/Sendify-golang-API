package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gopkg.in/validator.v2"
)

type Shipment struct {
	ID                   string
	SenderName           string  `validate:"min=3,max=30,regexp=^[a-zA-Z]*$"`
	SenderEmail          string  `validate:"min=6"`
	SenderAddress        string  `validate:"max=100"`
	SenderCountryCode    string  `validate:"len=2,regexp=^[A-Z]*$"`
	RecipientName        string  `validate:"min=3,max=30,regexp=^[a-zA-Z]*$"`
	RecipientEmail       string  `validate:"min=6"`
	RecipientAddress     string  `validate:"max=100"`
	RecipientCountryCode string  `validate:"len=2,regexp=^[A-Z]*$"`
	Weight               float64 `validate:"max=1000"`
	Price                string
}

func getAllShipments(w http.ResponseWriter, r *http.Request) {

	allShipments := getAllShipmentsFromDB()

	if err := json.NewEncoder(w).Encode(allShipments); err != nil {
		fmt.Println(err)
	}
}

func getShipmentByID(w http.ResponseWriter, r *http.Request) {

	shipmentID := mux.Vars(r)["id"]

	shipment := getShipmentFromDB(shipmentID)

	if err := json.NewEncoder(w).Encode(shipment); err != nil {
		fmt.Println(err)
	}
}

func addShipment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var shipment Shipment
	_ = json.NewDecoder(r.Body).Decode(&shipment)

	if errs := validator.Validate(shipment); errs != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Input invalid. Check documentation", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(shipment)

	shipment.Price = setFinalPrice(shipment.SenderCountryCode, shipment.Weight)

	insertDataInDB(shipment)
}

func main() {
	ConnectDB()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/shipment", addShipment).Methods("POST")
	router.HandleFunc("/shipments/{id}", getShipmentByID).Methods("GET")
	router.HandleFunc("/shipments", getAllShipments).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
