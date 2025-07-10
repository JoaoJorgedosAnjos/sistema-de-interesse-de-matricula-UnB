package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func NewRouter(alunoHandler *AlunoHandler, cursoHandler *CursoHandler, registroHandler *RegistroInteresseHandler, historicoHandler *HistoricoEscolarHandler, authHandler *AuthHandler) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")

			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	})

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API do Projeto UnB no ar!"))
	})
	router.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("doc.json")))

	router.Route("/alunos", func(r chi.Router) {
		r.Post("/", alunoHandler.CreateAluno)
		r.Get("/", alunoHandler.GetAllAlunos)
		r.Route("/{matricula}", func(r chi.Router) {
			r.Get("/", alunoHandler.GetAlunoByMatricula) // <-- CORRIGIDO AQUI
			r.Put("/", alunoHandler.UpdateAluno)
			r.Delete("/", alunoHandler.DeleteAluno)
			r.Post("/foto", alunoHandler.UploadFotoAluno)
			r.Get("/historico", historicoHandler.GetByMatricula)
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

	router.Route("/historico", func(r chi.Router) {
		r.Post("/", historicoHandler.Create)
		r.Route("/{id}", func(r chi.Router) {
			r.Put("/", historicoHandler.Update)
			r.Delete("/", historicoHandler.Delete)
		})
	})

	router.Group(func(r chi.Router) {
		r.Use(AuthMiddleware)

		r.Get("/meus-interesses", registroHandler.GetMeusInteresses)
	})

	router.Post("/login", authHandler.Login)

	return router
}
