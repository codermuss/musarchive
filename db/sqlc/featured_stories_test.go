package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mustafayilmazdev/musarchive/util"
	"github.com/stretchr/testify/require"
)

func createRandomFeaturedStory(t *testing.T) FeaturedStory {
	randomBlog := createRandomPost(t)
	arg := InsertFeaturedStoryParams{
		PostID: randomBlog.ID,
		FeaturedDate: pgtype.Date{
			Valid: true,
			Time:  util.DateNow(),
		},
	}
	featuredStory, err := testStore.InsertFeaturedStory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, featuredStory)

	require.Equal(t, arg.PostID, featuredStory.PostID)
	require.Equal(t, arg.FeaturedDate.Time, featuredStory.FeaturedDate.Time)

	return featuredStory
}

func TestCreateRandomFeaturedStory(t *testing.T) {
	createRandomFeaturedStory(t)
}

func TestGetFeaturedStory(t *testing.T) {
	randomFeaturedStory := createRandomFeaturedStory(t)
	featuredStory, err := testStore.GetFeaturedStory(context.Background(), randomFeaturedStory.ID)
	require.NoError(t, err)
	require.NotEmpty(t, featuredStory)
	require.Equal(t, randomFeaturedStory.ID, featuredStory.ID)
	require.Equal(t, randomFeaturedStory.PostID, featuredStory.PostID)
	require.Equal(t, randomFeaturedStory.FeaturedDate, featuredStory.FeaturedDate)

}

func TestUpdateFeaturedStory(t *testing.T) {
	randomFeaturedStory := createRandomFeaturedStory(t)
	randomBlog := createRandomPost(t)
	updateFeaturedStory := UpdateFeaturedStoryParams{
		ID:     randomFeaturedStory.ID,
		PostID: randomBlog.ID,
		FeaturedDate: pgtype.Date{
			Valid: true,
			Time:  util.DateYesterday(),
		},
	}
	featuredStory, err := testStore.UpdateFeaturedStory(context.Background(), updateFeaturedStory)

	require.NoError(t, err)
	require.NotEmpty(t, featuredStory)
	require.Equal(t, updateFeaturedStory.ID, featuredStory.ID)
	require.Equal(t, updateFeaturedStory.PostID, featuredStory.PostID)
	require.Equal(t, updateFeaturedStory.FeaturedDate, featuredStory.FeaturedDate)

	require.NotEqual(t, randomFeaturedStory.FeaturedDate, featuredStory.FeaturedDate)
	require.WithinDuration(t, randomFeaturedStory.FeaturedDate.Time, featuredStory.FeaturedDate.Time, 24*time.Hour)
}

func TestDeleteFeaturedStory(t *testing.T) {
	randomFeaturedStory := createRandomFeaturedStory(t)
	err := testStore.DeleteFeaturedStory(context.Background(), randomFeaturedStory.ID)
	require.NoError(t, err)
	nilFeaturedStory, err := testStore.GetFeaturedStory(context.Background(), randomFeaturedStory.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, nilFeaturedStory)
}
