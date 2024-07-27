// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	DeleteBlog(ctx context.Context, id int32) error
	DeleteBlogCategory(ctx context.Context, arg DeleteBlogCategoryParams) error
	DeleteBlogTag(ctx context.Context, arg DeleteBlogTagParams) error
	DeleteCategory(ctx context.Context, id int32) error
	DeleteComment(ctx context.Context, id int32) error
	DeleteFeaturedStory(ctx context.Context, id int32) error
	DeleteOnboarding(ctx context.Context, id int32) error
	DeleteProfile(ctx context.Context, userID int32) error
	DeleteTag(ctx context.Context, id int32) error
	DeleteUser(ctx context.Context, id int32) error
	DeleteUserFollower(ctx context.Context, arg DeleteUserFollowerParams) error
	DeleteUserPost(ctx context.Context, arg DeleteUserPostParams) error
	GetBlog(ctx context.Context, id int32) (Blog, error)
	GetBlogs(ctx context.Context, arg GetBlogsParams) ([]Blog, error)
	GetCategoriesForBlog(ctx context.Context, blogID int32) ([]Category, error)
	GetCategory(ctx context.Context, id int32) (Category, error)
	GetCommentsForBlog(ctx context.Context, blogID int32) ([]Comment, error)
	GetFeaturedStory(ctx context.Context, id int32) (FeaturedStory, error)
	GetFollowersOfUser(ctx context.Context, userID int32) ([]GetFollowersOfUserRow, error)
	GetFollowingUsers(ctx context.Context, followerID int32) ([]GetFollowingUsersRow, error)
	GetOnboarding(ctx context.Context, id int32) (Onboarding, error)
	GetProfile(ctx context.Context, userID int32) (Profile, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetTag(ctx context.Context, id int32) (Tag, error)
	GetTagsForBlog(ctx context.Context, blogID int32) ([]Tag, error)
	GetUser(ctx context.Context, username string) (User, error)
	GetUserBlog(ctx context.Context, arg GetUserBlogParams) (GetUserBlogRow, error)
	GetUserBlogs(ctx context.Context, userID int32) ([]GetUserBlogsRow, error)
	InsertBlog(ctx context.Context, arg InsertBlogParams) (Blog, error)
	InsertBlogCategory(ctx context.Context, arg InsertBlogCategoryParams) (BlogCategory, error)
	InsertBlogTag(ctx context.Context, arg InsertBlogTagParams) (BlogTag, error)
	InsertCategory(ctx context.Context, name string) (Category, error)
	InsertComment(ctx context.Context, arg InsertCommentParams) (Comment, error)
	InsertFeaturedStory(ctx context.Context, arg InsertFeaturedStoryParams) (FeaturedStory, error)
	InsertOnboarding(ctx context.Context, arg InsertOnboardingParams) (Onboarding, error)
	InsertProfile(ctx context.Context, arg InsertProfileParams) (Profile, error)
	InsertSession(ctx context.Context, arg InsertSessionParams) (Session, error)
	InsertTag(ctx context.Context, name string) (Tag, error)
	InsertUser(ctx context.Context, arg InsertUserParams) (User, error)
	InsertUserFollower(ctx context.Context, arg InsertUserFollowerParams) (UserFollower, error)
	InsertUserPost(ctx context.Context, arg InsertUserPostParams) (UserPost, error)
	ListOnboarding(ctx context.Context) ([]Onboarding, error)
	UpdateBlog(ctx context.Context, arg UpdateBlogParams) (Blog, error)
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error)
	UpdateFeaturedStory(ctx context.Context, arg UpdateFeaturedStoryParams) (FeaturedStory, error)
	UpdateOnboarding(ctx context.Context, arg UpdateOnboardingParams) (Onboarding, error)
	UpdateProfile(ctx context.Context, arg UpdateProfileParams) (Profile, error)
	UpdateTag(ctx context.Context, arg UpdateTagParams) (Tag, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
