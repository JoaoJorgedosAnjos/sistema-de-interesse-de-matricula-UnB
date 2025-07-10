package repository

import (
	"context"
	"fmt"

	"github.com/JoaoJorgedosAnjos/sistema-de-interesse-de-matricula-UnB/internal/domain"
	"github.com/jackc/pgx/v5"
)

type HistoricoEscolarRepository struct{
	db *pgx.Conn
}

func NewHistoricoEscolarRepository(db *pgx.Conn)*HistoricoEscolarRepository{
	return &HistoricoEscolarRepository{db:db}
}

func (r *HistoricoEscolarRepository) FindAllByMatricula(matricula string) ([]domain.HistoricoEscolar, error) {
	query := `
		SELECT
			h.id_historico_item, h.matricula_aluno, h.cod_materia, h.id_turma_origem,
			h.semestre_conclusao, h.nota_final, h.status_conclusao,
			h.data_registro_historico, m.nome_materia, m.cod_disciplina
		FROM
			HISTORICO_ESCOLAR h
		JOIN
			MATERIAS m ON h.cod_materia = m.cod_mat
		WHERE
			h.matricula_aluno = $1
		ORDER BY
			h.semestre_conclusao
	`

	rows, err := r.db.Query(context.Background(), query, matricula)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var historico []domain.HistoricoEscolar
	for rows.Next() {
		var item domain.HistoricoEscolar
		
		if err := rows.Scan(
			&item.IDHistoricoItem, &item.MatriculaAluno, &item.CodMateria, &item.IdTurmaOrigem,
			&item.SemestreConclusao, &item.NotaFinal, &item.StatusConclusao,
			&item.DataRegistroHistorico, &item.NomeMateria, &item.CodDisciplina,
		); err != nil {
			return nil, err
		}
		
		historico = append(historico, item)
	}
	
	return historico, nil
}

func (r *HistoricoEscolarRepository) Create(item domain.HistoricoEscolar) (domain.HistoricoEscolar, error) {
	query := `
		INSERT INTO HISTORICO_ESCOLAR (matricula_aluno, cod_materia, id_turma_origem, semestre_conclusao, nota_final, status_conclusao)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id_historico_item, matricula_aluno, cod_materia, id_turma_origem, semestre_conclusao, nota_final, status_conclusao, data_registro_historico
	`
	var novoItem domain.HistoricoEscolar
	err := r.db.QueryRow(context.Background(), query,
		item.MatriculaAluno,
		item.CodMateria,
		item.IdTurmaOrigem,
		item.SemestreConclusao,
		item.NotaFinal,
		item.StatusConclusao,
	).Scan(
		&novoItem.IDHistoricoItem,
		&novoItem.MatriculaAluno, &novoItem.CodMateria, &novoItem.IdTurmaOrigem,
		&novoItem.SemestreConclusao, &novoItem.NotaFinal, &novoItem.StatusConclusao,
		&novoItem.DataRegistroHistorico,
	)

	if err != nil {
		return domain.HistoricoEscolar{}, err
	}
	return novoItem, nil
}

func (r *HistoricoEscolarRepository) Update(id int, item domain.HistoricoEscolar) (domain.HistoricoEscolar, error) {
	query := `
		UPDATE HISTORICO_ESCOLAR
		SET nota_final = $1, status_conclusao = $2
		WHERE id_historico_item = $3
		RETURNING id_historico_item, matricula_aluno, cod_materia, id_turma_origem, semestre_conclusao, nota_final, status_conclusao, data_registro_historico
	`
	var itemAtualizado domain.HistoricoEscolar
	err := r.db.QueryRow(context.Background(), query,
		item.NotaFinal,
		item.StatusConclusao,
		id,
	).Scan(
		&itemAtualizado.IDHistoricoItem,
		&itemAtualizado.MatriculaAluno, &itemAtualizado.CodMateria, &itemAtualizado.IdTurmaOrigem,
		&itemAtualizado.SemestreConclusao, &itemAtualizado.NotaFinal, &itemAtualizado.StatusConclusao,
		&itemAtualizado.DataRegistroHistorico,
	)

	if err != nil {
		return domain.HistoricoEscolar{}, err
	}
	return itemAtualizado, nil
}

func (r *HistoricoEscolarRepository) Delete(id int) error {
	query := "DELETE FROM HISTORICO_ESCOLAR WHERE id_historico_item = $1"

	res, err := r.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("nenhum item no histórico encontrado com o ID %d", id)
	}

	return nil
}
