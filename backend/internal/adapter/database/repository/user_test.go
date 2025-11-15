package repository_test

import (
	"context"
	"testing"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/goawwer/devclash/internal/adapter/database"
	"github.com/goawwer/devclash/internal/adapter/database/repository"
	"github.com/goawwer/devclash/internal/domain/usermodel"
	"github.com/goawwer/devclash/pkg/testenv"
	"github.com/goawwer/devclash/utils"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
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

	u := &usermodel.User{
		Email:          "test@test.com",
		Username:       "testUser",
		HashedPassword: hashPass,
		CreatedAt:      time.Now(),
	}

	create(t, ctx, r, u)
	getByEmail(t, ctx, r, u.Email)
}

func create(t *testing.T, ctx context.Context, r repository.Repository, u *usermodel.User) {
	require.NoError(t, r.Create(ctx, u))
}

func getByEmail(t *testing.T, ctx context.Context, r repository.Repository, email string) {
	user, err := r.GetUserByEmail(ctx, email)
	require.NoError(t, err)

	require.Equal(t, "test@test.com", user.Email)

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte("password"))
	require.NoError(t, err)
}
