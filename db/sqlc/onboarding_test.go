package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mustafayilmazdev/musarchive/util"
	"github.com/stretchr/testify/require"
)

func createRandomOnboarding(t *testing.T) Onboarding {
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
	return onboarding
}

func TestCreateRandomOnboarding(t *testing.T) {
	createRandomOnboarding(t)
}

func TestGetOnboarding(t *testing.T) {
	randomOnboarding := createRandomOnboarding(t)
	onboarding, err := testStore.GetOnboarding(context.Background(), randomOnboarding.ID)
	require.NoError(t, err)
	require.NotEmpty(t, onboarding)
	require.Equal(t, randomOnboarding.ID, onboarding.ID)
	require.Equal(t, randomOnboarding.Title, onboarding.Title)
	require.Equal(t, randomOnboarding.Description, onboarding.Description)
	require.Equal(t, randomOnboarding.Image, onboarding.Image)
}

func TestUpdateOnboarding(t *testing.T) {
	randomOnboarding := createRandomOnboarding(t)
	updateOnboarding := UpdateOnboardingParams{
		ID:          randomOnboarding.ID,
		Title:       util.RandomTitle(),
		Description: util.RandomDescription(),
		Image: pgtype.Text{
			Valid:  true,
			String: util.RandomImage(),
		},
	}
	onboarding, err := testStore.UpdateOnboarding(context.Background(), updateOnboarding)

	require.NoError(t, err)
	require.NotEmpty(t, onboarding)
	require.Equal(t, randomOnboarding.ID, onboarding.ID)
	require.Equal(t, updateOnboarding.Title, onboarding.Title)
	require.Equal(t, updateOnboarding.Description, onboarding.Description)
	require.Equal(t, updateOnboarding.Image, onboarding.Image)
}

func TestDeleteOnboarding(t *testing.T) {
	randomOnboarding := createRandomOnboarding(t)
	err := testStore.DeleteOnboarding(context.Background(), randomOnboarding.ID)
	require.NoError(t, err)
	onboarding, err := testStore.GetOnboarding(context.Background(), randomOnboarding.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, onboarding)
}
