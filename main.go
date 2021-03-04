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

type shipment struct {
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
	Price                int64   `json:"Price"`
}

type allShipments []shipment

var shipments = allShipments{
	{
		ID:                   "1",
		SenderName:           "Nilsson, Tim",
		SenderEmail:          "",
		SenderAddress:        "",
		SenderCountryCode:    "",
		RecipientName:        "Jack",
		RecipientEmail:       "",
		RecipientAddress:     "",
		RecipientCountryCode: "",
		Weight:               0.451,
		Price:                0,
	},
}

func addShipment(w http.ResponseWriter, r *http.Request) {

	insert, err := DB.Query("insert into `Shipments` (`SenderName`) values('Dick Harrison')")
	if err != nil {
		panic(err)
	}

	defer insert.Close()
}

func getOneShipment(w http.ResponseWriter, r *http.Request) {
	shipmentID := mux.Vars(r)["id"]

	for _, singleShipment := range shipments {
		if singleShipment.ID == shipmentID {
			json.NewEncoder(w).Encode(singleShipment)
		}
	}
}

func getAllShipments(w http.ResponseWriter, r *http.Request) {
	getAll, err := DB.Query("SELECT * FROM `Shipments` WHERE `ID`=1")
	if err != nil {
		panic(err)
	}
	fmt.Println(getAll)
	defer getAll.Close()

	var jsonData, erro1 = json.Marshal(getAll)
	if erro1 != nil {
		log.Println(erro1)
	}

	fmt.Println(string(jsonData))
	json.NewEncoder(w).Encode(jsonData)
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
