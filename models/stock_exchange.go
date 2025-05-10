package models


// StockExchange representa uma criptomoeda na exchange
type StockExchange struct {
	// ID único da criptomoeda
	Id int `json:"id" example:"1"`
	// Nome da criptomoeda
	Name string `json:"name" example:"Bitcoin"`
	// Símbolo da criptomoeda
	Symbol string `json:"symbol" example:"BTC"`
	// Preço atual em USD
	Price float64 `json:"price" example:"45000.50"`
	// Valor em BRL
	Valor float64 `json:"valor" example:"225000.75"`
	// Data e hora de criação do registro
	CreatedAt string `json:"created_at" example:"2025-05-10 15:58:00"`
}

const (
	TableName = "stock_exchange"
	CreateTableSQL = `CREATE TABLE IF NOT EXISTS stock_exchange (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		symbol TEXT NOT NULL,
		price REAL NOT NULL,
		valor REAL NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`
)