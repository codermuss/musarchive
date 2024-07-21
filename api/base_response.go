package api

import "github.com/gin-gonic/gin"

type MessageType string

const (
	SUCCESS MessageType = "SUCCESS"
	ERROR   MessageType = "ERROR"
	INFO    MessageType = "INFO"
	WARNING MessageType = "WARNING"
)

type ResponseMessage struct {
	Type    MessageType `json:"type"`
	Content string      `json:"content"`
}

type BaseResponse struct {
	Code    int `json:"code"`
	Data    any `json:"data"`
	Message any `json:"message"`
}

func BuildResponse(ctx *gin.Context, response BaseResponse) {
	ctx.JSON(response.Code, response)
}
