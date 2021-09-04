package sql

import (
	"context"
	"database/sql"

	"github.com/turao/go-ddd/projects/domain/project"
)

type repo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) (*repo, error) {
	return &repo{
		db: db,
	}, nil
}

func (r *repo) Create(ctx context.Context, p project.Project) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(
		ctx,
		"INSERT INTO projects VALUES ($1, $2)",
		p.ID,
		p.Title,
	)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *repo) Update(ctx context.Context, p project.Project) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(
		ctx,
		"UPDATE projects SET title = $1 WHERE id = $2",
		p.Title,
		p.ID,
	)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id project.ProjectID) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(
		ctx,
		"DELETE FROM projects WHERE id = $1",
		id,
	)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *repo) FindProjectByID(ctx context.Context, id project.ProjectID) (*project.Project, error) {
	var p project.Project
	err := r.db.QueryRowContext(
		ctx,
		"SELECT * FROM projects WHERE id = $1",
		id,
	).Scan(&p.ID, &p.Title)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
