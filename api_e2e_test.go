package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/gorilla/mux"
	"github.com/alexandre-henrique-rp/simples_crud_go/routes"
)

// Função utilitária para criar um banco de dados em memória para testes
func setupTestDB() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("Erro ao criar banco de dados em memória: %v", err)
	}
	// Criação da tabela conforme esperado pelo controller
	tabela := `CREATE TABLE stock_exchange (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		symbol TEXT,
		price REAL,
		valor REAL,
		created_at TEXT DEFAULT (datetime('now'))
	)`
	_, err = db.Exec(tabela)
	if err != nil {
		log.Fatalf("Erro ao criar tabela: %v", err)
	}
	return db
}

// Função utilitária para criar o servidor de testes
func setupTestServer() (*httptest.Server, *sql.DB) {
	db := setupTestDB()
	router := mux.NewRouter()
	routes.SetupRoutes(router, db)
	server := httptest.NewServer(router)
	return server, db
}

// Estrutura que representa o modelo esperado pela API
// Os nomes e tipos devem estar alinhados com o controller e a tabela

type StockExchange struct {
	Id        int     `json:"id,omitempty"`
	Name      string  `json:"name"`
	Symbol    string  `json:"symbol"`
	Price     float64 `json:"price"`
	Valor     float64 `json:"valor"`
	CreatedAt string  `json:"created_at,omitempty"`
}

// Teste E2E para rotas da API
func TestAPIStockExchangeE2E(t *testing.T) {
	server, db := setupTestServer()
	defer server.Close()
	defer db.Close()

	// 1. Testa POST /api/stock-exchange
	// Agora enviando todos os campos obrigatórios
	newStock := StockExchange{
		Name:   "B3",
		Symbol: "B3SA3",
		Price:  100.5,
		Valor:  100.5,
	}
	jsonBody, _ := json.Marshal(newStock)
	resp, err := http.Post(server.URL+"/api/stock-exchange", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Erro ao fazer POST: %v", err)
	}
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		var buf bytes.Buffer
		buf.ReadFrom(resp.Body)
		t.Errorf("Status esperado 201 ou 200, obtido: %d. Corpo da resposta: %s", resp.StatusCode, buf.String())
	}
	resp.Body.Close()

	// 2. Testa GET /api/stock-exchange
	resp, err = http.Get(server.URL + "/api/stock-exchange")
	if err != nil {
		t.Fatalf("Erro ao fazer GET: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status esperado 200, obtido: %d", resp.StatusCode)
	}
	var stocks []StockExchange
	json.NewDecoder(resp.Body).Decode(&stocks)
	resp.Body.Close()
	if len(stocks) == 0 {
		t.Errorf("Esperava pelo menos um registro na resposta do GET")
	}

	// 3. Testa PUT /api/stock-exchange/{id}
	if len(stocks) > 0 {
		// Atualiza apenas os campos necessários
		update := StockExchange{Name: "B3 Atualizada", Symbol: "B3SA3", Price: 200.0, Valor: 200.0}
		jsonBody, _ = json.Marshal(update)
		url := server.URL + "/api/stock-exchange/" + strconv.Itoa(stocks[0].Id)
		req, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err = client.Do(req)
		if err != nil {
			t.Fatalf("Erro ao fazer PUT: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Status esperado 200 no PUT, obtido: %d", resp.StatusCode)
		}
		resp.Body.Close()
	}

	// 4. Testa DELETE /api/stock-exchange/{id}
	if len(stocks) > 0 {
		url := server.URL + "/api/stock-exchange/" + strconv.Itoa(stocks[0].Id)
		req, _ := http.NewRequest(http.MethodDelete, url, nil)
		client := &http.Client{}
		resp, err = client.Do(req)
		if err != nil {
			t.Fatalf("Erro ao fazer DELETE: %v", err)
		}
		if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
			t.Errorf("Status esperado 200 ou 204 no DELETE, obtido: %d", resp.StatusCode)
		}
		resp.Body.Close()
	}
}

// Dica: Para rodar os testes, execute: go test -v
