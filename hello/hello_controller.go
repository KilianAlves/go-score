package hello

import (
	"but3/go-score/services/send"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ReadAll(res http.ResponseWriter, _ *http.Request) {
	result, _ := repository.FindAll()

	send.Json(result, res)
}

func Create(res http.ResponseWriter, req *http.Request) {
	var helloData HelloData
	err := json.NewDecoder(req.Body).Decode(&helloData)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	repository.insertOne(helloData)
	send.Json(helloData, res)
}

func Read(res http.ResponseWriter, req *http.Request) {
	parts := strings.Split(req.URL.Path, "/")
	id := parts[len(parts)-1]

	objectId, _ := primitive.ObjectIDFromHex(id)
	repository.findById(objectId)
	// send.Json(helloData, res)
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
