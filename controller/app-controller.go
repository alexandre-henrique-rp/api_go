package controller

// CryptoStats representa os dados agregados da consulta
// Usamos nomes claros e tags JSON para facilitar o retorno para o frontend
// Isso segue Clean Code e facilita manutenção e entendimento
import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type AppController struct {
	Db *sql.DB
}

func NewAppController(db *sql.DB) *AppController {
	return &AppController{Db: db}
}


// CryptoStats representa as estatísticas de preço de uma criptomoeda
type CryptoStats struct {
	// Preço mínimo registrado no período
	MinPrice float64 `json:"min_price" example:"42000.50"`
	// Preço máximo registrado no período
	MaxPrice float64 `json:"max_price" example:"45000.75"`
	// Média dos preços mínimos diários
	MediaMinima float64 `json:"media_minima" example:"43000.25"`
	// Média dos preços máximos diários
	MediaMaxima float64 `json:"media_maxima" example:"44500.80"`
}

// FindFilter busca estatísticas de criptomoedas pelo símbolo
// @Summary Pesquisa estatísticas de criptomoedas
// @Description Retorna preços mínimos, máximos e médias dos últimos 21 dias
// @Tags Pesquisa
// @Accept json
// @Produce json
// @Param symbol path string true "Símbolo da criptomoeda (ex: BTC, ETH)"
// @Success 200 {object} CryptoStats
// @Failure 400 {object} string "Símbolo inválido"
// @Failure 500 {object} string "Erro ao buscar estatísticas"
// @Router /app/pesquisa/{symbol} [get]
func (c *AppController) FindFilter(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]
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
				tag = ?
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
	err := c.Db.QueryRowContext(r.Context(), query, symbol).Scan(
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
