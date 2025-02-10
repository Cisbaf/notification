package routes

import (
	"whatsappbot/src/wpp"

	"github.com/gin-gonic/gin"
	"go.mau.fi/whatsmeow/types"
)

type NotificationRequest struct {
	To      string `json:"to"`
	Message string `json:"message"`
	IsGroup bool   `json:"is_group"`
}

func NotificationRoute(ctx *gin.Context) {

	var notificationRequest NotificationRequest

	err := ctx.ShouldBindJSON(&notificationRequest)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "cannot bind Json" + err.Error(),
		})
		return
	}

	err = wpp.SendMessage(wpp.Notification{
		Number: func() string {
			if notificationRequest.IsGroup {
				return wpp.GetGroup(notificationRequest.To)
			}
			return notificationRequest.To
		}(),
		Message: notificationRequest.Message,
		Server: func() string {
			if notificationRequest.IsGroup {
				return types.GroupServer
			}
			return types.DefaultUserServer
		}(),
	})

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "erro ao enviar mensagem" + err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "sucesso",
	})
}
