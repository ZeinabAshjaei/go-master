package db

import (
	"context"
	"github.com/ZeinabAshjaei/go-master/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       utils.RandomOwner(),
		HashedPassword: "secret",
		FullName:       utils.RandomOwner(),
		Email:          utils.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.FullName, arg.FullName)
	require.Equal(t, user.Email, arg.Email)
	require.Equal(t, user.Username, arg.Username)
	require.Equal(t, user.HashedPassword, arg.HashedPassword)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestQueries_CreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestQueries_GetUser(t *testing.T) {
	insertedUser := createRandomUser(t)
	retrievedUser, err := testQueries.GetUser(context.Background(), insertedUser.Username)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedUser)

	require.Equal(t, insertedUser.Username, retrievedUser.Username)
	require.Equal(t, insertedUser.Email, insertedUser.Email)
	require.Equal(t, insertedUser.FullName, insertedUser.FullName)
	require.Equal(t, insertedUser.HashedPassword, retrievedUser.HashedPassword)
	require.WithinDuration(t, insertedUser.CreatedAt, retrievedUser.CreatedAt, time.Second)
}
