package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) GetTags(ctx *gin.Context) {

	tags, err := server.store.GetTags(ctx)
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
		Data: tags,
	})
}
