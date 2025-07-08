package handler

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func NewRouter(alunoHandler *AlunoHandler, cursoHandler *CursoHandler, registroHandler *RegistroInteresseHandler) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	// O Recoverer normal não está nos mostrando o erro, então vamos criar o nosso.
	// router.Use(middleware.Recoverer) 

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API do Projeto UnB no ar!"))
	})
	
    // Rota do Swagger com um wrapper de depuração
	router.Get("/swagger/*", func(w http.ResponseWriter, r *http.Request) {
		// Este defer vai capturar qualquer "panic" que acontecer dentro do httpSwagger.Handler
		defer func() {
			if r := recover(); r != nil {
				// Se um pânico for capturado, escrevemos o erro diretamente na resposta HTTP
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("Ocorreu um pânico: %v\n", r)))
				w.Write(debug.Stack()) // E também o stack trace completo
			}
		}()

        // Chamamos o handler original do swagger aqui dentro
		httpSwagger.Handler(httpSwagger.URL("doc.json")).ServeHTTP(w, r)
	})


	// As outras rotas continuam iguais...
	router.Route("/alunos", func(r chi.Router) {
		r.Post("/", alunoHandler.CreateAluno)
		r.Get("/", alunoHandler.GetAllAlunos)
		r.Route("/{matricula}", func(r chi.Router) {
			r.Get("/", alunoHandler.GetAlunoByMatricula)
			r.Put("/", alunoHandler.UpdateAluno)
			r.Delete("/", alunoHandler.DeleteAluno)
		})
	})

	router.Route("/cursos", func(r chi.Router) {
		r.Get("/", cursoHandler.GetAllCursos)
		r.Post("/", cursoHandler.CreateCurso)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", cursoHandler.GetCursoByID)
			r.Put("/", cursoHandler.UpdateCurso)
			r.Delete("/", cursoHandler.DeleteCurso)
		})
	})

	router.Route("/interesses", func(r chi.Router) {
		r.Get("/", registroHandler.GetAll)
		r.Post("/", registroHandler.Create)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", registroHandler.GetByID)
			r.Put("/", registroHandler.Update)
			r.Delete("/", registroHandler.Delete)
		})
	})

	return router
}