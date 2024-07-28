package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mustafayilmazdev/musarchive/util"

	"github.com/stretchr/testify/require"
)

func createRandomPost(t *testing.T) Post {
	user := createRandomUser(t)
	arg := InsertPostParams{
		UserID: pgtype.Int4{
			Valid: true,
			Int32: user.ID,
		},
		Title:   util.RandomTitle(),
		Summary: util.RandomString(15),
		Content: util.RandomDescription(),
		CoverImage: pgtype.Text{
			Valid:  true,
			String: util.RandomImage(),
		},
	}
	post, err := testStore.InsertPost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post)

	require.Equal(t, arg.UserID, post.UserID)
	require.Equal(t, arg.Title, post.Title)
	require.Equal(t, arg.Summary, post.Summary)
	require.Equal(t, arg.Content, post.Content)
	require.Equal(t, arg.Likes, post.Likes)
	require.Equal(t, arg.CreatedAt, post.CreatedAt)
	require.Equal(t, arg.UpdatedAt, post.UpdatedAt)

	return post
}

func TestCreateBlog(t *testing.T) {
	createRandomPost(t)
}

func TestGetBlog(t *testing.T) {
	randomBlog := createRandomPost(t)
	post, err := testStore.GetPost(context.Background(), randomBlog.ID)
	require.NoError(t, err)
	require.NotEmpty(t, post)

	require.Equal(t, randomBlog.UserID, post.UserID)
	require.Equal(t, randomBlog.Title, post.Title)
	require.Equal(t, randomBlog.Summary, post.Summary)
	require.Equal(t, randomBlog.Content, post.Content)
	require.Equal(t, randomBlog.Likes, post.Likes)
	require.Equal(t, randomBlog.CreatedAt, post.CreatedAt)
	require.Equal(t, randomBlog.UpdatedAt, post.UpdatedAt)

}

func TestUpdatePostTitle(t *testing.T) {
	randomBlog := createRandomPost(t)
	updatePost := UpdatePostParams{
		ID: randomBlog.ID,
		Title: pgtype.Text{
			Valid:  true,
			String: util.RandomTitle(),
		},
	}
	post, err := testStore.UpdatePost(context.Background(), updatePost)
	require.NoError(t, err)
	require.NotEmpty(t, post)
	require.Equal(t, randomBlog.UserID, post.UserID)
	require.NotEqual(t, randomBlog.Title, post.Title)
	require.Equal(t, randomBlog.Summary, post.Summary)
	require.Equal(t, randomBlog.Content, post.Content)
	require.Equal(t, randomBlog.Likes, post.Likes)
	require.Equal(t, randomBlog.CreatedAt, post.CreatedAt)
	require.Equal(t, randomBlog.UpdatedAt, post.UpdatedAt)
}

func TestUpdatePostSummary(t *testing.T) {
	randomBlog := createRandomPost(t)
	updatePost := UpdatePostParams{
		ID: randomBlog.ID,
		Summary: pgtype.Text{
			Valid:  true,
			String: util.RandomString(10),
		},
	}
	post, err := testStore.UpdatePost(context.Background(), updatePost)
	require.NoError(t, err)
	require.NotEmpty(t, post)
	require.Equal(t, randomBlog.UserID, post.UserID)
	require.Equal(t, randomBlog.Title, post.Title)
	require.NotEqual(t, randomBlog.Summary, post.Summary)
	require.Equal(t, randomBlog.Content, post.Content)
	require.Equal(t, randomBlog.Likes, post.Likes)
	require.Equal(t, randomBlog.CreatedAt, post.CreatedAt)
	require.Equal(t, randomBlog.UpdatedAt, post.UpdatedAt)
}

func TestUpdatePostContent(t *testing.T) {
	randomBlog := createRandomPost(t)
	updatePost := UpdatePostParams{
		ID: randomBlog.ID,
		Content: pgtype.Text{
			Valid:  true,
			String: util.RandomDescription(),
		},
	}
	post, err := testStore.UpdatePost(context.Background(), updatePost)
	require.NoError(t, err)
	require.NotEmpty(t, post)

	require.NotEqual(t, randomBlog.Content, post.Content)
	require.Equal(t, updatePost.Content.String, post.Content)

	require.Equal(t, randomBlog.UserID, post.UserID)
	require.Equal(t, randomBlog.Title, post.Title)
	require.Equal(t, randomBlog.Summary, post.Summary)
	require.Equal(t, randomBlog.Likes, post.Likes)
	require.Equal(t, randomBlog.CreatedAt, post.CreatedAt)
	require.Equal(t, randomBlog.UpdatedAt, post.UpdatedAt)
}

func TestUpdatePostCover(t *testing.T) {
	randomBlog := createRandomPost(t)
	updatePost := UpdatePostParams{
		ID: randomBlog.ID,
		CoverImage: pgtype.Text{
			Valid:  true,
			String: util.RandomImage() + "/update",
		},
	}
	post, err := testStore.UpdatePost(context.Background(), updatePost)
	require.NoError(t, err)
	require.NotEmpty(t, post)

	require.NotEqual(t, randomBlog.CoverImage, post.CoverImage)
	require.Equal(t, updatePost.CoverImage, post.CoverImage)

	require.Equal(t, randomBlog.UserID, post.UserID)
	require.Equal(t, randomBlog.Title, post.Title)
	require.Equal(t, randomBlog.Content, post.Content)
	require.Equal(t, randomBlog.Likes, post.Likes)
	require.Equal(t, randomBlog.CreatedAt, post.CreatedAt)
	require.Equal(t, randomBlog.UpdatedAt, post.UpdatedAt)
}

func TestUpdatePostAll(t *testing.T) {
	randomBlog := createRandomPost(t)
	updatePost := UpdatePostParams{
		ID: randomBlog.ID,
		Title: pgtype.Text{
			Valid:  true,
			String: util.RandomTitle(),
		},
		Summary: pgtype.Text{
			Valid:  true,
			String: util.RandomString(10),
		},
		Content: pgtype.Text{
			Valid:  true,
			String: util.RandomDescription(),
		},
		Likes: pgtype.Int4{
			Valid: true,
			Int32: util.RandomLike(),
		},
		CoverImage: pgtype.Text{
			Valid:  true,
			String: util.RandomImage() + "/update",
		},
	}
	post, err := testStore.UpdatePost(context.Background(), updatePost)
	require.NoError(t, err)
	require.NotEmpty(t, post)

	require.NotEqual(t, randomBlog.CoverImage, post.CoverImage)
	require.NotEqual(t, randomBlog.Likes, post.Likes)
	require.NotEqual(t, randomBlog.Title, post.Title)
	require.NotEqual(t, randomBlog.Summary, post.Summary)
	require.NotEqual(t, randomBlog.Content, post.Content)

	require.Equal(t, updatePost.Title.String, post.Title)
	require.Equal(t, updatePost.Summary.String, post.Summary)
	require.Equal(t, updatePost.Content.String, post.Content)
	require.Equal(t, updatePost.Likes, post.Likes)
	require.Equal(t, updatePost.CoverImage, post.CoverImage)

	require.Equal(t, randomBlog.UserID, post.UserID)
	require.Equal(t, randomBlog.CreatedAt, post.CreatedAt)
	require.Equal(t, randomBlog.UpdatedAt, post.UpdatedAt)
}

func TestDeleteBlog(t *testing.T) {
	randomcategory := createRandomCategory(t)
	err := testStore.DeleteCategory(context.Background(), randomcategory.ID)
	require.NoError(t, err)
	category, err := testStore.GetCategory(context.Background(), randomcategory.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, category)
}
