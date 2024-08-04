package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type VerifyEmailTxParams struct {
	EmailID    int32
	SecretCode string
}
type VerifyEmailTxResult struct {
	User        User
	VerifyEmail VerifyEmail
}

// * Note [codermuss]: This method responsible with creating user.
// * Note [codermuss]: It uses execTx to handle DB Transaction error
func (store *SQLStore) VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error) {
	var result VerifyEmailTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.VerifyEmail, err = q.UpdateVerifyEmail(ctx, UpdateVerifyEmailParams{
			ID:         arg.EmailID,
			SecretCode: arg.SecretCode,
		})
		if err != nil {
			return err
		}
		user, err := q.GetUser(ctx, result.VerifyEmail.Username)
		if err != nil {
			return err
		}
		result.User, err = q.UpdateUser(ctx, UpdateUserParams{
			ID: user.ID,
			IsEmailVerified: pgtype.Bool{
				Valid: true,
				Bool:  true,
			},
		})
		return err
	})
	return result, err
}
