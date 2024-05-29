package handler

import (
	utils "hello/Utils"
	"log"

	"github.com/gin-gonic/gin"
)

func HandleTelegramWebHook(g *gin.Context) {

	// Parse incoming request
	var update, err = utils.ParseTelegramRequest(g.Request)
	if err != nil {
		log.Printf("error parsing update, %s", err.Error())
		return
	}

	var telegramResponseBody, errTelegram = utils.SendTextToTelegramChat(update.Message.Chat.Id, "Tempo Message")
	if errTelegram != nil {
		log.Printf("got error %s from telegram, reponse body is %s", errTelegram.Error(), telegramResponseBody)
	}
}
