package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/goawwer/devclash/internal/adapter/database"
	"github.com/goawwer/devclash/internal/adapter/database/repository"
	accountmodel "github.com/goawwer/devclash/internal/domain/account_model"
	usermodel "github.com/goawwer/devclash/internal/domain/user_model"
	"github.com/goawwer/devclash/pkg/testenv"
	"github.com/goawwer/devclash/utils"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestUserRepository(t *testing.T) {
	ctx := context.TODO()

	env := testenv.NewTestEnv(ctx, t)

	cleanup := env.SetupDb(ctx, t)
	t.Cleanup(func() {
		cleanup(t)
	})

	r := repository.NewRepository(database.Get())

	hashPass, err := utils.CreateHashPassword("password")
	require.NoError(t, err)

	a := &accountmodel.Account{
		ID:             uuid.New(),
		Email:          "test@test.com",
		Role:           "user",
		HashedPassword: hashPass,
		CreatedAt:      time.Now(),
	}

	u := &usermodel.User{
		AccountID: a.ID,
		Username:  "testUser",
	}

	create(t, ctx, r, u, a)
}

func create(t *testing.T, ctx context.Context, r repository.Repository, u *usermodel.User, a *accountmodel.Account) {
	require.NoError(t, r.CreateUser(ctx, a, u))
}
