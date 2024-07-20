package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomPost(t *testing.T) UserPost {
	randomBlog := createRandomBlog(t)
	arg := InsertUserPostParams{
		UserID: randomBlog.UserID.Int32,
		BlogID: randomBlog.ID,
	}
	userPost, err := testStore.InsertUserPost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, userPost)

	require.Equal(t, arg.UserID, userPost.UserID)
	require.Equal(t, arg.BlogID, userPost.BlogID)

	return userPost
}

func TestCreateUserPost(t *testing.T) {
	createRandomPost(t)
}

func TestGetUserPost(t *testing.T) {
	randomUserPost := createRandomPost(t)
	userPost, err := testStore.GetUserBlogs(context.Background(), randomUserPost.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, userPost)
}

func TestDeleteUserPost(t *testing.T) {
	randomUserPost := createRandomPost(t)
	err := testStore.DeleteUserPost(context.Background(), DeleteUserPostParams{
		UserID: randomUserPost.UserID,
		BlogID: randomUserPost.BlogID,
	})
	require.NoError(t, err)
	userPost, err := testStore.GetUserBlog(context.Background(), GetUserBlogParams{
		UserID: randomUserPost.UserID,
		ID:     randomUserPost.BlogID,
	})
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, userPost)
}
