package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomBlogCategory(t *testing.T) BlogCategory {
	blog := createRandomBlog(t)
	category := createRandomCategory(t)
	arg := InsertBlogCategoryParams{
		BlogID:     blog.ID,
		CategoryID: category.ID,
	}

	blogCategory, err := testStore.InsertBlogCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, blogCategory)
	return blogCategory
}

func TestCreateBlogCategory(t *testing.T) {
	createRandomBlogCategory(t)
}

func TestGetBlogCategory(t *testing.T) {
	randomBlogCategory := createRandomBlogCategory(t)
	categories, err := testStore.GetCategoriesForBlog(context.Background(), randomBlogCategory.BlogID)
	require.NoError(t, err)
	require.NotEmpty(t, categories)
}

func TestDeleteBlogCategory(t *testing.T) {
	randomBlogCategory := createRandomBlogCategory(t)
	categories, err := testStore.GetCategoriesForBlog(context.Background(), randomBlogCategory.BlogID)
	require.NoError(t, err)
	require.NotEmpty(t, categories)
}
