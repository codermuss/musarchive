package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomUserPost(t *testing.T) UserPost {
	randomPost := createRandomPost(t)
	arg := InsertUserPostParams{
		UserID: randomPost.UserID.Int32,
		PostID: randomPost.ID,
	}
	userPost, err := testStore.InsertUserPost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, userPost)

	require.Equal(t, arg.UserID, userPost.UserID)
	require.Equal(t, arg.PostID, userPost.PostID)

	return userPost
}

func TestCreateUserPost(t *testing.T) {
	createRandomUserPost(t)
}

func TestGetUserPost(t *testing.T) {
	randomUserPost := createRandomUserPost(t)
	userPost, err := testStore.GetUserPosts(context.Background(), randomUserPost.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, userPost)
}

func TestDeleteUserPost(t *testing.T) {
	randomUserPost := createRandomUserPost(t)
	err := testStore.DeleteUserPost(context.Background(), DeleteUserPostParams{
		UserID: randomUserPost.UserID,
		PostID: randomUserPost.PostID,
	})
	require.NoError(t, err)
	userPost, err := testStore.GetUserPost(context.Background(), GetUserPostParams{
		UserID: randomUserPost.UserID,
		ID:     randomUserPost.PostID,
	})
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, userPost)
}
