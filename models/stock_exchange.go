package models


type StockExchange struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Price  float64 `json:"price"`
	Valor  float64 `json:"valor"`
	// Usar string para facilitar o Scan do SQLite em testes E2E
	// Em produção, pode-se usar time.Time e tratar a conversão
	CreatedAt string `json:"created_at"`
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