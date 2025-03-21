package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	r := router.GenerateRouter()

	fmt.Printf("Rodando API na porta %d", config.Host)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Host), r))
}
