package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"

	db "github.com/mustafayilmazdev/musarchive/db/sqlc"
	localization "github.com/mustafayilmazdev/musarchive/locales"
	"github.com/mustafayilmazdev/musarchive/util"
)

type registerUserRequest struct {
	Username  string      `json:"username" binding:"required,alphanum"`
	Password  string      `json:"password" binding:"required,min=6"`
	FullName  string      `json:"full_name" binding:"required"`
	Email     string      `json:"email" binding:"required,email"`
	Avatar    pgtype.Text `json:"avatar"`
	BirthDate pgtype.Date `json:"birth_date" binding:"required"`
}

type registerUserResponse struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`

	FullName          string      `json:"full_name"`
	Email             string      `json:"email"`
	Avatar            pgtype.Text `json:"avatar"`
	BirthDate         pgtype.Date `json:"birth_date"`
	PasswordChangedAt time.Time   `json:"password_changed_at"`
	CreatedAt         time.Time   `json:"created_at"`
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

	arg := db.InsertUserParams{
		Username:  req.Username,
		Password:  hashedPassword,
		FullName:  req.FullName,
		Email:     req.Email,
		Avatar:    req.Avatar,
		BirthDate: req.BirthDate,
	}

	user, err := server.store.InsertUser(ctx, arg)

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

	responseUser := registerUserResponse{
		ID:                user.ID,
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		Avatar:            user.Avatar,
		BirthDate:         user.BirthDate,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}

	BuildResponse(ctx, BaseResponse{
		Code: http.StatusOK,
		Data: responseUser,
		Message: ResponseMessage{
			Type:    SUCCESS,
			Content: server.lm.Translate(localeValue, localization.User_RegisterSuccess),
		},
	})
}
