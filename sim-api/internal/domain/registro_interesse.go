package domain

import "time"

type RegistroInteresse struct {
	IDRegistroInteresse   int       `json:"id_registro_interesse"`
	MatriculaAluno        string    `json:"matricula_aluno"`
	IDTurma               int       `json:"id_turma"`
	DataRegistroInteresse time.Time `json:"data_registro_interesse"`
	StatusInteresse       string    `json:"status_interesse"`
	PrioridadeInteresse   int       `json:"prioridade_interesse"`

	NomeMateria     string `json:"nome_materia,omitempty"`
	CodDisciplina   string `json:"cod_disciplina,omitempty"`
	SemestreOferta  string `json:"semestre_oferta,omitempty"`
	Horario         string `json:"horario,omitempty"`
}