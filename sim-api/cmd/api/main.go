package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	_ "github.com/JoaoJorgedosAnjos/sistema-de-interesse-de-matricula-UnB/docs"

	"github.com/JoaoJorgedosAnjos/sistema-de-interesse-de-matricula-UnB/internal/database"
	"github.com/JoaoJorgedosAnjos/sistema-de-interesse-de-matricula-UnB/internal/handler"
	"github.com/JoaoJorgedosAnjos/sistema-de-interesse-de-matricula-UnB/internal/repository"
)

/// @title API do Sistema de Interesse de Matrícula - UnB
// @version 1.0
// @description Esta é a API para o projeto de banco de dados, permitindo a gestão de alunos, cursos e o registro de interesse em turmas.

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Digite "Bearer " seguido do seu token JWT. Exemplo: "Bearer eyJhbGciOi..."
func main() {
	databaseUrl := "postgres://postgres:postgres@localhost:5432/unb_database"

	db := database.NewDB(databaseUrl)
	defer db.Close(context.Background())

	alunoRepo := repository.NewAlunoRepository(db)
	alunoHandler := handler.NewAlunoHandler(alunoRepo)

	cursoRepo := repository.NewCursoRepository(db)
	cursoHandler := handler.NewCursoHandler(cursoRepo)

	registroRepo := repository.NewRegistroInteresseRepository(db)
	registroHandler := handler.NewRegistroInteresseHandler(registroRepo)

	historicoRepo := repository.NewHistoricoEscolarRepository(db)
	historicoHandler := handler.NewHistoricoEscolarHandler(historicoRepo)

	authHandler := handler.NewAuthHandler(alunoRepo)

	router := handler.NewRouter(alunoHandler, cursoHandler, registroHandler, historicoHandler, authHandler)

		port := ":8080"
	fmt.Printf("Servidor escutando na porta %s\n", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
