package db

import (
	"context"
	"testing"

	"github.com/mustafayilmazdev/musarchive/util"
	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) Category {
	arg := util.RandomString(6)
	category, err := testStore.InsertCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg, category.Name)

	return category
}

func TestCreateRandomcategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetCategory(t *testing.T) {
	randomCategory := createRandomCategory(t)
	category, err := testStore.GetCategory(context.Background(), randomCategory.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category)
	require.Equal(t, randomCategory.ID, category.ID)
	require.Equal(t, randomCategory.Name, category.Name)

}

func TestUpdateCategory(t *testing.T) {
	randomcategory := createRandomCategory(t)
	newCategoryName := util.RandomString(6)
	updatecategory := UpdateCategoryParams{
		ID:   randomcategory.ID,
		Name: newCategoryName,
	}
	category, err := testStore.UpdateCategory(context.Background(), updatecategory)

	require.NoError(t, err)
	require.NotEmpty(t, category)
	require.Equal(t, randomcategory.ID, category.ID)
	require.Equal(t, randomcategory.ID, category.ID)
}

func TestDeleteCategory(t *testing.T) {
	randomcategory := createRandomCategory(t)
	err := testStore.DeleteCategory(context.Background(), randomcategory.ID)
	require.NoError(t, err)
	category, err := testStore.GetCategory(context.Background(), randomcategory.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, category)
}
