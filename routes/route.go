// routes/routes.go

package routes

import (
	"ESTIAM/dictionary"
	"ESTIAM/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRoutes configure les routes de l'API
func SetupRoutes(d *dictionary.Dictionary) *mux.Router {
	r := mux.NewRouter()

	// Ajoute le middleware de journalisation à toutes les routes
	r.Use(middleware.LoggingMiddleware)

	// Ajoute le middleware d'authentification à toutes les routes
	r.Use(middleware.AuthenticationMiddleware)

	// Route pour ajouter une entrée au dictionnaire (POST)
	r.HandleFunc("/entry", AddEntryHandler(d)).Methods("POST")

	// Route pour récupérer une définition par mot (GET)
	r.HandleFunc("/definition/{word}", GetDefinitionHandler(d)).Methods("GET")

	// Route pour supprimer une entrée par mot (DELETE)
	r.HandleFunc("/entry/{word}", RemoveEntryHandler(d)).Methods("DELETE")

	return r
}

// AddEntryHandler gère la création d'une entrée dans le dictionnaire
func AddEntryHandler(d *dictionary.Dictionary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Votre logique pour ajouter une entrée ici
	}
}

// GetDefinitionHandler gère la récupération d'une définition par mot
func GetDefinitionHandler(d *dictionary.Dictionary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Votre logique pour récupérer une définition ici
	}
}

// RemoveEntryHandler gère la suppression d'une entrée par mot
func RemoveEntryHandler(d *dictionary.Dictionary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Votre logique pour supprimer une entrée ici
	}
}
