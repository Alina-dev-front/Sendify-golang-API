package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gopkg.in/validator.v2"
)

// DB is a global variable to hold db connection
var DB *sql.DB

type Shipment struct {
	ID                   string
	SenderName           string `validate:"min=3,max=30,regexp=^[a-zA-Z]*$"`
	SenderEmail          string `validate:"min=6"`
	SenderAddress        string
	SenderCountryCode    string `validate:"len=2"`
	RecipientName        string `validate:"min=3,max=30,regexp=^[a-zA-Z]*$"`
	RecipientEmail       string `validate:"min=6"`
	RecipientAddress     string
	RecipientCountryCode string `validate:"len=2"`
	Weight               float64
	Price                float64
}

func getAllShipments(w http.ResponseWriter, r *http.Request) {

	rows, err := DB.Query("SELECT * FROM `Shipments`")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var allShipments []*Shipment
	for rows.Next() {
		s := new(Shipment)
		err := rows.Scan(&s.ID, &s.SenderName, &s.SenderEmail, &s.SenderAddress, &s.SenderCountryCode, &s.RecipientName, &s.RecipientEmail, &s.RecipientAddress, &s.RecipientCountryCode, &s.Weight, &s.Price)
		if err != nil {
			return
		}
		allShipments = append(allShipments, s)
	}

	if err := rows.Err(); err != nil {
		return
	}

	if err := json.NewEncoder(w).Encode(allShipments); err != nil {
		fmt.Println(err)
	}
}

func getOneShipment(w http.ResponseWriter, r *http.Request) {

	shipmentID := mux.Vars(r)["id"]

	result, err := DB.Query("SELECT * FROM `Shipments` WHERE id = ?", shipmentID)
	if err != nil {
		panic(err)
	}

	defer result.Close()

	var s Shipment

	for result.Next() {
		err := result.Scan(&s.ID, &s.SenderName, &s.SenderEmail, &s.SenderAddress, &s.SenderCountryCode, &s.RecipientName, &s.RecipientEmail, &s.RecipientAddress, &s.RecipientCountryCode, &s.Weight, &s.Price)
		if err != nil {
			panic(err)
		}
	}

	json.NewEncoder(w).Encode(s)
}

func addShipment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var shipment Shipment
	_ = json.NewDecoder(r.Body).Decode(&shipment)

	if errs := validator.Validate(shipment); errs != nil {
		// values not valid, deal with errors here
		fmt.Println("Xuy")
		fmt.Println(errs)
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Input invalid. Check documentation", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(shipment)

	price := setFinalPrice(shipment.SenderCountryCode, shipment.Weight)

	insert, err := DB.Query("INSERT INTO Shipments(SenderName, SenderEmail, SenderAddress, SenderCountryCode, RecipientName, RecipientEmail, RecipientAddress, RecipientCountryCode, Weight, Price) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", shipment.SenderName, shipment.SenderEmail, shipment.SenderAddress, shipment.SenderCountryCode, shipment.RecipientName, shipment.RecipientEmail, shipment.RecipientAddress, shipment.RecipientCountryCode, shipment.Weight, price)
	if err != nil {
		panic(err)
	}

	defer insert.Close()
}

func main() {

	db, err := sql.Open("mysql", "root:root123@tcp(127.0.0.1:3306)/Sendify")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	DB = db

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/shipment", addShipment).Methods("POST")
	router.HandleFunc("/shipments/{id}", getOneShipment).Methods("GET")
	router.HandleFunc("/shipments", getAllShipments).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
