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
		return fmt.Errorf("migration")
	}

	connectionString := db.GetConnectionString()

	m, err := migrate.NewWithSourceInstance("iofs", d, connectionString)
	if err != nil {
		return fmt.Errorf("migration read from iofs %w", err)
	}
	defer m.Close()

	if err = m.Up(); !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migration up %w", err)
	}

	return nil
}
