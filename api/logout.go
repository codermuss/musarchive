package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	db "github.com/mustafayilmazdev/musarchive/db/sqlc"
	localization "github.com/mustafayilmazdev/musarchive/locales"
	"github.com/mustafayilmazdev/musarchive/token"
	"github.com/mustafayilmazdev/musarchive/util"
)

func (server *Server) LogoutUser(ctx *gin.Context) {
	locale := ctx.Query(util.Locale)
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	getSession, err := server.store.GetSession(ctx, authPayload.ID)
	if err != nil {
		fmt.Println("Session not found")
	}
	fmt.Println(getSession)
	if getSession.IsBlocked {

		BuildResponse(ctx, BaseResponse{
			Code: http.StatusUnauthorized,
			Message: ResponseMessage{
				Type:    ERROR,
				Content: "Session expired",
			},
		})
		return
	}
	err = server.store.UpdateSession(ctx, db.UpdateSessionParams{ID: authPayload.ID, IsBlocked: true})
	fmt.Println(authPayload.ID)
	if err != nil {
		BuildResponse(ctx, BaseResponse{
			Code: http.StatusInternalServerError,
			Message: ResponseMessage{
				Type:    ERROR,
				Content: err.Error(),
			},
		})
	}
	BuildResponse(ctx, BaseResponse{
		Code:    http.StatusInternalServerError,
		Message: server.lm.Translate(locale, localization.User_LogoutSuccess),
	})
}
