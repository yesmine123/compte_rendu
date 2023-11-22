// middleware/logging.go

package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// LoggingMiddleware enregistre chaque requête dans un fichier journal
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Début de la requête
		start := time.Now()

		// Passe la requête au gestionnaire suivant
		next.ServeHTTP(w, r)

		// Fin de la requête
		end := time.Now()

		// Calcul de la durée totale
		duration := end.Sub(start)

		// Format du message de journalisation
		logMessage := fmt.Sprintf("%s - %s %s %s - %v", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.RequestURI, r.RemoteAddr, duration)

		// Ajoutez le message de journalisation à un fichier (ou imprimez-le sur la console, selon vos besoins)
		appendToLogFile(logMessage)
	})
}

func appendToLogFile(message string) {
	// Ouvre (ou crée) un fichier journal en mode ajout
	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier journal :", err)
		return
	}
	defer file.Close()

	// Ajoute le message au fichier journal
	if _, err := file.WriteString(message + "\n"); err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier journal :", err)
		return
	}
}
