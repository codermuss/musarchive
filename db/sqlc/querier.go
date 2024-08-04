// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateVerifyEmail(ctx context.Context, arg CreateVerifyEmailParams) (VerifyEmail, error)
	DeleteCategory(ctx context.Context, id int32) error
	DeleteComment(ctx context.Context, id int32) error
	DeleteFeaturedStory(ctx context.Context, id int32) error
	DeleteOnboarding(ctx context.Context, id int32) error
	DeletePost(ctx context.Context, id int32) error
	DeletePostCategory(ctx context.Context, arg DeletePostCategoryParams) error
	DeletePostTag(ctx context.Context, arg DeletePostTagParams) error
	DeleteProfile(ctx context.Context, userID int32) error
	DeleteSession(ctx context.Context, id uuid.UUID) error
	DeleteTag(ctx context.Context, id int32) error
	DeleteUser(ctx context.Context, id int32) error
	DeleteUserFollower(ctx context.Context, arg DeleteUserFollowerParams) error
	DeleteUserPost(ctx context.Context, arg DeleteUserPostParams) error
	GetCategories(ctx context.Context) ([]Category, error)
	GetCategoriesForPost(ctx context.Context, postID int32) ([]Category, error)
	GetCategory(ctx context.Context, id int32) (Category, error)
	GetComment(ctx context.Context, id int32) (Comment, error)
	GetCommentsForPost(ctx context.Context, postID int32) ([]Comment, error)
	GetFeaturedStory(ctx context.Context, id int32) (FeaturedStory, error)
	GetFollowedPosts(ctx context.Context, arg GetFollowedPostsParams) ([]Post, error)
	GetFollowersOfUser(ctx context.Context, userID int32) ([]GetFollowersOfUserRow, error)
	GetFollowingUsers(ctx context.Context, followerID int32) ([]GetFollowingUsersRow, error)
	GetOnboarding(ctx context.Context, id int32) (Onboarding, error)
	GetPost(ctx context.Context, id int32) (Post, error)
	GetPosts(ctx context.Context, arg GetPostsParams) ([]Post, error)
	GetPostsWithFilter(ctx context.Context, arg GetPostsWithFilterParams) ([]Post, error)
	GetProfile(ctx context.Context, userID int32) (Profile, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetTag(ctx context.Context, id int32) (Tag, error)
	GetTags(ctx context.Context) ([]Tag, error)
	GetTagsForPost(ctx context.Context, postID int32) ([]Tag, error)
	GetUser(ctx context.Context, username string) (User, error)
	GetUserPost(ctx context.Context, arg GetUserPostParams) (GetUserPostRow, error)
	GetUserPosts(ctx context.Context, userID int32) ([]GetUserPostsRow, error)
	InsertCategory(ctx context.Context, name string) (Category, error)
	InsertComment(ctx context.Context, arg InsertCommentParams) (Comment, error)
	InsertFeaturedStory(ctx context.Context, arg InsertFeaturedStoryParams) (FeaturedStory, error)
	InsertOnboarding(ctx context.Context, arg InsertOnboardingParams) (Onboarding, error)
	InsertPost(ctx context.Context, arg InsertPostParams) (Post, error)
	InsertPostCategory(ctx context.Context, arg InsertPostCategoryParams) (PostCategory, error)
	InsertPostTag(ctx context.Context, arg InsertPostTagParams) (PostTag, error)
	InsertProfile(ctx context.Context, arg InsertProfileParams) (Profile, error)
	InsertSession(ctx context.Context, arg InsertSessionParams) (Session, error)
	InsertTag(ctx context.Context, name string) (Tag, error)
	InsertUser(ctx context.Context, arg InsertUserParams) (User, error)
	InsertUserFollower(ctx context.Context, arg InsertUserFollowerParams) (UserFollower, error)
	InsertUserPost(ctx context.Context, arg InsertUserPostParams) (UserPost, error)
	ListOnboarding(ctx context.Context) ([]Onboarding, error)
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error)
	UpdateFeaturedStory(ctx context.Context, arg UpdateFeaturedStoryParams) (FeaturedStory, error)
	UpdateOnboarding(ctx context.Context, arg UpdateOnboardingParams) (Onboarding, error)
	UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error)
	UpdateProfile(ctx context.Context, arg UpdateProfileParams) (Profile, error)
	UpdateSession(ctx context.Context, arg UpdateSessionParams) error
	UpdateTag(ctx context.Context, arg UpdateTagParams) (Tag, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateVerifyEmail(ctx context.Context, arg UpdateVerifyEmailParams) (VerifyEmail, error)
}

var _ Querier = (*Queries)(nil)
