package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	// "sendify-api/sqldb"
)

// DB is a global variable to hold db connection
var DB *sql.DB

type Shipment struct {
	ID                   string  `json:"ID"`
	SenderName           string  `json:"SenderName"`
	SenderEmail          string  `json:"SenderEmail"`
	SenderAddress        string  `json:"SenderAddress"`
	SenderCountryCode    string  `json:"SenderCountryCode"`
	RecipientName        string  `json:"RecipientName"`
	RecipientEmail       string  `json:"RecipientEmail"`
	RecipientAddress     string  `json:"RecipientAddress"`
	RecipientCountryCode string  `json:"RecipientCountryCode"`
	Weight               float64 `json:"Weight"`
	Price                float64 `json:"Price"`
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
			fmt.Println(err)
			return
		}
		allShipments = append(allShipments, s)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
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

	insert, err := DB.Query("INSERT INTO `Shipments` (`SenderName`) VALUES('Dick Harrison')")
	if err != nil {
		panic(err)
	}

	defer insert.Close()
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {

	db, err := sql.Open("mysql", "root:root123@tcp(127.0.0.1:3306)/Sendify")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	DB = db

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/shipment", addShipment).Methods("POST")
	router.HandleFunc("/shipments/{id}", getOneShipment).Methods("GET")
	router.HandleFunc("/shipments", getAllShipments).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
