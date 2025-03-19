package migration

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Migrate(migrationFileDir, databaseURL string) error {
	migration, err := migrate.New(fmt.Sprintf("file://%s", migrationFileDir), databaseURL)
	if err != nil {
		return err
	}

	err = migration.Up()

	if errors.Is(err, migrate.ErrNoChange) {
		return nil
	}

	return err
}
