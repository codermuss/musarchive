package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgtype"

	db "github.com/mustafayilmazdev/musarchive/db/sqlc"
	localization "github.com/mustafayilmazdev/musarchive/locales"
	"github.com/mustafayilmazdev/musarchive/util"
	"github.com/mustafayilmazdev/musarchive/worker"
)

type registerUserRequest struct {
	Username  string      `json:"username" binding:"required,alphanum"`
	Password  string      `json:"password" binding:"required,min=6"`
	FullName  string      `json:"full_name" binding:"required"`
	Email     string      `json:"email" binding:"required,email"`
	Avatar    pgtype.Text `json:"avatar"`
	BirthDate pgtype.Date `json:"birth_date" binding:"required"`
}

type UserResponse struct {
	ID                int32       `json:"id"`
	Username          string      `json:"username"`
	FullName          string      `json:"full_name"`
	Email             string      `json:"email"`
	Role              string      `json:"role"`
	Avatar            pgtype.Text `json:"avatar"`
	BirthDate         pgtype.Date `json:"birth_date"`
	PasswordChangedAt time.Time   `json:"password_changed_at"`
	CreatedAt         time.Time   `json:"created_at"`
}

type RegisterResponse struct {
	User    UserResponse `json:"user"`
	Profile db.Profile   `json:"profile"`
}

func (server *Server) RegisterUser(ctx *gin.Context) {

	localeValue := ctx.Query("locale")
	var req registerUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		BuildResponse(ctx, BaseResponse{
			Code: http.StatusBadRequest,
			Message: ResponseMessage{
				Type:    ERROR,
				Content: err.Error(),
			},
		})
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		BuildResponse(ctx, BaseResponse{
			Code: http.StatusInternalServerError,
			Message: ResponseMessage{
				Type:    ERROR,
				Content: err.Error(),
			},
		})
		return
	}

	argUser := db.InsertUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
		Role:           util.Standard,
		Avatar:         req.Avatar,
		BirthDate:      req.BirthDate,
	}
	argAfterCreate := func(user db.User) error {
		taskPayload := &worker.PayloadSendVerifyEmail{Username: user.Username}
		opts := []asynq.Option{
			asynq.MaxRetry(10),
			asynq.ProcessIn(10 * time.Second),
			asynq.Queue(worker.QueueCritical),
		}
		return server.taskDistributor.DistributeTaskSendVerifyEmail(ctx, taskPayload, opts...)
	}

	arg := db.RegisterUserTxParams{
		InsertUserParams: argUser,
		AfterCreate:      argAfterCreate,
	}

	userAndProfile, err := server.store.RegisterUserTx(ctx, arg)

	if err != nil {

		if db.ErrorCode(err) == db.UniqueViolation {
			BuildResponse(ctx, BaseResponse{
				Code: http.StatusForbidden,
				Message: ResponseMessage{
					Type:    ERROR,
					Content: server.lm.Translate(localeValue, localization.Errors_InternalError, err),
				},
			})
			return
		}

		BuildResponse(ctx, BaseResponse{
			Code: http.StatusInternalServerError,
			Message: ResponseMessage{
				Type:    ERROR,
				Content: server.lm.Translate(localeValue, localization.Errors_InternalError, err),
			},
		})
		return
	}

	registerResponse := RegisterResponse{
		User: UserResponse{
			ID:                userAndProfile.User.ID,
			Username:          userAndProfile.User.Username,
			FullName:          userAndProfile.User.FullName,
			Email:             userAndProfile.User.Email,
			Role:              userAndProfile.User.Role,
			Avatar:            userAndProfile.User.Avatar,
			BirthDate:         userAndProfile.User.BirthDate,
			PasswordChangedAt: userAndProfile.User.PasswordChangedAt,
			CreatedAt:         userAndProfile.User.CreatedAt,
		},
		Profile: userAndProfile.Profile,
	}

	BuildResponse(ctx, BaseResponse{
		Code: http.StatusOK,
		Data: registerResponse,
		Message: ResponseMessage{
			Type:    SUCCESS,
			Content: server.lm.Translate(localeValue, localization.User_RegisterSuccess),
		},
	})
}
