package routes

import (
	"whatsappbot/src/wpp"

	"github.com/gin-gonic/gin"
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

	ntf := wpp.MakeNotification(
		notificationRequest.To,
		notificationRequest.Message,
		notificationRequest.IsGroup,
	)
	err = wpp.SendMessage(ntf)

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
