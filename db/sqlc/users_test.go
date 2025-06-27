package db

import (
	"context"
	"testing"
	"time"

	"github.com/dborowsky/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	return user
}

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: "secret",
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, user.Username, arg.Username)
	require.Equal(t, user.HashedPassword, arg.HashedPassword)
	require.Equal(t, user.FullName, arg.FullName)
	require.NotZero(t, user.Email, arg.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotEmpty(t, user.CreatedAt)
}

func TestGetUserByUsername(t *testing.T) {
	user := createRandomUser(t)
	got, err := testQueries.GetUser(context.Background(), user.Username)

	require.NoError(t, err)
	require.NotEmpty(t, got)

	require.Equal(t, got.Username, user.Username)
	require.Equal(t, got.HashedPassword, user.HashedPassword)
	require.Equal(t, got.FullName, user.FullName)
	require.Equal(t, got.Email, user.Email)
	require.WithinDuration(t, got.CreatedAt, user.CreatedAt, time.Second)
	require.WithinDuration(t, got.PasswordChangedAt, user.PasswordChangedAt, time.Second)
}
