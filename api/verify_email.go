package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/mustafayilmazdev/musarchive/db/sqlc"
	localization "github.com/mustafayilmazdev/musarchive/locales"
)

type VerifyEmailRequest struct {
	EmailID    int32  `form:"email_id" binding:"required"`
	SecretCode string `form:"secret_code" binding:"required"`
	Locale     string `form:"locale" binding:"required"`
}

func (server *Server) VerifyEmail(ctx *gin.Context) {
	var req VerifyEmailRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		BuildResponse(ctx, BaseResponse{
			Code: http.StatusBadRequest,
			Message: ResponseMessage{
				Type:    ERROR,
				Content: err.Error(),
			},
		})
		return
	}
	_, err := server.store.VerifyEmailTx(ctx, db.VerifyEmailTxParams{
		EmailID:    req.EmailID,
		SecretCode: req.SecretCode,
	})

	if err != nil {
		data := gin.H{
			"Title":        server.lm.Translate(req.Locale, localization.User_VerifyEmailErrorTitle),
			"Timestamp":    time.Now().Format(time.RFC1123),
			"Content":      server.lm.Translate(req.Locale, localization.Errors_AnErrorOccured),
			"ErrorMessage": err,
		}
		ctx.HTML(http.StatusOK, "error_verify_email.html", data)
		return
	}

	data := gin.H{
		"Title":          server.lm.Translate(req.Locale, localization.User_VerifyEmailSuccessTitle),
		"Timestamp":      time.Now().Format(time.RFC1123),
		"Content":        server.lm.Translate(req.Locale, localization.User_VerifyEmailSuccess),
		"WelcomeMessage": server.lm.Translate(req.Locale, localization.User_VerifyEmailMessage),
	}
	ctx.HTML(http.StatusOK, "success_verify_email.html", data)

}
