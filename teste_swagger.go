// ARQUIVO: teste_swagger.go (VERSÃO CORRIGIDA)
package main

import (
	"log"
	"net/http"

	_ "github.com/JoaoJorgedosAnjos/sistema-de-interesse-de-matricula-UnB/docs"

	"github.com/go-chi/chi/v5" // <-- CORRIGIDO AQUI
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title API de Teste Mínimo
// @version 1.0
// @description Teste para depurar o Swagger.
// @host localhost:8080
// @BasePath /
func main() {
	router := chi.NewRouter()

	// Uma única rota de teste
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Servidor de teste no ar!"))
	})
	
	// A rota do Swagger
	router.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("doc.json")))

	log.Println("Servidor de teste escutando na porta :8080")
	log.Println("Acesse http://localhost:8080/swagger/index.html")

	// Inicia o servidor
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Erro ao iniciar o servidor de teste: %v", err)
	}
}