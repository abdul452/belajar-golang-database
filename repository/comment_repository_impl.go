package repository

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/abdul452/belajar-golang-database/entity"
	_ "github.com/go-sql-driver/mysql"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comments) (entity.Comments, error) {
	script := `INSERT INTO comment(email, comment) VALUES(?, ?)`
	result, err := repository.DB.ExecContext(ctx, script, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}

	comment.Id = int(id)
	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int) (entity.Comments, error) {
	script := `SELECT id, email, comment FROM comment WHERE id = ? LIMIT 1`
	rows, err := repository.DB.QueryContext(ctx, script, id)
	comment := entity.Comments{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()

	if rows.Next() {
		// ada
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		// tidak ada
		return comment, errors.New("Id " + strconv.Itoa(id) + " Not Found")
	}
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comments, error) {
	script := `SELECT id, email, comment FROM comment`
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []entity.Comments
	for rows.Next() {
		var comment entity.Comments
		err := rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
