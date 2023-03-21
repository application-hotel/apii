package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Facture struct {
	ID      int     `json:"id"`
	Numero  int     `json:"numero"`
	Date    string  `json:"date"`
	Montant float64 `json:"montant"`
}

// CreateFacture inserts a new Facture record into the database
func CreateFacture(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the Facture data
	var f Facture
	err := json.NewDecoder(r.Body).Decode(&f)
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

	// Insert the Facture record into the database
	stmt, err := db.Prepare("INSERT INTO factures (numero, date, montant) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(f.Numero, f.Date, f.Montant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the ID of the newly inserted Facture record
	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the ID of the Facture object to the ID of the newly inserted record
	f.ID = int(id)

	// Marshal the Facture object to JSON
	responseJSON, err := json.Marshal(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseJSON)
}

// UpdateFacture updates an existing Facture record in the database
func UpdateFacture(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the Facture data
	var f Facture
	err := json.NewDecoder(r.Body).Decode(&f)
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

	// Update the Facture record in the database
	stmt, err := db.Prepare("UPDATE factures SET date=?, montant=? WHERE numero=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(f.Date, f.Montant, f.Numero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Marshal the updated Facture object to JSON
	responseJSON, err := json.Marshal(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
func deleteFacture(id int) error {
	// Ouvrir une connexion à la base de données
	db, err := sql.Open("postgres", "postgres://user:password@localhost/mydb?sslmode=disable")
	if err != nil {
		return err
	}
	defer db.Close()

	// Préparer la requête de suppression
	stmt, err := db.Prepare("DELETE FROM factures WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Exécuter la requête de suppression
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
