// Package routes fornece as rotas da API
package routes

import (
	"database/sql"

	"github.com/alexandre-henrique-rp/simples_crud_go/controller"
	_ "github.com/alexandre-henrique-rp/simples_crud_go/models" // Usado apenas na documentação Swagger
	"github.com/gorilla/mux"
)

func ApiRoutes(router *mux.Router, db *sql.DB) {
	// Criar uma base de rota
	api := router.PathPrefix("/api").Subrouter()

	apiController := controller.NewApiControllerGet(db)

	// @Summary Lista todos os registros
	// @Description Lista os últimos 100 registros de criptomoedas
	// @Tags StockExchange
	// @Accept json
	// @Produce json
	// @Success 200 {array} models.StockExchange
	// @Failure 500 {object} string "Erro ao buscar registros"
	// @Router /stock-exchange [get]
	api.HandleFunc("/stock-exchange", apiController.FindAll).Methods("GET")

	// @Summary Busca um registro
	// @Description Busca um registro de criptomoeda pelo ID
	// @Tags StockExchange
	// @Accept json
	// @Produce json
	// @Param id path int true "StockExchange ID"
	// @Success 200 {object} models.StockExchange
	// @Failure 404 {object} string "Registro não encontrado"
	// @Router /stock-exchange/{id} [get]
	api.HandleFunc("/stock-exchange/{id}", apiController.FindById).Methods("GET")

	// @Summary Cria um novo registro
	// @Description Cria um novo registro de criptomoeda
	// @Tags StockExchange
	// @Accept json
	// @Produce json
	// @Param stockExchange body models.StockExchange true "Dados da criptomoeda"
	// @Success 201 {object} models.StockExchange
	// @Failure 400 {object} string "Erro nos dados enviados"
	// @Router /stock-exchange [post]
	api.HandleFunc("/stock-exchange", apiController.Create).Methods("POST")

	// @Summary Atualiza um registro
	// @Description Atualiza um registro de criptomoeda
	// @Tags StockExchange
	// @Accept json
	// @Produce json
	// @Param id path int true "StockExchange ID"
	// @Param stockExchange body models.StockExchange true "Dados atualizados"
	// @Success 200 {object} models.StockExchange
	// @Failure 404 {object} string "Registro não encontrado"
	// @Router /stock-exchange/{id} [put]
	api.HandleFunc("/stock-exchange/{id}", apiController.Update).Methods("PUT")

	// @Summary Remove um registro
	// @Description Remove um registro de criptomoeda
	// @Tags StockExchange
	// @Accept json
	// @Produce json
	// @Param id path int true "StockExchange ID"
	// @Success 204 {object} string "Registro removido com sucesso"
	// @Failure 404 {object} string "Registro não encontrado"
	// @Router /stock-exchange/{id} [delete]
	api.HandleFunc("/stock-exchange/{id}", apiController.Delete).Methods("DELETE")
}