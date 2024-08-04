package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mustafayilmazdev/musarchive/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {

	arg := InsertUserParams{
		Username:       util.RandomUsername(),
		HashedPassword: util.RandomString(10),
		FullName:       util.RandomString(10),
		Email:          util.RandomEmail(),
		Avatar: pgtype.Text{
			Valid:  true,
			String: util.RandomImage(),
		},
		BirthDate: pgtype.Date{
			Valid: true,
			Time:  time.Date(1997, 7, 7, 7, 7, 7, 7, time.Local),
		},
	}
	user, err := testStore.InsertUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.NotEmpty(t, arg.Username, user.Username)
	require.Equal(t, arg.Username, user.Username)

	require.NotEmpty(t, arg.FullName, user.FullName)
	require.Equal(t, arg.FullName, user.FullName)

	require.NotEmpty(t, arg.Email, user.Email)
	require.Equal(t, arg.Email, arg.Email)

	require.NotEmpty(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.HashedPassword, arg.HashedPassword)

	require.NotEmpty(t, arg.Avatar, user.Avatar)
	require.Equal(t, arg.Avatar, arg.Avatar)

	require.NotEmpty(t, arg.BirthDate, user.BirthDate)
	require.Equal(t, arg.BirthDate, arg.BirthDate)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	randomUser := createRandomUser(t)
	user, err := testStore.GetUser(context.Background(), randomUser.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, randomUser.ID, user.ID)
	require.Equal(t, randomUser.Username, user.Username)
	require.Equal(t, randomUser.FullName, user.FullName)
	require.Equal(t, randomUser.Email, user.Email)
	require.Equal(t, randomUser.HashedPassword, user.HashedPassword)
	require.Equal(t, randomUser.PasswordChangedAt, user.PasswordChangedAt)
	require.Equal(t, randomUser.Avatar, user.Avatar)
	require.Equal(t, randomUser.BirthDate, user.BirthDate)
	require.Equal(t, randomUser.CreatedAt, user.CreatedAt)
	require.WithinDuration(t, randomUser.PasswordChangedAt, user.PasswordChangedAt, time.Second)
	require.WithinDuration(t, randomUser.CreatedAt, user.CreatedAt, time.Second)
}

func TestUpdateUser(t *testing.T) {
	randomUser := createRandomUser(t)
	updateUser := UpdateUserParams{
		ID: randomUser.ID,
		Username: pgtype.Text{
			Valid:  true,
			String: util.RandomUsername(),
		},
		HashedPassword: pgtype.Text{
			Valid:  true,
			String: util.RandomString(10),
		},
		FullName: pgtype.Text{
			Valid:  true,
			String: util.RandomString(10),
		},
		Email: pgtype.Text{
			Valid:  true,
			String: util.RandomEmail(),
		},
		Avatar: pgtype.Text{
			Valid:  true,
			String: util.RandomImage(),
		},
		BirthDate: pgtype.Date{
			Valid: true,
			Time:  time.Date(1995, 7, 7, 7, 7, 7, 7, time.Local),
		},
	}

	user, err := testStore.UpdateUser(context.Background(), updateUser)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, randomUser.ID, user.ID)
	require.Equal(t, updateUser.Username.String, user.Username)
	require.Equal(t, updateUser.FullName.String, user.FullName)
	require.Equal(t, updateUser.Email.String, user.Email)
	require.Equal(t, updateUser.HashedPassword.String, user.HashedPassword)
	require.Equal(t, updateUser.Avatar, user.Avatar)
}

func TestDeleteuser(t *testing.T) {
	randomUser := createRandomUser(t)
	err := testStore.DeleteUser(context.Background(), randomUser.ID)
	require.NoError(t, err)
	user, err := testStore.GetUser(context.Background(), randomUser.Username)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, user)
}
