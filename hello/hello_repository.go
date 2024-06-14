package hello

import (
	"but3/go-score/services/mongodb"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// structure (non exportée) représentant un repository
type helloRepository struct {
	Collection *mongo.Collection
}

// instance unique (non exportée) du repository
var repository = helloRepository{
	Collection: mongodb.Collection("hello"),
}

// accesseur à l'unique instance du repository
func Repository() helloRepository {
	return repository
}

// méthode retournant toutes les données de la collection, ou une erreur en cas de problème
func (repository *helloRepository) FindAll() ([]HelloData, error) {
	var result []HelloData
	cursor, err := repository.Collection.Find(context.Background(), bson.D{})
	cursor.All(context.Background(), &result)

	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (repository *helloRepository) insertOne(message HelloData) (*mongo.InsertOneResult, error) {
	return repository.Collection.InsertOne(context.Background(), message)

}

func (repository *helloRepository) findById(Id primitive.ObjectID) (HelloData, error) {
	var result HelloData
	filter := bson.D{{Key: "_id", Value: Id}}
	cursor := repository.Collection.FindOne(context.Background(), filter)
	cursor.Decode(&result)
	return result, nil
}

func (repository *helloRepository) Delete(Id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.D{{Key: "_id", Value: Id}}
	result, err := repository.Collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
