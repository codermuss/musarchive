package db

import (
	"context"
	"testing"

	"github.com/mustafayilmazdev/musarchive/util"
	"github.com/stretchr/testify/require"
)

func createRandomTag(t *testing.T) Tag {
	arg := util.RandomString(6)
	tag, err := testStore.InsertTag(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, tag)

	require.Equal(t, arg, tag.Name)

	return tag
}

func TestCreateTag(t *testing.T) {
	createRandomTag(t)
}

func TestGetTag(t *testing.T) {
	randomtag := createRandomTag(t)
	tag, err := testStore.GetTag(context.Background(), randomtag.ID)
	require.NoError(t, err)
	require.NotEmpty(t, tag)
	require.Equal(t, randomtag.ID, tag.ID)
	require.Equal(t, randomtag.Name, tag.Name)

}
func TestGetTags(t *testing.T) {
	createRandomTag(t)
	createRandomTag(t)
	createRandomTag(t)

	tags, err := testStore.GetTags(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, tags)
	require.True(t, len(tags) > 2)

}

func TestUpdateTag(t *testing.T) {
	randomtag := createRandomTag(t)
	newTagName := util.RandomString(6)
	updatetag := UpdateTagParams{
		ID:   randomtag.ID,
		Name: newTagName,
	}
	tag, err := testStore.UpdateTag(context.Background(), updatetag)

	require.NoError(t, err)
	require.NotEmpty(t, tag)
	require.Equal(t, randomtag.ID, tag.ID)
	require.Equal(t, randomtag.ID, tag.ID)
}

func TestDeleteTag(t *testing.T) {
	randomtag := createRandomTag(t)
	err := testStore.DeleteTag(context.Background(), randomtag.ID)
	require.NoError(t, err)
	tag, err := testStore.GetTag(context.Background(), randomtag.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, tag)
}
