package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type RegisterUserTxParams struct {
	InsertUserParams
	// * Note [codermuss]: This function will be executed after the user is inserted,
	// * Note [codermuss]: inside the same transaction. And its output error will be used to decide
	// * Note [codermuss]: whether to commit or rollback the transaction.
	AfterCreate func(user User) error
}
type RegisterUserTxResult struct {
	User    User
	Profile Profile
}

// * Note [codermuss]: This method responsible with creating user.
// * Note [codermuss]: It uses execTx to handle DB Transaction error
func (store *SQLStore) RegisterUserTx(ctx context.Context, arg RegisterUserTxParams) (RegisterUserTxResult, error) {
	var result RegisterUserTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.User, err = q.InsertUser(ctx, arg.InsertUserParams)
		if err != nil {
			return err
		}
		profileArg := InsertProfileParams{
			UserID: result.User.ID,
			Bio:    pgtype.Text{Valid: true, String: ""},
			PostCount: pgtype.Int4{
				Valid: true, Int32: 0,
			},
			LikeCount: pgtype.Int4{
				Valid: true, Int32: 0,
			},
			FollowerCount: pgtype.Int4{
				Valid: true, Int32: 0,
			},
		}

		result.Profile, err = q.InsertProfile(ctx, profileArg)
		if err != nil {
			return err
		}

		return arg.AfterCreate(result.User)
	})
	return result, err
}
