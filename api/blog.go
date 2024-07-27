package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/mustafayilmazdev/musarchive/db/sqlc"
	localization "github.com/mustafayilmazdev/musarchive/locales"
	"github.com/rs/zerolog/log"
)

func (server *Server) GetBlogs(ctx *gin.Context) {
	localeValue := ctx.Query("locale")
	pageStr := ctx.Query("page")
	sizeStr := ctx.Query("size")
	page := 1
	size := 10
	log.Info().Msg(pageStr + " " + sizeStr)
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
		}
	}

	arg := db.GetBlogsParams{
		Limit:  int32(size),
		Offset: int32((page - 1) * size),
	}

	blogs, err := server.store.GetBlogs(ctx, arg)
	log.Info().Msgf("blogs: %v", len(blogs))
	if err != nil {
		BuildResponse(ctx, BaseResponse{
			Code: http.StatusInternalServerError,
			Message: ResponseMessage{
				Type:    ERROR,
				Content: "Internal error: " + err.Error(),
			},
		})
		return
	}

	BuildResponse(ctx, BaseResponse{
		Code: http.StatusOK,
		Data: blogs,
	})
}
