package routes

import (
	"database/sql"

	"github.com/alexandre-henrique-rp/simples_crud_go/controller"
	"github.com/gorilla/mux"
)

func AppRoutes(router *mux.Router, db *sql.DB)  {
	//criar uma base de rota
	app := router.PathPrefix("/app").Subrouter()

	appController := controller.NewAppController(db)

	//criar rota get
	app.HandleFunc("/pesquisa", appController.FindFilter).Methods("GET")
}