package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) GetOnboardings(ctx *gin.Context) {

	onboardings, err := server.store.ListOnboarding(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, onboardings)
}
