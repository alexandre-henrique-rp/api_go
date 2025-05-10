package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/alexandre-henrique-rp/simples_crud_go/models"
	"github.com/gorilla/mux"
)

type ApiController struct {
	Db *sql.DB
}

func NewApiControllerGet(db *sql.DB) *ApiController {
	return &ApiController{Db: db}
}

// FindAll lista todos os registros de StockExchange
// @Summary Lista todos os registros
// @Description Lista os últimos 100 registros de criptomoedas
// @Tags StockExchange
// @Accept json
// @Produce json
// @Success 200 {array} models.StockExchange
// @Failure 500 {object} string "Erro ao buscar registros"
// @Router /api/stock-exchange [get]
func (apiController *ApiController) FindAll(writer http.ResponseWriter, request *http.Request) {
	var stockExchanges []models.StockExchange

	rows, err := apiController.Db.Query("SELECT * FROM stock_exchange ORDER BY id DESC LIMIT 100")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	for rows.Next() {
		var stockExchange models.StockExchange
		err := rows.Scan(&stockExchange.Id, &stockExchange.Name, &stockExchange.Symbol, &stockExchange.Price, &stockExchange.Valor, &stockExchange.CreatedAt)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		stockExchanges = append(stockExchanges, stockExchange)
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(stockExchanges)
	
}

// FindById busca um registro de StockExchange pelo ID
// @Summary Busca um registro
// @Description Busca um registro de criptomoeda pelo ID
// @Tags StockExchange
// @Accept json
// @Produce json
// @Param id path int true "StockExchange ID"
// @Success 200 {object} models.StockExchange
// @Failure 404 {object} string "Registro não encontrado"
// @Router /api/stock-exchange/{id} [get]
func (apiController *ApiController) FindById(writer http.ResponseWriter, request *http.Request) {
	var stockExchange models.StockExchange

	id := mux.Vars(request)["id"]

	row, err := apiController.Db.Query("SELECT * FROM stock_exchange WHERE id = ?", id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	for row.Next() {
		err := row.Scan(&stockExchange.Id, &stockExchange.Name, &stockExchange.Symbol, &stockExchange.Price, &stockExchange.Valor, &stockExchange.CreatedAt)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(stockExchange)
}

// Create cria um novo registro de StockExchange
// @Summary Cria um novo registro
// @Description Cria um novo registro de criptomoeda
// @Tags StockExchange
// @Accept json
// @Produce json
// @Param stockExchange body models.StockExchange true "Dados da criptomoeda"
// @Success 201 {object} models.StockExchange
// @Failure 400 {object} string "Erro nos dados enviados"
// @Router /api/stock-exchange [post]
func (apiController *ApiController) Create(writer http.ResponseWriter, request *http.Request) {
	var stockExchange models.StockExchange

	json.NewDecoder(request.Body).Decode(&stockExchange)

	rowCreate, err := apiController.Db.Exec("INSERT INTO stock_exchange (name, symbol, price, valor) VALUES (?, ?, ?, ?)", stockExchange.Name, stockExchange.Symbol, stockExchange.Price, stockExchange.Valor)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	
	data, err := rowCreate.RowsAffected()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	
	id := data

	row, errGet := apiController.Db.Query("SELECT * FROM stock_exchange WHERE id = ?", id)
	if errGet != nil {
		http.Error(writer, errGet.Error(), http.StatusInternalServerError)
		return
	}
	for row.Next() {
		err := row.Scan(&stockExchange.Id, &stockExchange.Name, &stockExchange.Symbol, &stockExchange.Price, &stockExchange.Valor, &stockExchange.CreatedAt)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(stockExchange)
}


// Update atualiza um registro de StockExchange pelo ID
// @Summary Atualiza um registro
// @Description Atualiza um registro de criptomoeda
// @Tags StockExchange
// @Accept json
// @Produce json
// @Param id path int true "StockExchange ID"
// @Param stockExchange body models.StockExchange true "Dados atualizados"
// @Success 200 {object} models.StockExchange
// @Failure 404 {object} string "Registro não encontrado"
// @Router /api/stock-exchange/{id} [put]
func (apiController *ApiController) Update(writer http.ResponseWriter, request *http.Request) {
	var stockExchange models.StockExchange

	id := mux.Vars(request)["id"]

	json.NewDecoder(request.Body).Decode(&stockExchange)

	_, errUpdate := apiController.Db.Exec("UPDATE stock_exchange SET name = ?, symbol = ?, price = ?, valor = ? WHERE id = ?", stockExchange.Name, stockExchange.Symbol, stockExchange.Price, stockExchange.Valor, id)
	if errUpdate != nil {
		http.Error(writer, errUpdate.Error(), http.StatusInternalServerError)
		return
	}
	
	row, errGet := apiController.Db.Query("SELECT * FROM stock_exchange WHERE id = ?", id)
	if errGet != nil {
		http.Error(writer, errGet.Error(), http.StatusInternalServerError)
		return
	}
	for row.Next() {
		err := row.Scan(&stockExchange.Id, &stockExchange.Name, &stockExchange.Symbol, &stockExchange.Price, &stockExchange.Valor, &stockExchange.CreatedAt)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(stockExchange)
}

// Delete remove um registro de StockExchange pelo ID
// @Summary Remove um registro
// @Description Remove um registro de criptomoeda
// @Tags StockExchange
// @Accept json
// @Produce json
// @Param id path int true "StockExchange ID"
// @Success 200 {object} string "Registro deletado com sucesso"
// @Failure 404 {object} string "Registro não encontrado"
// @Router /api/stock-exchange/{id} [delete]
func (apiController *ApiController) Delete(writer http.ResponseWriter, request *http.Request) {

	id := mux.Vars(request)["id"]

	_, errDelete := apiController.Db.Exec("DELETE FROM stock_exchange WHERE id = ?", id)
	if errDelete != nil {
		http.Error(writer, errDelete.Error(), http.StatusInternalServerError)
		return
	}

	//retornar um texto
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Registro deletado com sucesso"))
}
	