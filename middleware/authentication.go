// middleware/authentication.go

package middleware

import (
	"net/http"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Récupère le jeton d'authentification de l'en-tête
		token := r.Header.Get("Authorization")

		// Logique d'authentification (personnalisez selon vos besoins)
		if isValidToken(token) {
			// Le jeton est valide, autorisez l'accès
			next.ServeHTTP(w, r)
		} else {
			// Le jeton n'est pas valide, renvoie une réponse d'erreur non autorisée
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})
}

// isValidToken est une fonction de logique d'authentification (à personnaliser)
func isValidToken(token string) bool {
	// Logique d'authentification : vérifiez si le jeton est valide
	// Vous pouvez implémenter une vérification plus complexe ici
	return token == "votre_jeton_secret"
}
