package testhelper

import (
	"testing"

	"github.com/andrersp/favorites/internal/adapters/resources/gorm"
	"github.com/andrersp/favorites/internal/adapters/resources/migration"
)

func GetGormInstance(t *testing.T, pg *PostgresDB) (*gorm.GormInstance, error) {
	gorm := gorm.Connect(&gorm.PgOptions{
		Host:     pg.Host(t),
		User:     pg.User,
		Password: pg.Pass,
		DBName:   pg.DBName,
		Port:     pg.Port(t),
	}, gorm.PgConfig{})

	err := migration.Migrate("../../../../db/migrations", gorm.URL())

	return gorm, err
}
