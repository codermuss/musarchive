package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/mustafayilmazdev/musarchive/db/sqlc"
	localization "github.com/mustafayilmazdev/musarchive/locales"
	"github.com/mustafayilmazdev/musarchive/token"
	"github.com/mustafayilmazdev/musarchive/util"
)

func (server *Server) GetPosts(ctx *gin.Context) {
	localeValue := ctx.Query(util.Locale)
	pageStr := ctx.Query(util.Page)
	sizeStr := ctx.Query(util.Size)
	page := util.PageCount
	size := util.SizeCount

	if pageStr != "" {
		var err error
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			BuildResponse(ctx, BaseResponse{
				Code: http.StatusBadRequest,
				Message: ResponseMessage{
					Type:    ERROR,
					Content: server.lm.Translate(localeValue, localization.Pagination_pageError, pageStr),
				},
			})
			return
		}

	}

	if sizeStr != "" {
		var err error
		size, err = strconv.Atoi(sizeStr)
		if err != nil {
			BuildResponse(ctx, BaseResponse{
				Code: http.StatusBadRequest,
				Message: ResponseMessage{
					Type:    ERROR,
					Content: server.lm.Translate(localeValue, localization.Pagination_sizeError, sizeStr),
				},
			})
			return
		}

	}

	arg := db.GetPostsParams{
		Limit:  int32(size),
		Offset: int32((page - 1) * size),
	}

	blogs, err := server.store.GetPosts(ctx, arg)
	if err != nil {
		BuildResponse(ctx, BaseResponse{
			Code: http.StatusInternalServerError,
			Message: ResponseMessage{
				Type:    ERROR,
				Content: server.lm.Translate(localeValue, localization.Errors_InternalError, err.Error()),
			},
		})
		return
	}

	BuildResponse(ctx, BaseResponse{
		Code: http.StatusOK,
		Data: blogs,
	})
}

func (server *Server) GetFollowedPosts(ctx *gin.Context) {
	localeValue := ctx.Query(util.Locale)
	pageStr := ctx.Query(util.Page)
	sizeStr := ctx.Query(util.Size)
	page := util.PageCount
	size := util.SizeCount

	if pageStr != "" {
		var err error
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			BuildResponse(ctx, BaseResponse{
				Code: http.StatusBadRequest,
				Message: ResponseMessage{
					Type:    ERROR,
					Content: server.lm.Translate(localeValue, localization.Pagination_pageError, pageStr),
				},
			})
			return
		}

	}

	if sizeStr != "" {
		var err error
		size, err = strconv.Atoi(sizeStr)
		if err != nil {
			BuildResponse(ctx, BaseResponse{
				Code: http.StatusBadRequest,
				Message: ResponseMessage{
					Type:    ERROR,
					Content: server.lm.Translate(localeValue, localization.Pagination_sizeError, sizeStr),
				},
			})
			return
		}

	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	fmt.Println(authPayload.UserID)
	arg := db.GetFollowedPostsParams{
		FollowerID: int32(authPayload.UserID),
		Limit:      int32(size),
		Offset:     int32((page - 1) * size),
	}

	blogs, err := server.store.GetFollowedPosts(ctx, arg)
	if err != nil {
		BuildResponse(ctx, BaseResponse{
			Code: http.StatusInternalServerError,
			Message: ResponseMessage{
				Type:    ERROR,
				Content: server.lm.Translate(localeValue, localization.Errors_InternalError, err.Error()),
			},
		})
		return
	}

	BuildResponse(ctx, BaseResponse{
		Code: http.StatusOK,
		Data: blogs,
	})
}
