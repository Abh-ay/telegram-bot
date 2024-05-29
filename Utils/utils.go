package utils

import (
	"encoding/json"
	"fmt"
	models "hello/Models"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func ParseTelegramRequest(r *http.Request) (*models.Update, error) {
	var update models.Update
	fmt.Println(r.Body)
	fmt.Println("))))))))))))))*************")
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("could not decode incoming update %s", err.Error())
		return nil, err
	}
	fmt.Println("=================")
	fmt.Println(update.Message.Chat.Id)
	return &update, nil
}

func SendTextToTelegramChat(chatId int, text string) (string, error) {

	log.Printf("Sending %s to chat_id: %d", text, chatId)
	var telegramApi string = "https://api.telegram.org/bot" + os.Getenv("TELEGRAM_BOT_TOKEN") + "/sendMessage"
	response, err := http.PostForm(
		telegramApi,
		url.Values{
			"chat_id": {strconv.Itoa(chatId)},
			"text":    {text},
		})

	if err != nil {
		log.Printf("error when posting text to the chat: %s", err.Error())
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = io.ReadAll(response.Body)
	if errRead != nil {
		log.Printf("error in parsing telegram answer %s", errRead.Error())
		return "", err
	}
	bodyString := string(bodyBytes)
	log.Printf("Body of Telegram Response: %s", bodyString)

	return bodyString, nil
}