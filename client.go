// Package main is the entry point for the program
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
type Client struct {
	ID        int    `json:"id"`
	Nom       string `json:"nom"`
	Prenom    string `json:"prenom"`
	Adresse   string `json:"adresse"`
	Telephone string `json:"telephone"`
}

// CreateClient inserts a new client record into the database
func CreateClient(w http.ResponseWriter, r *http.Request) {
	// Decode the request body into a Client struct
	var client Client
	err := json.NewDecoder(r.Body).Decode(&client)
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

	// Insert the new client record into the database
	stmt, err := db.Prepare("INSERT INTO clients(nom, prenom, adresse, telephone) values(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(client.Nom, client.Prenom, client.Adresse, client.Telephone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusCreated)
}

// UpdateClient updates an existing client record in the database
func UpdateClient(w http.ResponseWriter, r *http.Request) {
	// Get the client ID from the request parameters
	params := mux.Vars(r)
	clientID := params["id"]

	// Decode the request body into a Client struct
	var client Client
	err := json.NewDecoder(r.Body).Decode(&client)
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

	// Update the client record in the database
	stmt, err := db.Prepare("UPDATE clients SET nom=?, prenom=?, adresse=?, telephone=? WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(client.Nom, client.Prenom, client.Adresse, client.Telephone, clientID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
}

// DeleteClient deletes an existing client record from the database
func DeleteClient(w http.ResponseWriter, r *http.Request) {
	// Get the client ID from the request parameters
	params := mux.Vars(r)
	clientID := params["id"]

	// Open the database connection
	db, err := sql.Open("sqlite3", "./gestionHotel")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Delete the client record from the database
	stmt, err := db.Prepare("DELETE FROM clients WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(clientID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
}
