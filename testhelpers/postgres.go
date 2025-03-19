package testhelper

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	timeout = time.Minute * 3
)

type PostgresDB struct {
	instance testcontainers.Container
	User     string
	Pass     string
	DBName   string
}

func NewPostgres(t *testing.T) *PostgresDB {
	t.Helper()

	var (
		env = map[string]string{
			"POSTGRES_USER":     "user",
			"POSTGRES_PASSWORD": "password",
			"POSTGRES_DB":       "postgres",
		}
		port = "5432/tcp"
	)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:14.3-alpine",
		ExposedPorts: []string{port},
		Env:          env,
		WaitingFor:   wait.ForListeningPort("5432/tcp").WithStartupTimeout(timeout),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	require.NoError(t, err)

	return &PostgresDB{
		User:     env["POSTGRES_USER"],
		Pass:     env["POSTGRES_PASSWORD"],
		DBName:   env["POSTGRES_DB"],
		instance: container,
	}
}

func (r *PostgresDB) Host(t *testing.T) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	p, err := r.instance.Host(ctx)
	require.NoError(t, err)

	return p
}

func (r *PostgresDB) Port(t *testing.T) int {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	p, err := r.instance.MappedPort(ctx, "5432")
	require.NoError(t, err)

	return p.Int()
}

func (r *PostgresDB) Close(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	require.NoError(t, r.instance.Terminate(ctx))
}

func (r *PostgresDB) DatabaseName() string {
	return r.DBName
}
