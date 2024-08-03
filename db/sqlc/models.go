// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Category struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	ID        int32              `json:"id"`
	PostID    int32              `json:"post_id"`
	UserID    int32              `json:"user_id"`
	Content   string             `json:"content"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

type FeaturedStory struct {
	ID           int32       `json:"id"`
	PostID       int32       `json:"post_id"`
	FeaturedDate pgtype.Date `json:"featured_date"`
}

type Onboarding struct {
	ID          int32       `json:"id"`
	Image       pgtype.Text `json:"image"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
}

type Post struct {
	ID         int32       `json:"id"`
	UserID     pgtype.Int4 `json:"user_id"`
	Title      string      `json:"title"`
	Content    string      `json:"content"`
	CoverImage pgtype.Text `json:"cover_image"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	Likes      int32       `json:"likes"`
}

type PostCategory struct {
	PostID     int32 `json:"post_id"`
	CategoryID int32 `json:"category_id"`
}

type PostLike struct {
	PostID int32 `json:"post_id"`
	UserID int32 `json:"user_id"`
}

type PostTag struct {
	PostID int32 `json:"post_id"`
	TagID  int32 `json:"tag_id"`
}

type Profile struct {
	UserID        int32       `json:"user_id"`
	Bio           pgtype.Text `json:"bio"`
	PostCount     pgtype.Int4 `json:"post_count"`
	LikeCount     pgtype.Int4 `json:"like_count"`
	FollowerCount pgtype.Int4 `json:"follower_count"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	UserID       int32     `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type Tag struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID                int32       `json:"id"`
	Username          string      `json:"username"`
	HashedPassword    string      `json:"hashed_password"`
	FullName          string      `json:"full_name"`
	Email             string      `json:"email"`
	Avatar            pgtype.Text `json:"avatar"`
	Role              string      `json:"role"`
	BirthDate         pgtype.Date `json:"birth_date"`
	IsEmailVerified   bool        `json:"is_email_verified"`
	PasswordChangedAt time.Time   `json:"password_changed_at"`
	CreatedAt         time.Time   `json:"created_at"`
}

type UserFollower struct {
	UserID     int32 `json:"user_id"`
	FollowerID int32 `json:"follower_id"`
}

type UserPost struct {
	UserID int32 `json:"user_id"`
	PostID int32 `json:"post_id"`
}

type VerifyEmail struct {
	ID         int32     `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	SecretCode string    `json:"secret_code"`
	IsUsed     bool      `json:"is_used"`
	CreatedAt  time.Time `json:"created_at"`
	ExpiredAt  time.Time `json:"expired_at"`
}
