package main

import (
	"log"
	"net/http"

	"github.com/alexandre-henrique-rp/simples_crud_go/config"
	"github.com/alexandre-henrique-rp/simples_crud_go/models"
	"github.com/alexandre-henrique-rp/simples_crud_go/routes"
	"github.com/gorilla/mux"
)

func main() {
	connection := config.SetupDB()
	_, err := connection.Exec(models.CreateTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	routes.SetupRoutes(router, connection)

	defer connection.Close()
	// inicializa o servidor
	servidor := http.ListenAndServe(":3030", router)
	if servidor != nil {
		log.Fatal(servidor)
	}
}
