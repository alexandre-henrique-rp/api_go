package routes

import (
	"database/sql"

	"github.com/alexandre-henrique-rp/simples_crud_go/controller"
	"github.com/gorilla/mux"
)

// @title Rotas do App
// @version 1.0
// @description Rotas para consultas e estatísticas
func AppRoutes(router *mux.Router, db *sql.DB) {
	// Criar uma base de rota
	app := router.PathPrefix("/app").Subrouter()

	appController := controller.NewAppController(db)

	// @Summary Pesquisa estatísticas de criptomoedas
	// @Description Retorna preços mínimos, máximos e médias dos últimos 21 dias
	// @Tags Pesquisa
	// @Accept json
	// @Produce json
	// @Param symbol path string true "Símbolo da criptomoeda (ex: BTC, ETH)"
	// @Success 200 {object} controller.CryptoStats
	// @Failure 400 {object} string "Símbolo inválido"
	// @Failure 500 {object} string "Erro ao buscar estatísticas"
	// @Router /app/pesquisa/{symbol} [get]
	app.HandleFunc("/pesquisa/{symbol}", appController.FindFilter).Methods("GET")
}