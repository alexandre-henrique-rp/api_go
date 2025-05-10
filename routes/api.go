package routes

import (
	"database/sql"
	"github.com/alexandre-henrique-rp/simples_crud_go/controller"
	"github.com/gorilla/mux"
)

func ApiRoutes(router *mux.Router, db *sql.DB)  {
	//criar uma base de rota
	api := router.PathPrefix("/api").Subrouter()

	apiController := controller.NewApiControllerGet(db)

	//criar rota get
	api.HandleFunc("/stock-exchange", apiController.FindAll).Methods("GET")

	//criar rota post
	api.HandleFunc("/stock-exchange", apiController.Create).Methods("POST")

	//criar rota put
	api.HandleFunc("/stock-exchange/{id}", apiController.Update).Methods("PUT")

	//criar rota delete
	api.HandleFunc("/stock-exchange/{id}", apiController.Delete).Methods("DELETE")
}

	