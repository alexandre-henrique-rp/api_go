// @title API de Criptomoedas
// @version 1.0
// @description API para gerenciamento de dados de criptomoedas
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3030
// @BasePath /
// @schemes http

// @tag.name StockExchange
// @tag.description Operações com criptomoedas

// @tag.name Pesquisa
// @tag.description Consultas e estatísticas

package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/alexandre-henrique-rp/simples_crud_go/controller" // Import para documentação Swagger
	"github.com/alexandre-henrique-rp/simples_crud_go/config"
	"github.com/alexandre-henrique-rp/simples_crud_go/models"
	_ "github.com/alexandre-henrique-rp/simples_crud_go/docs" // importação dos docs swagger
	"github.com/alexandre-henrique-rp/simples_crud_go/routes"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	connection := config.SetupDB()
	_, err := connection.Exec(models.CreateTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	routes.ApiRoutes(router, connection)
	routes.AppRoutes(router, connection)

	// Obter a porta do ambiente
	port := os.Getenv("PORT")
	if port == "" {
		port = "3030" // Porta padrão se não estiver definida
	}

	// Serve a documentação Swagger na rota /swagger/
	// Use o IP do servidor para evitar problemas de CORS
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // Caminho relativo para evitar problemas de CORS
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("list"),
		httpSwagger.DomID("swagger-ui"),
	))

	defer connection.Close()

	// Middleware simples para liberar CORS
	enableCORS := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

	// Inicializa o servidor
	log.Printf("Server running on http://0.0.0.0:%s", port)
	log.Printf("Swagger running on http://0.0.0.0:%s/swagger/index.html", port)
	if err := http.ListenAndServe(":"+port, enableCORS(router)); err != nil {
		log.Fatal(err)
	}
}