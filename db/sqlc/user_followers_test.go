package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomUserFollower(t *testing.T) UserFollower {
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)
	arg := InsertUserFollowerParams{
		UserID:     user1.ID,
		FollowerID: user2.ID,
	}
	userFollower, err := testStore.InsertUserFollower(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, userFollower)

	require.Equal(t, arg.UserID, userFollower.UserID)
	require.Equal(t, arg.FollowerID, userFollower.FollowerID)

	return userFollower
}

func TestCreateUserFollower(t *testing.T) {
	createRandomUserFollower(t)
}

func TestGetFollowersOfUsers(t *testing.T) {
	randomUser := createRandomUserFollower(t)
	userFollowers, err := testStore.GetFollowersOfUser(context.Background(), randomUser.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, userFollowers)
}

func TestGetFollowingUsers(t *testing.T) {
	randomUser := createRandomUserFollower(t)
	userFollowers, err := testStore.GetFollowingUsers(context.Background(), randomUser.FollowerID)
	require.NoError(t, err)
	require.NotEmpty(t, userFollowers)
}

func TestDeleteUserFollower(t *testing.T) {
	randomUser := createRandomUserFollower(t)

	err := testStore.DeleteUserFollower(context.Background(), DeleteUserFollowerParams{
		UserID:     randomUser.UserID,
		FollowerID: randomUser.FollowerID,
	})
	require.NoError(t, err)

	followers, err := testStore.GetFollowersOfUser(context.Background(), randomUser.UserID)
	require.NoError(t, err)
	require.Empty(t, followers)
}
