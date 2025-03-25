package main

import (
	"api/src/config"
	"api/src/router"

	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	config.Load()
	r := router.GenerateRouter()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	handler := corsHandler.Handler(r)

	fmt.Printf("Rodando API na porta %d\n", config.Host)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Host), handler))
}
