package player

import (
	"but3/go-score/services/send"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ReadAll(res http.ResponseWriter, req *http.Request) {

	lastname := req.URL.Query().Get("lastname")
	country := req.URL.Query().Get("country")
	tour := req.URL.Query().Get("tour")

	result, _ := repository.FindAll(lastname, country, tour)

	send.Json(result, res)
}

func Create(res http.ResponseWriter, req *http.Request) {
	var playerData PlayerData
	err := json.NewDecoder(req.Body).Decode(&playerData)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	Data, _ := repository.insertOne(playerData)

	send.Json(Data, res)
}

func Read(res http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	result, _ := repository.findById(objectId)
	send.Json(result, res)
}

func Delete(res http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	result, _ := repository.Delete(objectId)
	send.Json(result, res)
}
