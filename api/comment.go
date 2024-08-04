package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/mustafayilmazdev/musarchive/db/sqlc"
	localization "github.com/mustafayilmazdev/musarchive/locales"
	"github.com/mustafayilmazdev/musarchive/token"
	"github.com/mustafayilmazdev/musarchive/util"
)

type CreateCommentRequest struct {
	PostID  int32  `json:"post_id"`
	UserID  int32  `json:"user_id"`
	Content string `json:"content"`
}

func (server *Server) CreateComment(ctx *gin.Context) {
	locale := ctx.Query(util.Locale)
	var req CreateCommentRequest
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

	_, err := server.store.InsertComment(ctx, db.InsertCommentParams{
		PostID:  req.PostID,
		UserID:  req.UserID,
		Content: req.Content,
	})
	if err != nil {
		BuildResponse(ctx, BaseResponse{
			Code: http.StatusBadRequest,
			Message: ResponseMessage{
				Type:    ERROR,
				Content: err.Error(),
			},
		})
		return
	}

	BuildResponse(ctx, BaseResponse{
		Code: http.StatusOK,
		Message: ResponseMessage{
			Type:    SUCCESS,
			Content: server.lm.Translate(locale, localization.Comment_Added),
		},
	})
}

type PostIDRequest struct {
	PostID int32 `json:"post_id"`
}

func (server *Server) GetPostComments(ctx *gin.Context) {
	var req PostIDRequest
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

	postComments, err := server.store.GetCommentsForPost(ctx, req.PostID)
	if err != nil {
		BuildResponse(ctx, BaseResponse{
			Code: http.StatusBadRequest,
			Message: ResponseMessage{
				Type:    ERROR,
				Content: err.Error(),
			},
		})
		return
	}

	BuildResponse(ctx, BaseResponse{
		Code: http.StatusOK,
		Data: postComments,
	})
}

type CommentIDRequest struct {
	CommentID int32 `json:"comment_id"`
}

func (server *Server) DeletePostComment(ctx *gin.Context) {
	locale := ctx.Query(util.Locale)
	var req CommentIDRequest
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

	comment, err := server.store.GetComment(ctx, req.CommentID)
	if err != nil {
		BuildResponse(ctx, BaseResponse{
			Code: http.StatusBadRequest,
			Message: ResponseMessage{
				Type:    ERROR,
				Content: err.Error(),
			},
		})
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if comment.UserID != int32(authPayload.UserID) {
		BuildResponse(ctx, BaseResponse{
			Code: http.StatusBadRequest,
			Message: ResponseMessage{
				Type:    ERROR,
				Content: server.lm.Translate(locale, localization.Comment_CantDeleteOthersComment),
			},
		})
		return
	}

	err = server.store.DeleteComment(ctx, req.CommentID)
	if err != nil {
		BuildResponse(ctx, BaseResponse{
			Code: http.StatusBadRequest,
			Message: ResponseMessage{
				Type:    ERROR,
				Content: err.Error(),
			},
		})
		return
	}

	BuildResponse(ctx, BaseResponse{
		Code: http.StatusOK,
		Message: ResponseMessage{
			Type:    SUCCESS,
			Content: server.lm.Translate(locale, localization.Comment_Deleted),
		},
	})
}
