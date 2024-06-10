package main

import (
	"but3/go-score/hello"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	// création du routeur
	router := http.NewServeMux()

	// association de la route /api/hello/world (avec de la méthode GET) à la fonction HelloWorld
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
