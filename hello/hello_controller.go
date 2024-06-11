package hello

import (
	"but3/go-score/services/send"
	"net/http"
	"strconv"
	"strings"
)

func ReadAll(res http.ResponseWriter, _ *http.Request) {
	result, _ := repository.FindAll()

	send.Json(result, res)
}

func HelloWorld(res http.ResponseWriter, _ *http.Request) {
	// préparation des données à envoyer
	data := HelloWorldData{
		Message: "hello",
	}

	send.Json(data, res)
}

func Square(res http.ResponseWriter, req *http.Request) {
	// Recup le numero dans l'URL
	parts := strings.Split(req.URL.Path, "/")
	num, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		http.Error(res, "Nombre invalide", http.StatusBadRequest)
		return
	}

	// Calculate the square
	square := num * num
	data := struct {
		Result int `json:"result"`
	}{
		Result: square,
	}
	send.Json(data, res)
}
