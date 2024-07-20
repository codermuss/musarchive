// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
)

type Querier interface {
	DeleteOnboarding(ctx context.Context, id int32) error
	DeleteUser(ctx context.Context, id int32) error
	GetOnboarding(ctx context.Context, id int32) (Onboarding, error)
	GetUser(ctx context.Context, id int32) (User, error)
	InsertOnboarding(ctx context.Context, arg InsertOnboardingParams) (Onboarding, error)
	InsertUser(ctx context.Context, arg InsertUserParams) (User, error)
	UpdateOnboarding(ctx context.Context, arg UpdateOnboardingParams) (Onboarding, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
