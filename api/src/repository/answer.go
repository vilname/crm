package repository

import (
	"api/config/storage"
	"api/src/model"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AnswerRepository struct {
	db  *pgxpool.Pool
	ctx *gin.Context
}

func NewAnswerRepository(ctx *gin.Context) *AnswerRepository {
	return &AnswerRepository{
		db:  storage.GetDB(),
		ctx: ctx,
	}
}

func (repository *AnswerRepository) GetList(limit uint16, offset uint16) ([]model.Answer, error) {
	var answerList []model.Answer

	query := `select * from answers limit $1 offset $2`

	rows, err := repository.db.Query(repository.ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var answer model.Answer

		err = rows.Scan(&answer.Id, &answer.Title, &answer.Text)

		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}

		answerList = append(answerList, answer)
	}

	rows.Close()

	return answerList, nil
}

func (repository *AnswerRepository) Get(id int) (model.Answer, error) {
	var answer model.Answer

	query := `select * from answers where id=$1`
	row := repository.db.QueryRow(repository.ctx, query, id)

	err := row.Scan(&answer.Id, &answer.Title, &answer.Text)

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return answer, err
	}

	return answer, nil
}

func (repository *AnswerRepository) Count() (uint16, error) {
	var total uint16

	query := `select count(*) from answers`
	row := repository.db.QueryRow(repository.ctx, query)

	err := row.Scan(&total)

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return total, err
	}

	return total, nil
}

func (repository *AnswerRepository) Save(answer model.Answer) error {
	query := `insert into answers (title, text) values ($1, $2)`
	_, err := repository.db.Exec(repository.ctx, query, answer.Title, answer.Text)

	if err != nil {
		return err
	}

	return nil
}
