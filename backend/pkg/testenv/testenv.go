package testenv

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/golang-migrate/migrate/v4"

	"github.com/goawwer/devclash/config"
	"github.com/goawwer/devclash/internal/adapter/database"
	"github.com/goawwer/devclash/pkg/logger"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

type TestEnv struct {
	Db     *sqlx.DB
	Config *config.Config
}

func NewTestEnv(ctx context.Context, t *testing.T) *TestEnv {
	os.Setenv("ENV", string(database.EnvTest))

	cfg, err := config.New()
	require.NoError(t, err)

	err = database.Init(ctx, &cfg.Database)
	require.NoError(t, err)

	return &TestEnv{
		Db:     database.Get(),
		Config: cfg,
	}
}

func (te *TestEnv) SetupDb(ctx context.Context, t *testing.T) func(t *testing.T) {
	te.dropSchema(ctx, t)

	output, err := exec.Command("go", "list", "-f", "{{.Module.Dir}}").Output()
	if err != nil {
		logger.Error("failed to take output from exec.Command(pwd)")
	}

	m, err := migrate.New(
		fmt.Sprintf("file:///%s/migrations", strings.TrimSpace(string(output))),
		te.Config.Database.DSN(),
	)
	require.NoError(t, err)
	require.Contains(t, te.Config.Database.DSN(), "devclash_test")

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		require.NoError(t, err)
	}

	return te.teardownDb
}

func (te *TestEnv) dropSchema(ctx context.Context, t *testing.T) {
	tx, err := te.Db.BeginTxx(ctx, nil)
	require.NoError(t, err)

	// Drop everything that is not a system object
	_, err = tx.Exec(`
		DROP SCHEMA IF EXISTS public CASCADE;
		CREATE SCHEMA public;
		GRANT ALL ON SCHEMA public TO CURRENT_USER;
		GRANT ALL ON SCHEMA public TO public;
	`)
	require.NoError(t, err)
	require.NoError(t, tx.Commit())
}

func (te *TestEnv) teardownDb(t *testing.T) {
	_, err := te.Db.Exec(`
		DO $$
			DECLARE r RECORD;
		BEGIN
  			FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = current_schema()) LOOP
    			EXECUTE 'TRUNCATE TABLE ' || quote_ident(r.tablename) || ' RESTART IDENTITY CASCADE';
  			END LOOP;
		END $$;
	`)
	require.NoError(t, err)
}
