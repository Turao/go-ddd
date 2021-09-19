package sql

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type migrator struct {
	db      *sql.DB
	migrate *migrate.Migrate
}

func NewMigrator(db *sql.DB, dir string) (*migrator, error) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	migrate, err := migrate.NewWithDatabaseInstance(dir, "postgres", driver)
	if err != nil {
		return nil, err
	}

	return &migrator{
		db:      db,
		migrate: migrate,
	}, nil
}

func (m *migrator) Up() error {
	err := m.migrate.Up()
	if err != nil {
		if err != migrate.ErrNoChange {
			return errors.Wrap(err, "unable to perform up migration")
		}
	}
	return nil
}

func (m *migrator) Down() error {
	err := m.migrate.Down()
	if err != nil {
		return errors.Wrap(err, "unable to perform down migration")
	}
	return nil
}
