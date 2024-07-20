package db

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/mustafayilmazdev/musarchive/util"
	"github.com/stretchr/testify/require"
)

func createSession(t *testing.T) Session {
	randomUser := createRandomUser(t)
	arg := InsertSessionParams{
		ID:           uuid.New(),
		UserID:       randomUser.ID,
		RefreshToken: util.RandomString(10),
		UserAgent:    util.RandomString(10),
		ClientIp:     util.RandomString(10),
		IsBlocked:    false,
		CreatedAt:    util.DateYesterday(),
		ExpiresAt:    util.DateNow(),
	}

	session, err := testStore.InsertSession(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, session)

	return session
}

func TestCreateCreateSession(t *testing.T) {
	createSession(t)
}

func TestGetSession(t *testing.T) {
	session := createSession(t)
	comments, err := testStore.GetSession(context.Background(), session.ID)
	require.NoError(t, err)
	require.NotEmpty(t, comments)
}
