package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	localization "github.com/mustafayilmazdev/musarchive/locales"
)

func (server *Server) GetTags(ctx *gin.Context) {

	tags, err := server.store.GetTags(ctx)
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

	BuildResponse(ctx, BaseResponse{
		Code: http.StatusOK,
		Data: tags,
	})
}

type InsertTagRequest struct {
	Name string `json:"name"`
}

func (server *Server) InsertTag(ctx *gin.Context) {
	locale, _ := ctx.Get(localeKey)

	var req InsertTagRequest
	err := ctx.ShouldBindBodyWithJSON(&req)
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
	tag, err := server.store.InsertTag(ctx, req.Name)
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
	BuildResponse(ctx, BaseResponse{
		Code:    http.StatusOK,
		Message: server.lm.Translate(locale.(string), localization.Tag_Inserted, tag.Name),
	})
}
