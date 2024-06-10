package send

import (
	"encoding/json"
	"net/http"
)

func Json(data any, res http.ResponseWriter) {
	// génération des données JSON
	jsonData, err := json.Marshal(data)

	// arrêt en cas d'erreur
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	// définition du type de contenu
	res.Header().Set("Content-Type", "application/json")

	// envoi des données
	res.Write(jsonData)
}
