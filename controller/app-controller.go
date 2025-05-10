package controller

// CryptoStats representa os dados agregados da consulta
// Usamos nomes claros e tags JSON para facilitar o retorno para o frontend
// Isso segue Clean Code e facilita manutenção e entendimento
import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type AppController struct {
	Db *sql.DB
}

func NewAppController(db *sql.DB) *AppController {
	return &AppController{Db: db}
}


type CryptoStats struct {
	MinPrice    float64 `json:"min_price"`
	MaxPrice    float64 `json:"max_price"`
	MediaMinima float64 `json:"media_minima"`
	MediaMaxima float64 `json:"media_maxima"`
}

// FindFilter executa a consulta agregada para ETH nos últimos 21 dias
// Retorna estatísticas de preço em formato JSON
func (c *AppController) FindFilter(w http.ResponseWriter, r *http.Request) {
	// Consulta SQL utilizando CTE (WITH) para calcular estatísticas diárias
	query := `
		WITH daily_prices AS (
			SELECT
				DATE(created_at) AS data,
				MIN(CAST(REPLACE(REPLACE(price_usd, '$', ''), ',', '') AS REAL)) AS min_price,
				MAX(CAST(REPLACE(REPLACE(price_usd, '$', ''), ',', '') AS REAL)) AS max_price
			FROM
				cryptos
			WHERE
				tag = 'ETH'
				AND DATE(created_at) >= DATE('now', '-21 days')
			GROUP BY
				DATE(created_at)
		)
		SELECT
			ROUND(MIN(min_price), 2) AS min_price,
			ROUND(MAX(max_price), 2) AS max_price,
			ROUND(AVG(min_price), 2) AS media_minima,
			ROUND(AVG(max_price), 2) AS media_maxima
		FROM
			daily_prices;
	`

	// Executa a query e obtém os resultados agregados
	var stats CryptoStats
	err := c.Db.QueryRowContext(r.Context(), query).Scan(
		&stats.MinPrice,
		&stats.MaxPrice,
		&stats.MediaMinima,
		&stats.MediaMaxima,
	)
	if err != nil {
		// Tratamento de erro: retorna mensagem amigável e status HTTP 500
		http.Error(w, "Erro ao buscar estatísticas: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Define o tipo de conteúdo da resposta como JSON
	w.Header().Set("Content-Type", "application/json")
	// Retorna o resultado em JSON para o frontend
	json.NewEncoder(w).Encode(stats)
}

