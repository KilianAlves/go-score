package player

import (
	"but3/go-score/services/mongodb"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type playerRepository struct {
	Collection *mongo.Collection
}

// instance unique (non exportée) du repository
var repository = playerRepository{
	Collection: mongodb.Collection("player"),
}

// accesseur à l'unique instance du repository
func Repository() playerRepository {
	return repository
}

// méthode retournant toutes les données de la collection, ou une erreur en cas de problème
func (repository *playerRepository) FindAll(lastName, country, tour string) ([]PlayerData, error) {
	var result []PlayerData
	cursor, err := repository.Collection.Find(context.Background(), bson.D{
		{Key: "lastname", Value: lastName},
		{Key: "country", Value: country},
		{Key: "tour", Value: tour},
	})
	cursor.All(context.Background(), &result)

	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (repository *playerRepository) insertOne(Data PlayerData) (PlayerData, error) {
	result, err := repository.Collection.InsertOne(context.Background(), Data)
	Data.Id = result.InsertedID.(primitive.ObjectID)
	return Data, err
}

func (repository *playerRepository) findById(Id primitive.ObjectID) (PlayerData, error) {
	var result PlayerData
	filter := bson.D{{Key: "_id", Value: Id}}
	cursor := repository.Collection.FindOne(context.Background(), filter)
	cursor.Decode(&result)
	return result, nil
}

func (repository *playerRepository) Delete(Id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.D{{Key: "_id", Value: Id}}
	result, err := repository.Collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
