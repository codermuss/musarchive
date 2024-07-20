package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomBlogTag(t *testing.T, tagCount int) []BlogTag {
	var blogTags []BlogTag
	randomBlog := createRandomBlog(t)

	for i := 0; i < tagCount; i++ {
		blogTag, err := testStore.InsertBlogTag(context.Background(), InsertBlogTagParams{
			BlogID: randomBlog.ID,
			TagID:  createRandomTag(t).ID,
		})
		require.NoError(t, err)
		require.NotEmpty(t, blogTag)
		blogTags = append(blogTags, blogTag)

	}
	return blogTags
}

func TestCreateBlogTag(t *testing.T) {
	createRandomBlogTag(t, 1)
}

func TestGetBlogTag(t *testing.T) {
	randomBlogTag := createRandomBlogTag(t, 1)
	tags, err := testStore.GetTagsForBlog(context.Background(), randomBlogTag[0].BlogID)
	require.NoError(t, err)
	require.NotEmpty(t, tags)
}

func TestDeleteBlogTag(t *testing.T) {
	randomBlogTag := createRandomBlogTag(t, 5)
	tags, err := testStore.GetTagsForBlog(context.Background(), randomBlogTag[0].BlogID)
	require.NoError(t, err)
	require.NotEmpty(t, tags)
	fmt.Println(tags)
	err = testStore.DeleteBlogTag(context.Background(), DeleteBlogTagParams{
		BlogID: randomBlogTag[0].BlogID,
		TagID:  tags[len(tags)-1].ID,
	})
	fmt.Println(testStore.GetTagsForBlog(context.Background(), randomBlogTag[0].BlogID))
	require.NoError(t, err)
}
