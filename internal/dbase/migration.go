package dbase

import (
	"embed"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
)

func (db *DB) RunMigration(fs embed.FS) error {
	d, err := iofs.New(fs, ".")
	if err != nil {
		return fmt.Errorf("can't read migration from iofs: %w", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, db.GetConnectionString())
	if err != nil {
		return fmt.Errorf("can't read migration from iofs: %w", err)
	}
	defer m.Close()

	if err = m.Up(); !errors.Is(err, migrate.ErrNoChange) {
		if err != nil {
			return fmt.Errorf("can't apply migration: %w", err)
		}

		return nil
	}

	return nil
}
