package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mustafayilmazdev/musarchive/util"
	"github.com/stretchr/testify/require"
)

func TestCreateRandomOnboarding(t *testing.T) {
	title, description, image := util.RandomTitle(), util.RandomDescription(), util.RandomImage()
	arg := InsertOnboardingParams{
		Title:       title,
		Description: description,
		Image: pgtype.Text{
			Valid:  true,
			String: image,
		},
	}
	onboarding, err := testStore.InsertOnboarding(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, onboarding)

	require.NotEmpty(t, title, onboarding.Title)
	require.Equal(t, title, onboarding.Title)

	require.NotEmpty(t, description, onboarding.Description)
	require.Equal(t, description, onboarding.Description)

	require.NotEmpty(t, image, onboarding.Image.String)
	require.Equal(t, image, onboarding.Image.String)
}
