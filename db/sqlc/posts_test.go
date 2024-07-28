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
	blog, err := testStore.InsertPost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, blog)

	require.Equal(t, arg.UserID, blog.UserID)
	require.Equal(t, arg.Title, blog.Title)
	require.Equal(t, arg.Summary, blog.Summary)
	require.Equal(t, arg.Content, blog.Content)
	require.Equal(t, arg.Likes, blog.Likes)
	require.Equal(t, arg.CreatedAt, blog.CreatedAt)
	require.Equal(t, arg.UpdatedAt, blog.UpdatedAt)

	return blog
}

func TestCreateBlog(t *testing.T) {
	createRandomPost(t)
}

func TestGetBlog(t *testing.T) {
	randomBlog := createRandomPost(t)
	blog, err := testStore.GetPost(context.Background(), randomBlog.ID)
	require.NoError(t, err)
	require.NotEmpty(t, blog)

	require.Equal(t, randomBlog.UserID, blog.UserID)
	require.Equal(t, randomBlog.Title, blog.Title)
	require.Equal(t, randomBlog.Summary, blog.Summary)
	require.Equal(t, randomBlog.Content, blog.Content)
	require.Equal(t, randomBlog.Likes, blog.Likes)
	require.Equal(t, randomBlog.CreatedAt, blog.CreatedAt)
	require.Equal(t, randomBlog.UpdatedAt, blog.UpdatedAt)

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
	blog, err := testStore.UpdatePost(context.Background(), updatePost)
	require.NoError(t, err)
	require.NotEmpty(t, blog)
	require.Equal(t, randomBlog.UserID, blog.UserID)
	require.NotEqual(t, randomBlog.Title, blog.Title)
	require.Equal(t, randomBlog.Summary, blog.Summary)
	require.Equal(t, randomBlog.Content, blog.Content)
	require.Equal(t, randomBlog.Likes, blog.Likes)
	require.Equal(t, randomBlog.CreatedAt, blog.CreatedAt)
	require.Equal(t, randomBlog.UpdatedAt, blog.UpdatedAt)
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
	blog, err := testStore.UpdatePost(context.Background(), updatePost)
	require.NoError(t, err)
	require.NotEmpty(t, blog)
	require.Equal(t, randomBlog.UserID, blog.UserID)
	require.Equal(t, randomBlog.Title, blog.Title)
	require.NotEqual(t, randomBlog.Summary, blog.Summary)
	require.Equal(t, randomBlog.Content, blog.Content)
	require.Equal(t, randomBlog.Likes, blog.Likes)
	require.Equal(t, randomBlog.CreatedAt, blog.CreatedAt)
	require.Equal(t, randomBlog.UpdatedAt, blog.UpdatedAt)
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
	blog, err := testStore.UpdatePost(context.Background(), updatePost)
	require.NoError(t, err)
	require.NotEmpty(t, blog)

	require.NotEqual(t, randomBlog.Content, blog.Content)
	require.Equal(t, updatePost.Content.String, blog.Content)

	require.Equal(t, randomBlog.UserID, blog.UserID)
	require.Equal(t, randomBlog.Title, blog.Title)
	require.Equal(t, randomBlog.Summary, blog.Summary)
	require.Equal(t, randomBlog.Likes, blog.Likes)
	require.Equal(t, randomBlog.CreatedAt, blog.CreatedAt)
	require.Equal(t, randomBlog.UpdatedAt, blog.UpdatedAt)
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
	blog, err := testStore.UpdatePost(context.Background(), updatePost)
	require.NoError(t, err)
	require.NotEmpty(t, blog)

	require.NotEqual(t, randomBlog.CoverImage, blog.CoverImage)
	require.Equal(t, updatePost.CoverImage, blog.CoverImage)

	require.Equal(t, randomBlog.UserID, blog.UserID)
	require.Equal(t, randomBlog.Title, blog.Title)
	require.Equal(t, randomBlog.Content, blog.Content)
	require.Equal(t, randomBlog.Likes, blog.Likes)
	require.Equal(t, randomBlog.CreatedAt, blog.CreatedAt)
	require.Equal(t, randomBlog.UpdatedAt, blog.UpdatedAt)
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
	blog, err := testStore.UpdatePost(context.Background(), updatePost)
	require.NoError(t, err)
	require.NotEmpty(t, blog)

	require.NotEqual(t, randomBlog.CoverImage, blog.CoverImage)
	require.NotEqual(t, randomBlog.Likes, blog.Likes)
	require.NotEqual(t, randomBlog.Title, blog.Title)
	require.NotEqual(t, randomBlog.Summary, blog.Summary)
	require.NotEqual(t, randomBlog.Content, blog.Content)

	require.Equal(t, updatePost.Title.String, blog.Title)
	require.Equal(t, updatePost.Summary.String, blog.Summary)
	require.Equal(t, updatePost.Content.String, blog.Content)
	require.Equal(t, updatePost.Likes, blog.Likes)
	require.Equal(t, updatePost.CoverImage, blog.CoverImage)

	require.Equal(t, randomBlog.UserID, blog.UserID)
	require.Equal(t, randomBlog.CreatedAt, blog.CreatedAt)
	require.Equal(t, randomBlog.UpdatedAt, blog.UpdatedAt)
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
