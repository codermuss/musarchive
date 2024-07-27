package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/mustafayilmazdev/musarchive/db/sqlc"
	localization "github.com/mustafayilmazdev/musarchive/locales"
	"github.com/mustafayilmazdev/musarchive/util"
	"github.com/rs/zerolog/log"
)

func (server *Server) GetBlogs(ctx *gin.Context) {
	localeValue := ctx.Query(util.Locale)
	pageStr := ctx.Query(util.Page)
	sizeStr := ctx.Query(util.Size)
	page := util.PageCount
	size := util.SizeCount
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

	arg := db.GetBlogsParams{
		Limit:  int32(size),
		Offset: int32((page - 1) * size),
	}

	blogs, err := server.store.GetBlogs(ctx, arg)
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
