package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Hotel struct {
	Nom        string `json:"nom"`
	NbNiveaux  int    `json:"nb_niveaux"`
	NbChambres int    `json:"nb_chambres"`
}

// CreateHotel inserts a new Hotel record into the database
func CreateHotel(w http.ResponseWriter, r *http.Request) {
	// Parse the Hotel data from the request body
	var hotel Hotel
	err := json.NewDecoder(r.Body).Decode(&hotel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Open the database connection
	db, err := sql.Open("sqlite3", "./gestionHotel")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Insert the Hotel record into the database
	stmt, err := db.Prepare("INSERT INTO hotels(nom, nbNiveaux, nbChambres) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(hotel.Nom, hotel.NbNiveaux, hotel.NbChambres)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
}

// UpdateHotel updates an existing Hotel record in the database
func UpdateHotel(w http.ResponseWriter, r *http.Request) {
	// Parse the Hotel data from the request body
	var hotel Hotel
	err := json.NewDecoder(r.Body).Decode(&hotel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Open the database connection
	db, err := sql.Open("sqlite3", "./gestionHotel")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Update the Hotel record in the database
	stmt, err := db.Prepare("UPDATE hotels SET nom=?, nbNiveaux=?, nbChambres=? WHERE nom=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(hotel.Nom, hotel.NbNiveaux, hotel.NbChambres, hotel.Nom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
}

// DeleteHotel deletes an existing Hotel record from the database
func DeleteHotel(w http.ResponseWriter, r *http.Request) {
	// Parse the Hotel data from the request body
	var hotel Hotel
	err := json.NewDecoder(r.Body).Decode(&hotel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Open the database connection
	db, err := sql.Open("sqlite3", "./gestionHotel")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Delete the Hotel record from the database
	stmt, err := db.Prepare("DELETE FROM hotels WHERE nom=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(hotel.Nom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
}
