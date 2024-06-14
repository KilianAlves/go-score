package main

import (
	"but3/go-score/hello"
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
	router.HandleFunc("GET /api/hello", hello.ReadAll)
	router.HandleFunc("POST /api/hello", hello.Create)
	router.HandleFunc("GET /api/hello/", hello.Read)
	router.HandleFunc("DELETE /api/hello/", hello.Delete)
	router.HandleFunc("GET /api/hello/world", hello.HelloWorld)
	router.HandleFunc("GET /api/hello/square/", hello.Square)
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
