package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// Client is a struct that represents the client object
type Room struct {
	Numero int    `json:"numero"`
	Etat   string `json:"etat"`
}

// CreateRoom inserts a new room record into the database
func CreateRoom(w http.ResponseWriter, r *http.Request) {
	// Parse the room data from the request body
	var room Room
	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Open the database connection
	db, err := sql.Open("sqlite3", "./hotel.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Insert the room record into the database
	stmt, err := db.Prepare("INSERT INTO chambres(numero, etat) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(room.Numero, room.Etat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
}

// UpdateRoom updates an existing room record in the database
func UpdateRoom(w http.ResponseWriter, r *http.Request) {
	// Parse the room data from the request body
	var room Room
	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Open the database connection
	db, err := sql.Open("sqlite3", "./hotel.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Update the room record in the database
	stmt, err := db.Prepare("UPDATE chambres SET etat = ? WHERE numero = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(room.Etat, room.Numero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
}

// DeleteRoom deletes an existing room record from the database
func DeleteRoom(w http.ResponseWriter, r *http.Request) {
	// Parse the room number from the request URL
	vars := mux.Vars(r)
	numero := vars["numero"]

	// Open the database connection
	db, err := sql.Open("sqlite3", "./hotel.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Delete the room record from the database
	stmt, err := db.Prepare("DELETE FROM chambres WHERE numero = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(numero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
}
