package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mustafayilmazdev/musarchive/util"

	"github.com/stretchr/testify/require"
)

func createRandomBlog(t *testing.T) Blog {
	user := createRandomUser(t)
	arg := InsertBlogParams{
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
	blog, err := testStore.InsertBlog(context.Background(), arg)
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
	createRandomBlog(t)
}

func TestGetBlog(t *testing.T) {
	randomBlog := createRandomBlog(t)
	blog, err := testStore.GetBlog(context.Background(), randomBlog.ID)
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

func TestUpdateBlogTitle(t *testing.T) {
	randomBlog := createRandomBlog(t)
	updateBlog := UpdateBlogParams{
		ID: randomBlog.ID,
		Title: pgtype.Text{
			Valid:  true,
			String: util.RandomTitle(),
		},
	}
	blog, err := testStore.UpdateBlog(context.Background(), updateBlog)
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

func TestUpdateBlogSummary(t *testing.T) {
	randomBlog := createRandomBlog(t)
	updateBlog := UpdateBlogParams{
		ID: randomBlog.ID,
		Summary: pgtype.Text{
			Valid:  true,
			String: util.RandomString(10),
		},
	}
	blog, err := testStore.UpdateBlog(context.Background(), updateBlog)
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

func TestUpdateBlogContent(t *testing.T) {
	randomBlog := createRandomBlog(t)
	updateBlog := UpdateBlogParams{
		ID: randomBlog.ID,
		Content: pgtype.Text{
			Valid:  true,
			String: util.RandomDescription(),
		},
	}
	blog, err := testStore.UpdateBlog(context.Background(), updateBlog)
	require.NoError(t, err)
	require.NotEmpty(t, blog)

	require.NotEqual(t, randomBlog.Content, blog.Content)
	require.Equal(t, updateBlog.Content.String, blog.Content)

	require.Equal(t, randomBlog.UserID, blog.UserID)
	require.Equal(t, randomBlog.Title, blog.Title)
	require.Equal(t, randomBlog.Summary, blog.Summary)
	require.Equal(t, randomBlog.Likes, blog.Likes)
	require.Equal(t, randomBlog.CreatedAt, blog.CreatedAt)
	require.Equal(t, randomBlog.UpdatedAt, blog.UpdatedAt)
}

func TestUpdateBlogCover(t *testing.T) {
	randomBlog := createRandomBlog(t)
	updateBlog := UpdateBlogParams{
		ID: randomBlog.ID,
		CoverImage: pgtype.Text{
			Valid:  true,
			String: util.RandomImage() + "/update",
		},
	}
	blog, err := testStore.UpdateBlog(context.Background(), updateBlog)
	require.NoError(t, err)
	require.NotEmpty(t, blog)

	require.NotEqual(t, randomBlog.CoverImage, blog.CoverImage)
	require.Equal(t, updateBlog.CoverImage, blog.CoverImage)

	require.Equal(t, randomBlog.UserID, blog.UserID)
	require.Equal(t, randomBlog.Title, blog.Title)
	require.Equal(t, randomBlog.Content, blog.Content)
	require.Equal(t, randomBlog.Likes, blog.Likes)
	require.Equal(t, randomBlog.CreatedAt, blog.CreatedAt)
	require.Equal(t, randomBlog.UpdatedAt, blog.UpdatedAt)
}

func TestUpdateBlogAll(t *testing.T) {
	randomBlog := createRandomBlog(t)
	updateBlog := UpdateBlogParams{
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
	blog, err := testStore.UpdateBlog(context.Background(), updateBlog)
	require.NoError(t, err)
	require.NotEmpty(t, blog)

	require.NotEqual(t, randomBlog.CoverImage, blog.CoverImage)
	require.NotEqual(t, randomBlog.Likes, blog.Likes)
	require.NotEqual(t, randomBlog.Title, blog.Title)
	require.NotEqual(t, randomBlog.Summary, blog.Summary)
	require.NotEqual(t, randomBlog.Content, blog.Content)

	require.Equal(t, updateBlog.Title.String, blog.Title)
	require.Equal(t, updateBlog.Summary.String, blog.Summary)
	require.Equal(t, updateBlog.Content.String, blog.Content)
	require.Equal(t, updateBlog.Likes, blog.Likes)
	require.Equal(t, updateBlog.CoverImage, blog.CoverImage)

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
