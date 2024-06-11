package mongodb

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ouvre une connection au serveur MongoDB et retourne la base de données de l'application.
// Les informations (username, password, db) seront lues dans les variables d'environnement correspondantes
func connectToDB() *mongo.Database {
	godotenv.Load()

	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)

	user := os.Getenv("username")
	pass := os.Getenv("password")
	// port := os.Getenv("port") // 27017
	dbName := os.Getenv("db")
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+user+":"+pass+"@localhost:27017"))

	return client.Database(dbName)
}

// singleton
var db *mongo.Database = connectToDB()

// retourne une collection identifiée par son nom
func Collection(name string) *mongo.Collection {
	collection := db.Collection(name)
	return collection
}
