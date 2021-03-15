package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"Sendify-golang-API/counters"
	"Sendify-golang-API/data"

	_ "Sendify-golang-API/docs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gookit/validate"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Sendify Shipment API
// @version 1.0
// @description API server for Shipments
// @host localhost:8080
// @Basepath /

// @Summary Get details of all shipments
// @Tags Shipments
// @Description returns array of all shipments from database
// @ID get-all-shipments
// @Produce json
// @Success 200 {array} data.Shipment
// @Router /shipments [get]
func getAllShipments(w http.ResponseWriter, r *http.Request) {

	allShipments := getAllShipmentsFromDB()

	if err := json.NewEncoder(w).Encode(allShipments); err != nil {
		fmt.Println(err)
	}
}

// @Summary Get details for a given shipmentID
// @Description Get details of shipment corresponding to inserted id
// @Tags Shipments
// @Produce  json
// @Param shipmentId path int true "ID of the shipment"
// @Success 200 {object} data.Shipment
// @Router /shipments/{shipmentId} [get]
func getShipmentByID(w http.ResponseWriter, r *http.Request) {

	shipmentID := mux.Vars(r)["id"]

	shipment := getShipmentFromDB(shipmentID)

	if err := json.NewEncoder(w).Encode(shipment); err != nil {
		fmt.Println(err)
	}
}

// @Summary Create a new shipment
// @Description Create a new shipment with the input data
// @Tags Shipments
// @Accept  json
// @Produce  json
// @Param shipment body data.Shipment true "Create shipment"
// @Success 200 {object} data.Shipment
// @Router /shipment [post]
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
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
