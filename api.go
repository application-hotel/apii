package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Client struct {
	ID        int    `json:"id"`
	Nom       string `json:"nom"`
	Prenom    string `json:"prenom"`
	Adresse   string `json:"adresse"`
	Telephone string `json:"telephone"`
}

var clients []Client

func createClient(w http.ResponseWriter, r *http.Request) {
	var newClient Client
	err := json.NewDecoder(r.Body).Decode(&newClient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	clients = append(clients, newClient)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newClient)
}

func updateClient(w http.ResponseWriter, r *http.Request) {
	// extract client ID from URL
	urlSegments := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(urlSegments[len(urlSegments)-1])
	if err != nil {
		http.Error(w, "Invalid client ID", http.StatusBadRequest)
		return
	}

	// find client in list by ID
	var foundClient *Client
	for i := range clients {
		if clients[i].ID == id {
			foundClient = &clients[i]
			break
		}
	}
	if foundClient == nil {
		http.Error(w, "Client not found", http.StatusNotFound)
		return
	}

	// update client with new data
	var updatedClient Client
	err = json.NewDecoder(r.Body).Decode(&updatedClient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	foundClient.Nom = updatedClient.Nom
	foundClient.Prenom = updatedClient.Prenom
	foundClient.Adresse = updatedClient.Adresse
	foundClient.Telephone = updatedClient.Telephone

	// return updated client as JSON
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(foundClient)
}

func deleteClient(w http.ResponseWriter, r *http.Request) {
	// extract client ID from URL
	urlSegments := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(urlSegments[len(urlSegments)-1])
	if err != nil {
		http.Error(w, "Invalid client ID", http.StatusBadRequest)
		return
	}

	// find client in list by ID
	var foundClientIndex int = -1
	for i := range clients {
		if clients[i].ID == id {
			foundClientIndex = i
			break
		}
	}
	if foundClientIndex == -1 {
		http.Error(w, "Client not found", http.StatusNotFound)
		return
	}

	// remove client from list
	clients = append(clients[:foundClientIndex], clients[foundClientIndex+1:]...)

	// return success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Client deleted successfully"})
}

type Chambre struct {
	Numero string `json:"numero"`
	Etat   string `json:"etat"`
}

type ChambreList struct {
	Chambre []Chambre `json:"chambres"`
}

var chambreList ChambreList

func createChambreList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&chambreList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(chambreList)
}

func updateRoomList(w http.ResponseWriter, r *http.Request) {
	// Lecture du corps de la requête HTTP entrante
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Définition du type de données de la requête HTTP
	var updateRequest struct {
		Chambres []Chambre `json:"rooms"`
	}

	// Décodage du corps de la requête en JSON
	if err := json.Unmarshal(body, &updateRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var chambres []Chambre
	// Modification de la liste de chambres
	// Ici, nous supposons que la liste de chambres est stockée dans une variable globale appelée `chambres`
	chambres = updateRequest.Chambres

	// Encodage de la réponse en JSON et envoi de la réponse HTTP
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ChambreList{Chambre: chambres})
}

var chambres []Chambre

func deleteRoomList(w http.ResponseWriter, r *http.Request) {
	// Suppression de la liste de chambres
	// Ici, nous supposons que la liste de chambres est stockée dans une variable globale appelée `chambres`
	chambres = []Chambre{}

	// Encodage de la réponse en JSON et envoi de la réponse HTTP
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ChambreList{Chambre: chambres})
}

type Niveau struct {
	Numero    int `json:"numero"`
	Nbchambre int `json:"nb_chambre"`
}

var niveaux []Niveau

func createNiveau(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse le corps de la requête pour obtenir les données du nouveau niveau
	var newNiveau Niveau
	err := json.NewDecoder(r.Body).Decode(&newNiveau)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ajouter le nouveau niveau à la liste de niveaux
	niveaux = append(niveaux, newNiveau)

	// Encoder le nouveau niveau en JSON et renvoyer la réponse
	json.NewEncoder(w).Encode(newNiveau)
}
func updateNiveau(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Récupérer le numéro de niveau à modifier depuis les paramètres de la requête
	numero, err := strconv.Atoi(r.URL.Query().Get("numero"))
	if err != nil {
		http.Error(w, "Invalid niveau number", http.StatusBadRequest)
		return
	}

	// Rechercher le niveau correspondant dans la liste de niveaux
	index := -1
	for i, niveau := range niveaux {
		if niveau.Numero == numero {
			index = i
			break
		}
	}

	// Si le niveau n'existe pas, renvoyer une erreur
	if index == -1 {
		http.Error(w, "Niveau not found", http.StatusNotFound)
		return
	}

	// Analyser le corps de la requête pour obtenir les nouvelles données du niveau
	var updatedNiveau Niveau
	err = json.NewDecoder(r.Body).Decode(&updatedNiveau)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Modifier le niveau correspondant dans la liste de niveaux
	niveaux[index] = updatedNiveau

	// Encoder le niveau mis à jour en JSON et renvoyer la réponse
	json.NewEncoder(w).Encode(updatedNiveau)
}
func deleteNiveaux(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Supprimer tous les niveaux de la liste de niveaux
	niveaux = nil

	// Encoder la liste vide de niveaux en JSON et renvoyer la réponse
	json.NewEncoder(w).Encode(niveaux)
}

type Reservation struct {
	Numero          string `json:"numero"`
	DateEntree      string `json:"date_entree"`
	DateSortie      string `json:"date_sortie"`
	DateReservation string `json:"date_reservation"`
	Nuitee          int    `json:"nuitee"`
}

var reservations []Reservation

func CreateReservation(w http.ResponseWriter, r *http.Request) {
	var reservation Reservation
	err := json.NewDecoder(r.Body).Decode(&reservation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	reservations = append(reservations, reservation)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reservation)
}
func UpdateReservations(w http.ResponseWriter, r *http.Request) {
	var newReservations []Reservation
	err := json.NewDecoder(r.Body).Decode(&newReservations)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	reservations = newReservations

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reservations)
}
func DeleteReservations(w http.ResponseWriter, r *http.Request) {
	reservations = nil

	w.WriteHeader(http.StatusNoContent)
}

type Hotel struct {
	Nom        string `json:"nom"`
	NbNiveaux  int    `json:"nb_niveaux"`
	NbChambres int    `json:"nb_chambres"`
}

var hotels []Hotel

func CreateHotel(w http.ResponseWriter, r *http.Request) {
	var newHotel Hotel
	err := json.NewDecoder(r.Body).Decode(&newHotel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hotels = append(hotels, newHotel)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hotels)
}
func UpdateHotels(w http.ResponseWriter, r *http.Request) {
	var updatedHotels []Hotel
	err := json.NewDecoder(r.Body).Decode(&updatedHotels)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, updatedHotel := range updatedHotels {
		found := false
		for i, hotel := range hotels {
			if updatedHotel.Nom == hotel.Nom {
				hotels[i] = updatedHotel
				found = true
				break
			}
		}
		if !found {
			http.Error(w, fmt.Sprintf("Hotel %s not found", updatedHotel.Nom), http.StatusNotFound)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hotels)
}

func DeleteHotels(w http.ResponseWriter, r *http.Request) {
	hotels = nil

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hotels)
}

type Categorie struct {
	Classe       string  `json:"classe"`
	TarifNormal  float64 `json:"tarif_normal"`
	TarifSpecial float64 `json:"tarif_special"`
	TarifSpe     float64 `json:"tarif_spe"`
}

var categories []Categorie

func CreateCategorie(w http.ResponseWriter, r *http.Request) {
	var newCategorie Categorie
	err := json.NewDecoder(r.Body).Decode(&newCategorie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, categorie := range categories {
		if newCategorie.Classe == categorie.Classe {
			http.Error(w, fmt.Sprintf("Categorie %s already exists", newCategorie.Classe), http.StatusConflict)
			return
		}
	}

	categories = append(categories, newCategorie)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newCategorie)
}
func UpdateCategorie(w http.ResponseWriter, r *http.Request) {
	var updatedCategorie Categorie
	err := json.NewDecoder(r.Body).Decode(&updatedCategorie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, categorie := range categories {
		if updatedCategorie.Classe == categorie.Classe {
			categories[i] = updatedCategorie

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedCategorie)

			return
		}
	}

	http.Error(w, fmt.Sprintf("Categorie %s not found", updatedCategorie.Classe), http.StatusNotFound)
}
func DeleteCategorie(w http.ResponseWriter, r *http.Request) {
	classe := r.URL.Query().Get("classe")

	for i, categorie := range categories {
		if categorie.Classe == classe {
			categories = append(categories[:i], categories[i+1:]...)

			w.WriteHeader(http.StatusNoContent)

			return
		}
	}

	http.Error(w, fmt.Sprintf("Categorie %s not found", classe), http.StatusNotFound)
}

type Facture struct {
	Numero  int     `json:"numero"`
	Date    string  `json:"date"`
	Montant float64 `json:"montant"`
}

var factures []Facture

func CreateFacture(w http.ResponseWriter, r *http.Request) {
	var facture Facture

	err := json.NewDecoder(r.Body).Decode(&facture)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	factures = append(factures, facture)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(facture)
}

func UpdateFacture(w http.ResponseWriter, r *http.Request) {
	factureNumero, err := strconv.Atoi(r.URL.Query().Get("numero"))

	if err != nil {
		http.Error(w, "Invalid facture number", http.StatusBadRequest)
		return
	}

	var updatedFacture Facture

	err = json.NewDecoder(r.Body).Decode(&updatedFacture)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i := range factures {
		if factures[i].Numero == factureNumero {
			factures[i] = updatedFacture

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updatedFacture)

			return
		}
	}

	http.Error(w, fmt.Sprintf("Facture with number %d not found", factureNumero), http.StatusNotFound)
}

func DeleteFacture(w http.ResponseWriter, r *http.Request) {
	factureNumero, err := strconv.Atoi(r.URL.Query().Get("numero"))

	if err != nil {
		http.Error(w, "Invalid facture number", http.StatusBadRequest)
		return
	}

	for i := range factures {
		if factures[i].Numero == factureNumero {
			factures = append(factures[:i], factures[i+1:]...)

			w.WriteHeader(http.StatusOK)

			return
		}
	}

	http.Error(w, fmt.Sprintf("Facture with number %d not found", factureNumero), http.StatusNotFound)
}

type Service struct {
	ID                 int     `json:"id"`
	Phone              bool    `json:"phone"`
	Bar                bool    `json:"bar"`
	TarifPetitDejeuner float64 `json:"tarifPetitDejeuner"`
	TarifPhone         float64 `json:"tarifPhone"`
	TarifBar           float64 `json:"tarifBar"`
}

var services []Service

func CreateService(w http.ResponseWriter, r *http.Request) {
	var newService Service

	err := json.NewDecoder(r.Body).Decode(&newService)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	services = append(services, newService)

	w.WriteHeader(http.StatusCreated)
}

var nextID int = 1

func UpdateService(w http.ResponseWriter, r *http.Request) {
	serviceID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid service ID", http.StatusBadRequest)
		return
	}

	var updatedService Service

	err = json.NewDecoder(r.Body).Decode(&updatedService)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Vérification de l'existence du service avec l'ID donné
	serviceIndex := -1
	for i, s := range services {
		if s.ID == serviceID {
			serviceIndex = i
			break
		}
	}

	if serviceIndex == -1 {
		http.Error(w, "Service not found", http.StatusNotFound)
		return
	}

	// Mise à jour du service avec les nouvelles données
	updatedService.ID = serviceID
	services[serviceIndex] = updatedService

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Service with ID %d updated", serviceID)
}

func deleteServices(w http.ResponseWriter, r *http.Request) {
	var serviceIDs []int
	err := json.NewDecoder(r.Body).Decode(&serviceIDs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deletedServices := []Service{}
	for _, id := range serviceIDs {
		for i, service := range services {
			if service.ID == id {
				services = append(services[:i], services[i+1:]...)
				deletedServices = append(deletedServices, service)
				break
			}
		}
	}

	json.NewEncoder(w).Encode(deletedServices)
}

func main() {
	http.HandleFunc("/clients", createClient)
	http.HandleFunc("/chambreList", createChambreList)
	http.HandleFunc("/reservations", CreateReservation)
	http.HandleFunc("/reservations", UpdateReservations)
	http.HandleFunc("/reservations", DeleteReservations)
	http.HandleFunc("/hotels", CreateHotel)
	http.HandleFunc("/hotels", UpdateHotels)
	http.HandleFunc("/hotels", DeleteHotels)
	http.HandleFunc("/categories", CreateCategorie)
	http.HandleFunc("/categories", UpdateCategorie)
	http.HandleFunc("/categories", DeleteCategorie)
	http.HandleFunc("/factures", CreateFacture)
	http.HandleFunc("/factures/update", UpdateFacture)
	http.HandleFunc("/factures/delete", DeleteFacture)
	http.HandleFunc("/services", CreateService)
	http.HandleFunc("/services/update", UpdateService)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
