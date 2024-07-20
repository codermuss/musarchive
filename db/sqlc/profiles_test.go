package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mustafayilmazdev/musarchive/util"
	"github.com/stretchr/testify/require"
)

func createRandomUserProfile(t *testing.T) Profile {
	randomUser := createRandomUser(t)
	arg := InsertProfileParams{
		UserID: randomUser.ID,
		Bio: pgtype.Text{
			Valid:  true,
			String: util.RandomDescription(),
		},
	}
	profile, err := testStore.InsertProfile(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, profile)

	require.Equal(t, arg.UserID, profile.UserID)
	require.Equal(t, arg.Bio, profile.Bio)
	require.Equal(t, arg.BlogCount, profile.BlogCount)
	require.Equal(t, arg.FollowerCount, profile.FollowerCount)
	require.Equal(t, arg.LikeCount, profile.LikeCount)

	return profile
}

func TestCreateUserProfile(t *testing.T) {
	createRandomUserProfile(t)
}

func TestGetProfile(t *testing.T) {
	randomProfile := createRandomUserProfile(t)
	profile, err := testStore.GetProfile(context.Background(), randomProfile.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, profile)
	require.Equal(t, randomProfile.UserID, profile.UserID)
	require.Equal(t, randomProfile.Bio, profile.Bio)
	require.Equal(t, randomProfile.BlogCount, profile.BlogCount)
	require.Equal(t, randomProfile.FollowerCount, profile.FollowerCount)
	require.Equal(t, randomProfile.LikeCount, profile.LikeCount)

}

func TestUpdateProfileBio(t *testing.T) {
	randomProfile := createRandomUserProfile(t)
	updateProfile := UpdateProfileParams{
		UserID: randomProfile.UserID,
		Bio: pgtype.Text{
			Valid:  true,
			String: util.RandomDescription(),
		},
	}

	profile, err := testStore.UpdateProfile(context.Background(), updateProfile)

	require.NoError(t, err)
	require.NotEmpty(t, profile)

	require.Equal(t, randomProfile.UserID, profile.UserID)
	require.NotEqual(t, randomProfile.Bio, profile.Bio)
	require.Equal(t, randomProfile.BlogCount, profile.BlogCount)
	require.Equal(t, randomProfile.FollowerCount, profile.FollowerCount)
	require.Equal(t, randomProfile.LikeCount, profile.LikeCount)
}

func TestUpdateProfileBlogCount(t *testing.T) {
	randomProfile := createRandomUserProfile(t)
	updateProfile := UpdateProfileParams{
		UserID: randomProfile.UserID,
		BlogCount: pgtype.Int4{
			Valid: true,
			Int32: util.RandomLike(),
		},
	}

	profile, err := testStore.UpdateProfile(context.Background(), updateProfile)

	require.NoError(t, err)
	require.NotEmpty(t, profile)

	require.Equal(t, randomProfile.UserID, profile.UserID)
	require.Equal(t, randomProfile.Bio, profile.Bio)
	require.NotEqual(t, randomProfile.BlogCount, profile.BlogCount)
	require.Equal(t, randomProfile.FollowerCount, profile.FollowerCount)
	require.Equal(t, randomProfile.LikeCount, profile.LikeCount)
}

func TestUpdateProfileFollowerCount(t *testing.T) {
	randomProfile := createRandomUserProfile(t)
	updateProfile := UpdateProfileParams{
		UserID: randomProfile.UserID,
		FollowerCount: pgtype.Int4{
			Valid: true,
			Int32: util.RandomLike(),
		},
	}

	profile, err := testStore.UpdateProfile(context.Background(), updateProfile)

	require.NoError(t, err)
	require.NotEmpty(t, profile)

	require.Equal(t, randomProfile.UserID, profile.UserID)
	require.Equal(t, randomProfile.Bio, profile.Bio)
	require.Equal(t, randomProfile.BlogCount, profile.BlogCount)
	require.NotEqual(t, randomProfile.FollowerCount, profile.FollowerCount)
	require.Equal(t, randomProfile.LikeCount, profile.LikeCount)
}

func TestUpdateProfileLikeCount(t *testing.T) {
	randomProfile := createRandomUserProfile(t)
	updateProfile := UpdateProfileParams{
		UserID: randomProfile.UserID,
		LikeCount: pgtype.Int4{
			Valid: true,
			Int32: util.RandomLike(),
		},
	}

	profile, err := testStore.UpdateProfile(context.Background(), updateProfile)

	require.NoError(t, err)
	require.NotEmpty(t, profile)

	require.Equal(t, randomProfile.UserID, profile.UserID)
	require.Equal(t, randomProfile.Bio, profile.Bio)
	require.Equal(t, randomProfile.BlogCount, profile.BlogCount)
	require.Equal(t, randomProfile.FollowerCount, profile.FollowerCount)
	require.NotEqual(t, randomProfile.LikeCount, profile.LikeCount)
}

func TestUpdateProfileAll(t *testing.T) {
	randomProfile := createRandomUserProfile(t)
	updateProfile := UpdateProfileParams{
		UserID: randomProfile.UserID,
		Bio: pgtype.Text{
			Valid:  true,
			String: util.RandomDescription(),
		},
		BlogCount: pgtype.Int4{
			Valid: true,
			Int32: util.RandomLike(),
		},
		FollowerCount: pgtype.Int4{
			Valid: true,
			Int32: util.RandomLike(),
		},
		LikeCount: pgtype.Int4{
			Valid: true,
			Int32: util.RandomLike(),
		},
	}

	profile, err := testStore.UpdateProfile(context.Background(), updateProfile)

	require.NoError(t, err)
	require.NotEmpty(t, profile)

	require.Equal(t, randomProfile.UserID, profile.UserID)
	require.NotEqual(t, randomProfile.Bio, profile.Bio)
	require.NotEqual(t, randomProfile.BlogCount, profile.BlogCount)
	require.NotEqual(t, randomProfile.FollowerCount, profile.FollowerCount)
	require.NotEqual(t, randomProfile.LikeCount, profile.LikeCount)

	require.Equal(t, updateProfile.Bio, profile.Bio)
	require.Equal(t, updateProfile.BlogCount, profile.BlogCount)
	require.Equal(t, updateProfile.FollowerCount, profile.FollowerCount)
	require.Equal(t, updateProfile.LikeCount, profile.LikeCount)
}

func TestDeleteProfile(t *testing.T) {
	randomProfile := createRandomUserProfile(t)
	err := testStore.DeleteProfile(context.Background(), randomProfile.UserID)
	require.NoError(t, err)
	profile, err := testStore.GetProfile(context.Background(), randomProfile.UserID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, profile)
}
