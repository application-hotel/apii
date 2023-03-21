package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Niveau struct {
	Numero    int `json:"numero"`
	NbChambre int `json:"nbChambre"`
}

// InsertNiveau inserts a new Niveau record into the database
func InsertNiveau(w http.ResponseWriter, r *http.Request) {
	// Parse the Niveau data from the request body
	var niveau Niveau
	err := json.NewDecoder(r.Body).Decode(&niveau)
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

	// Insert the new Niveau record into the database
	stmt, err := db.Prepare("INSERT INTO niveaux(numero, nbChambre) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(niveau.Numero, niveau.NbChambre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
}

// UpdateNiveau updates an existing Niveau record in the database
func UpdateNiveau(w http.ResponseWriter, r *http.Request) {
	// Parse the Niveau data from the request body
	var niveau Niveau
	err := json.NewDecoder(r.Body).Decode(&niveau)
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

	// Update the Niveau record in the database
	stmt, err := db.Prepare("UPDATE niveaux SET nbChambre = ? WHERE numero = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(niveau.NbChambre, niveau.Numero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
}

// DeleteNiveau deletes an existing Niveau record from the database
func DeleteNiveau(w http.ResponseWriter, r *http.Request) {
	// Parse the Niveau number from the request URL
	vars := mux.Vars(r)
	numero := vars["numero"]

	// Open the database connection
	db, err := sql.Open("sqlite3", "./gestionHotel")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Delete the Niveau record from the database
	stmt, err := db.Prepare("DELETE FROM niveaux WHERE numero = ?")
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
