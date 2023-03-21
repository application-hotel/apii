package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Set up the router
	r := mux.NewRouter()

	// Add a route for creating a new client
	r.HandleFunc("/clients", CreateClient).Methods("POST")
	// Add a route for updating an existing client
	r.HandleFunc("/clients/{id}", UpdateClient).Methods("PUT")
	// Add a route for deleting an existing client
	r.HandleFunc("/clients/{id}", DeleteClient).Methods("DELETE")

	// chambre
	// Add a route for creating a new room
	r.HandleFunc("/chambres", CreateRoom).Methods("POST")
	// Add a route for updating an existing room
	r.HandleFunc("/chambres", UpdateRoom).Methods("PUT")
	// Add a route for deleting an existing room
	r.HandleFunc("/chambres/{numero}", DeleteRoom).Methods("DELETE")

	// hotel

	r.HandleFunc("/hotel", CreateHotel).Methods("POST")

	r.HandleFunc("/hotel", UpdateHotel).Methods("PUT")

	r.HandleFunc("/hotel/{nom}", DeleteHotel).Methods("DELETE")

	//categorie
	r.HandleFunc("/category", CreateCategory).Methods("POST")

	r.HandleFunc("/category", UpdateCategory).Methods("PUT")

	r.HandleFunc("/category/{id}", DeleteCategory).Methods("DELETE")

	//reservation
	r.HandleFunc("/reservation", CreateReservation).Methods("POST")

	r.HandleFunc("/reservation", UpdateReservation).Methods("PUT")

	r.HandleFunc("/reservation/{numero}", DeleteReservation).Methods("DELETE")

	// niveau
	r.HandleFunc("/niveau", InsertNiveau).Methods("POST")

	r.HandleFunc("/niveau", UpdateNiveau).Methods("PUT")

	r.HandleFunc("/niveau/{numero}", DeleteNiveau).Methods("DELETE")

	//facture
	r.HandleFunc("/facture", CreateFacture).Methods("POST")

	r.HandleFunc("/facture", UpdateFacture).Methods("PUT")

	// service
	r.HandleFunc("/service", CreateService).Methods("POST")

	r.HandleFunc("/service", UpdateService).Methods("PUT")

	r.HandleFunc("/service/{id}", DeleteService).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))

}
