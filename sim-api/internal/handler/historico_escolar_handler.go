package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JoaoJorgedosAnjos/sistema-de-interesse-de-matricula-UnB/internal/domain"
	"github.com/JoaoJorgedosAnjos/sistema-de-interesse-de-matricula-UnB/internal/repository"
	"github.com/go-chi/chi/v5"
)

type HistoricoEscolarHandler struct {
	repo *repository.HistoricoEscolarRepository
}

func NewHistoricoEscolarHandler(repo *repository.HistoricoEscolarRepository) *HistoricoEscolarHandler {
	return &HistoricoEscolarHandler{repo: repo}
}

// @Summary      Lista o histórico de um aluno
// @Description  Retorna o histórico escolar completo para uma dada matrícula
// @Tags         Histórico Escolar
// @Accept       json
// @Produce      json
// @Param        matricula  path      string  true  "Matrícula do Aluno"
// @Success      200        {array}   domain.HistoricoEscolar
// @Failure      500        {object}  map[string]string
// @Router       /alunos/{matricula}/historico [get]
func (h *HistoricoEscolarHandler) GetByMatricula(w http.ResponseWriter, r *http.Request) {
	matricula := chi.URLParam(r, "matricula")
	historico, err := h.repo.FindAllByMatricula(matricula)
	if err != nil {
		http.Error(w, "Erro ao buscar histórico: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(historico)
}

// @Summary      Adiciona um item ao histórico
// @Description  Adiciona um novo registro de matéria cursada ao histórico de um aluno
// @Tags         Histórico Escolar
// @Accept       json
// @Produce      json
// @Param        historico  body      domain.HistoricoEscolar  true  "Dados do Item do Histórico"
// @Success      201        {object}  domain.HistoricoEscolar
// @Failure      400        {object}  map[string]string
// @Router       /historico [post]
func (h *HistoricoEscolarHandler) Create(w http.ResponseWriter, r *http.Request) {
	var item domain.HistoricoEscolar
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	novoItem, err := h.repo.Create(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(novoItem)
}

// @Summary      Atualiza um item do histórico
// @Description  Atualiza a nota ou status de um item do histórico escolar
// @Tags         Histórico Escolar
// @Accept       json
// @Produce      json
// @Param        id         path      int                      true  "ID do Item do Histórico"
// @Param        historico  body      domain.HistoricoEscolar  true  "Dados para Atualizar"
// @Success      200        {object}  domain.HistoricoEscolar
// @Router       /historico/{id} [put]
func (h *HistoricoEscolarHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var item domain.HistoricoEscolar
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	itemAtualizado, err := h.repo.Update(id, item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(itemAtualizado)
}

// @Summary      Apaga um item do histórico
// @Description  Apaga um item específico do histórico escolar pelo seu ID
// @Tags         Histórico Escolar
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID do Item do Histórico"
// @Success      204  {object}  nil
// @Router       /historico/{id} [delete]
func (h *HistoricoEscolarHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}