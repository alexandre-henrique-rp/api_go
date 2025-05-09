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

func (apiController *ApiController) FindAll(writer http.ResponseWriter, request *http.Request) {
	var stockExchanges []models.StockExchange

	rows, err := apiController.Db.Query("SELECT * FROM stock_exchange")
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
	