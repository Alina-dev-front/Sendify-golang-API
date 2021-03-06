package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"sendify-api/counters"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gookit/validate"
	"github.com/gorilla/mux"
)

//Shipment is a main struct in project
type Shipment struct {
	ID                                      string
	SenderName, RecipientName               string  `validate:"required|maxLen:30|regex:^[^0-9]*$"`
	SenderEmail, RecipientEmail             string  `validate:"required|email"`
	SenderAddress, RecipientAddress         string  `validate:"required|maxLen:100"`
	SenderCountryCode, RecipientCountryCode string  `validate:"required|len:2|regex:^[A-Z]*$"`
	Weight                                  float64 `validate:"required|max:1000"`
	Price                                   string
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

	var shipment Shipment
	json.NewDecoder(r.Body).Decode(&shipment)

	v := validate.Struct(shipment)
	if !v.Validate() {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Input invalid. Check documentation", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(shipment)

	shipment.Price = counters.SetFinalPrice(shipment.SenderCountryCode, shipment.Weight)

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
