package mongodb

import "go.mongodb.org/mongo-driver/mongo"

// ouvre une connection au serveur MongoDB et retourne la base de données de l'application.
// Les informations (username, password, db) seront lues dans les variables d'environnement correspondantes
func connectToDB() *mongo.Database {
	return nil
}

// singleton
var db *mongo.Database = connectToDB()

// retourne une collection identifiée par son nom
func Collection(name string) *mongo.Collection {
	return nil
}
