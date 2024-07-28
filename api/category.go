package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) GetCategories(ctx *gin.Context) {

	categories, err := server.store.GetCategories(ctx)
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
		Data: categories,
	})
}
