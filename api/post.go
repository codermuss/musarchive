package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/mustafayilmazdev/musarchive/db/sqlc"
	localization "github.com/mustafayilmazdev/musarchive/locales"
	"github.com/mustafayilmazdev/musarchive/token"
	"github.com/mustafayilmazdev/musarchive/util"
)

type createPostRequest struct {
	Title      string      `json:"title" binding:"required"`
	Content    string      `json:"content" binding:"required"`
	CoverImage pgtype.Text `json:"cover_image"`
	Categories *[]int      `json:"categories"`
	Tags       *[]int      `json:"tags"`
}

func (server *Server) CreatePost(ctx *gin.Context) {
	locale := ctx.Query(util.Locale)
	var req createPostRequest
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
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.InsertPostParams{
		UserID: pgtype.Int4{
			Valid: true,
			Int32: int32(authPayload.UserID),
		},
		Title:      req.Title,
		Content:    req.Content,
		CoverImage: req.CoverImage,
	}

	post, err := server.store.InsertPost(ctx, arg)
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

	if req.Categories != nil {
		for _, categoryID := range *req.Categories {
			postCategoryArg := db.InsertPostCategoryParams{
				PostID:     post.ID,
				CategoryID: int32(categoryID),
			}
			_, err = server.store.InsertPostCategory(ctx, postCategoryArg)
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
		}
	}

	if req.Tags != nil {
		for _, v := range *req.Tags {
			postTag := db.InsertPostTagParams{
				PostID: post.ID,
				TagID:  int32(v),
			}
			_, err = server.store.InsertPostTag(ctx, postTag)

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
		}
	}

	BuildResponse(ctx, BaseResponse{
		Code: http.StatusOK,
		Data: post,
		Message: ResponseMessage{
			Type:    SUCCESS,
			Content: server.lm.Translate(locale, localization.Post_InsertSuccess),
		},
	})
}

type filteredPostsRequest struct {
	Categories []int32 `json:"categories"`
	Tags       []int32 `json:"tags"`
}

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
	var req filteredPostsRequest
	err := ctx.ShouldBindJSON(&req)
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

	arg := db.GetPostsWithFilterParams{
		Column3: req.Categories,
		Column4: req.Tags,
		Limit:   int32(size),
		Offset:  int32((page - 1) * size),
	}

	blogs, err := server.store.GetPostsWithFilter(ctx, arg)
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
