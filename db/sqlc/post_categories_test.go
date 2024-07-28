package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomPostCategory(t *testing.T) PostCategory {
	post := createRandomUserPost(t)
	category := createRandomCategory(t)
	arg := InsertPostCategoryParams{
		PostID:     post.PostID,
		CategoryID: category.ID,
	}

	postCategory, err := testStore.InsertPostCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, postCategory)
	return postCategory
}

func TestCreatepostCategory(t *testing.T) {
	createRandomPostCategory(t)
}

func TestGetpostCategory(t *testing.T) {
	randompostCategory := createRandomPostCategory(t)
	categories, err := testStore.GetCategoriesForPost(context.Background(), randompostCategory.PostID)
	require.NoError(t, err)
	require.NotEmpty(t, categories)
}

func TestDeletepostCategory(t *testing.T) {
	randompostCategory := createRandomPostCategory(t)
	categories, err := testStore.GetCategoriesForPost(context.Background(), randompostCategory.PostID)
	require.NoError(t, err)
	require.NotEmpty(t, categories)
}
