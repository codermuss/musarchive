package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	db "github.com/mustafayilmazdev/musarchive/db/sqlc"
	"github.com/mustafayilmazdev/musarchive/util"
)

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	SessionID             uuid.UUID    `json:"session_id"`
	AccessToken           string       `json:"access_token"`
	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
	RefreshToken          string       `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
	User                  UserResponse `json:"user"`
}

func newUserResponse(user db.User) UserResponse {
	return UserResponse{
		ID:                user.ID,
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		Avatar:            user.Avatar,
		BirthDate:         user.BirthDate,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

func (server *Server) LoginUser(ctx *gin.Context) {
	var req loginUserRequest
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

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			BuildResponse(ctx, BaseResponse{
				Code: http.StatusNotFound,
				Message: ResponseMessage{
					Type:    ERROR,
					Content: err.Error(),
				},
			})
			return
		}
		BuildResponse(ctx, BaseResponse{
			Code: http.StatusInternalServerError,
			Message: ResponseMessage{
				Type:    ERROR,
				Content: err.Error(),
			},
		})
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		BuildResponse(ctx, BaseResponse{
			Code: http.StatusUnauthorized,
			Message: ResponseMessage{
				Type:    ERROR,
				Content: err.Error(),
			},
		})
		return
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(int(user.ID), server.config.AccessTokenDuration)
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

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		int(user.ID),
		server.config.RefreshTokenDuration,
	)
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

	session, err := server.store.InsertSession(ctx, db.InsertSessionParams{
		ID:           refreshPayload.ID,
		UserID:       user.ID,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})

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

	rsp := loginUserResponse{
		SessionID:             session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  newUserResponse(user),
	}

	BuildResponse(ctx, BaseResponse{
		Code: http.StatusOK,
		Data: rsp,
	})
}
