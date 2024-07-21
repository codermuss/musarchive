package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) GetOnboardings(ctx *gin.Context) {

	onboardings, err := server.store.ListOnboarding(ctx)
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
		Data: onboardings,
	})
}
