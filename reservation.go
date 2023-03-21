package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Reservation struct {
	Numero          string `json:"numero"`
	DateEntree      string `json:"date_entree"`
	DateSortie      string `json:"date_sortie"`
	DateReservation string `json:"date_reservation"`
	Nuitee          int    `json:"nuitee"`
}

// CreateReservation inserts a new Reservation record into the database
func CreateReservation(w http.ResponseWriter, r *http.Request) {
	// Parse the Reservation data from the request body
	var reservation Reservation
	err := json.NewDecoder(r.Body).Decode(&reservation)
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

	// Insert the Reservation record into the database
	stmt, err := db.Prepare("INSERT INTO reservations(numero, dateEntree, dateSortie, dateReservation, nuitee) VALUES(?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(reservation.Numero, reservation.DateEntree, reservation.DateSortie, reservation.DateReservation, reservation.Nuitee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
}

// UpdateReservation updates an existing Reservation record in the database
func UpdateReservation(w http.ResponseWriter, r *http.Request) {
	// Parse the Reservation data from the request body
	var reservation Reservation
	err := json.NewDecoder(r.Body).Decode(&reservation)
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

	// Update the Reservation record in the database
	stmt, err := db.Prepare("UPDATE reservations SET dateEntree=?, dateSortie=?, dateReservation=?, nuitee=? WHERE numero=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(reservation.DateEntree, reservation.DateSortie, reservation.DateReservation, reservation.Nuitee, reservation.Numero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
}

// DeleteReservation deletes an existing Reservation record from the database
func DeleteReservation(w http.ResponseWriter, r *http.Request) {
	// Parse the Reservation data from the request body
	var reservation Reservation
	err := json.NewDecoder(r.Body).Decode(&reservation)
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

	// Delete the Reservation record from the database
	stmt, err := db.Prepare("DELETE FROM reservations WHERE numero=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(reservation.Numero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
}
