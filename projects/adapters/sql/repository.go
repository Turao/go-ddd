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

func (r *repo) Save(ctx context.Context, p project.Project) error {
	found, err := r.FindProjectByID(ctx, p.ID)
	if err != nil {
		return err
	}

	if found == nil {
		return r.insert(ctx, p)
	}

	return r.update(ctx, p)
}

func (r *repo) insert(ctx context.Context, p project.Project) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(
		ctx,
		"INSERT INTO projects VALUES ($1, $2, $3)",
		p.ID,
		p.Title,
		p.Active,
	)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *repo) update(ctx context.Context, p project.Project) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(
		ctx,
		"UPDATE projects SET title = $1, active = $2 WHERE id = $3",
		p.Title,
		p.Active,
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

func (r *repo) FindProjectByID(ctx context.Context, id project.ProjectID) (*project.Project, error) {
	var p project.Project
	err := r.db.QueryRowContext(
		ctx,
		"SELECT * FROM projects WHERE id = $1",
		id,
	).Scan(&p.ID, &p.Title, &p.Active)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &p, nil
}
