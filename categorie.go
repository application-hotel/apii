package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Category struct {
	Classe       string  `json:"classe"`
	TarifNormal  float64 `json:"tarif_normal"`
	TarifSpecial float64 `json:"tarif_special"`
	TarifSpe     float64 `json:"tarif_spe"`
}

// CreateCategory creates a new Category record in the database
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	// Parse the Category data from the request body
	var category Category
	err := json.NewDecoder(r.Body).Decode(&category)
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

	// Insert the Category record into the database
	stmt, err := db.Prepare("INSERT INTO categories(classe, tarifNormal, tarifSpecial, tarifSpe) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(category.Classe, category.TarifNormal, category.TarifSpecial, category.TarifSpe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusCreated)
}

// UpdateCategory updates a Category record in the database
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	// Parse the Category data from the request body
	var category Category
	err := json.NewDecoder(r.Body).Decode(&category)
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

	// Update the Category record in the database
	stmt, err := db.Prepare("UPDATE categories SET classe = ?, tarifNormal = ?, tarifSpecial = ?, tarifSpe = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(category.Classe, category.TarifNormal, category.TarifSpecial, category.TarifSpe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
}

// DeleteCategory deletes a Category record from the database
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	// Extract the ID of the Category record to delete from the URL query string
	id := r.URL.Query().Get("id")

	// Open the database connection
	db, err := sql.Open("sqlite3", "./gestionHotel")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Delete the Category record from the database
	stmt, err := db.Prepare("DELETE FROM categories WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
}
