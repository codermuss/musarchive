package db

import (
	"context"
	"testing"

	"github.com/mustafayilmazdev/musarchive/util"
	"github.com/stretchr/testify/require"
)

func createRandomComment(t *testing.T) Comment {
	randomBlog := createRandomBlog(t)
	arg := InsertCommentParams{
		BlogID:  randomBlog.ID,
		UserID:  randomBlog.UserID.Int32,
		Content: util.RandomDescription(),
	}
	comment, err := testStore.InsertComment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, comment)
	require.Equal(t, comment.BlogID, arg.BlogID)
	require.Equal(t, comment.UserID, arg.UserID)
	require.Equal(t, comment.Content, arg.Content)

	return comment
}

func TestCreateComment(t *testing.T) {
	createRandomComment(t)
}

func TestGetComments(t *testing.T) {
	randomComment := createRandomComment(t)
	comments, err := testStore.GetCommentsForBlog(context.Background(), randomComment.BlogID)
	require.NoError(t, err)
	require.NotEmpty(t, comments)
}

func TestDeleteComment(t *testing.T) {
	randomComment := createRandomComment(t)
	err := testStore.DeleteComment(context.Background(), randomComment.ID)
	require.NoError(t, err)
	category, err := testStore.GetCommentsForBlog(context.Background(), randomComment.ID)
	require.NoError(t, err)
	require.Empty(t, category)
}
