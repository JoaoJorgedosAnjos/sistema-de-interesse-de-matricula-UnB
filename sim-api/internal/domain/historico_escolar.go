package domain

import "time"

type HistoricoEscolar struct {
	IDHistoricoItem       int       `json:"id_historico_item"` 
	MatriculaAluno        string    `json:"matricula_aluno"`
	CodMateria            int       `json:"cod_materia"`
	IdTurmaOrigem         int       `json:"id_turma_origem"`
	SemestreConclusao     string    `json:"semestre_conclusao"`
	NotaFinal             float64    `json:"nota_final"`
	StatusConclusao       string    `json:"status_conclusao"`
	DataRegistroHistorico time.Time `json:"data_registro_historico"`

	NomeMateria    string `json:"nome_materia,omitempty"`
	CodDisciplina string `json:"cod_disciplina,omitempty"`
}
