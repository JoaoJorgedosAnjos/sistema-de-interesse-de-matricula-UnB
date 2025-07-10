package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/JoaoJorgedosAnjos/sistema-de-interesse-de-matricula-UnB/internal/repository"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("chave_secreta_do_projeto_unb")

type AuthHandler struct {
	alunoRepo *repository.AlunoRepository
}

func NewAuthHandler(alunoRepo *repository.AlunoRepository) *AuthHandler {
	return &AuthHandler{alunoRepo: alunoRepo}
}

type Credentials struct {
	Matricula string `json:"matricula"`
	Senha     string `json:"senha"`
}

// @Summary      Autentica um usuário
// @Description  Recebe matrícula e senha, retorna um token JWT se forem válidas
// @Tags         Autenticação
// @Accept       json
// @Produce      json
// @Param        credentials  body      Credentials  true  "Credenciais de Login"
// @Success      200          {object}  map[string]string
// @Failure      401          {object}  map[string]string
// @Router       /login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Requisição inválida", http.StatusBadRequest)
		return
	}

	aluno, err := h.alunoRepo.FindByID(creds.Matricula)
	if err != nil {
		http.Error(w, "Matrícula ou senha inválida", http.StatusUnauthorized)
		return
	}

	if aluno.Senha != creds.Senha {
		http.Error(w, "Matrícula ou senha inválida", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour) 
	claims := &jwt.RegisteredClaims{
		Subject:   aluno.Matricula,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}