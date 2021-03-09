package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"Sendify-golang-API/counters"
	"Sendify-golang-API/data"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gookit/validate"
	"github.com/gorilla/mux"
)

//getAllShipments returns the shipments from the db
func getAllShipments(w http.ResponseWriter, r *http.Request) {

	allShipments := getAllShipmentsFromDB()

	if err := json.NewEncoder(w).Encode(allShipments); err != nil {
		fmt.Println(err)
	}
}

//getShipmentByID returns one shipment according to inserted id
func getShipmentByID(w http.ResponseWriter, r *http.Request) {

	shipmentID := mux.Vars(r)["id"]

	shipment := getShipmentFromDB(shipmentID)

	if err := json.NewEncoder(w).Encode(shipment); err != nil {
		fmt.Println(err)
	}
}

//createShipment creates a new shipment according to user input data
func createShipment(w http.ResponseWriter, r *http.Request) {

	var shipment data.Shipment
	json.NewDecoder(r.Body).Decode(&shipment)

	v := validate.Struct(shipment)
	if !v.Validate() {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Input invalid. Check documentation", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(shipment)

	shipment.Price = counters.SetFinalPrice(shipment.SenderCountryCode, shipment.Weight)

	addShipmentToDB(shipment)
}

func main() {
	ConnectDB()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/shipment", createShipment).Methods("POST")
	router.HandleFunc("/shipments/{id}", getShipmentByID).Methods("GET")
	router.HandleFunc("/shipments", getAllShipments).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
