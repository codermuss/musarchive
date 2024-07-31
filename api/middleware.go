package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	localization "github.com/mustafayilmazdev/musarchive/locales"
	"github.com/mustafayilmazdev/musarchive/token"
	"github.com/mustafayilmazdev/musarchive/util"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload_key"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		locale := ctx.Query(util.Locale)
		if len(authorizationHeader) == 0 {
			AbortResponse(ctx, BaseResponse{
				Code: http.StatusUnauthorized,
				Message: ResponseMessage{
					Type:    ERROR,
					Content: localization.GetInstance().Translate(locale, localization.Middleware_HeaderIsNotProvided),
				},
			})

			return
		}
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			AbortResponse(ctx, BaseResponse{
				Code: http.StatusUnauthorized,
				Message: ResponseMessage{
					Type:    ERROR,
					Content: localization.GetInstance().Translate(locale, localization.Middleware_InvalidAuthorization),
				},
			})
			return
		}
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			AbortResponse(ctx, BaseResponse{
				Code: http.StatusUnauthorized,
				Message: ResponseMessage{
					Type:    ERROR,
					Content: localization.GetInstance().Translate(locale, localization.Middleware_UnsupportedAuthorization, authorizationType),
				},
			})
			return
		}
		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		fmt.Println(payload)
		if err != nil {
			AbortResponse(ctx, BaseResponse{
				Code: http.StatusUnauthorized,
				Message: ResponseMessage{
					Type:    ERROR,
					Content: err.Error(),
				},
			})
			return
		}
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
