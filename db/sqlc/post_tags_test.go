package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomPostTag(t *testing.T, tagCount int) []PostTag {
	var postTags []PostTag
	randomPost := createRandomPost(t)

	for i := 0; i < tagCount; i++ {
		PostTag, err := testStore.InsertPostTag(context.Background(), InsertPostTagParams{
			PostID: randomPost.ID,
			TagID:  createRandomTag(t).ID,
		})
		require.NoError(t, err)
		require.NotEmpty(t, PostTag)
		postTags = append(postTags, PostTag)

	}
	return postTags
}

func TestCreatePostTag(t *testing.T) {
	createRandomPostTag(t, 1)
}

func TestGetPostTag(t *testing.T) {
	randomPostTag := createRandomPostTag(t, 1)
	tags, err := testStore.GetTagsForPost(context.Background(), randomPostTag[0].PostID)
	require.NoError(t, err)
	require.NotEmpty(t, tags)
}

func TestDeletePostTag(t *testing.T) {
	randomPostTag := createRandomPostTag(t, 1)
	tags, err := testStore.GetTagsForPost(context.Background(), randomPostTag[0].PostID)
	require.NoError(t, err)
	require.NotEmpty(t, tags)
	err = testStore.DeletePostTag(context.Background(), DeletePostTagParams{
		PostID: randomPostTag[0].PostID,
		TagID:  tags[len(tags)-1].ID,
	})
	fmt.Println(testStore.GetTagsForPost(context.Background(), randomPostTag[0].PostID))
	require.NoError(t, err)
}
