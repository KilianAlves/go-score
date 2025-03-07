package main

import (
	"but3/go-score/hello"
	"but3/go-score/player"
	"but3/go-score/services/mongodb"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	// temp, verif mongoDB
	if mongodb.Collection("hello").Name() != "hello" {
		panic("Problème d'accès à la collection hello")
	}
	// création du routeur
	router := http.NewServeMux()

	// association de la route /api/hello/world (avec de la méthode GET) à la fonction HelloWorld
	router.Handle("/api/hello/", http.StripPrefix("/api/hello", hello.Router()))
	router.Handle("/api/player/", http.StripPrefix("/api/player", player.Router()))

	// configuration du serveur
	server := http.Server{
		Addr:    ":8888",
		Handler: cors.AllowAll().Handler(router),
	}

	log.Println("Serveur local démarré : http://localhost:8888")

	// démarrage du serveur
	err := server.ListenAndServe()

	log.Fatalln(err)
}
