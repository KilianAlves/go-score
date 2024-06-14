package player

import (
	"net/http"
)

// retourne un retour configuré avec les routes associées au préfixe /api/hello (mais sans le préfixe)
func Router() *http.ServeMux {
	router := http.NewServeMux()

	// association de la route /api/hello/world (avec de la méthode GET) à la fonction HelloWorld
	router.HandleFunc("POST /", Create)
	router.HandleFunc("GET /", ReadAll)
	router.HandleFunc("GET /{id}", Read)
	router.HandleFunc("DELETE /{id}", Delete)

	return router
}
