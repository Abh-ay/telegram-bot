package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	enums "hello/Enums"
	models "hello/Models"
	"net/http"
	"os"
	"strings"
)

func SendMessages(result models.Result) (resp *http.Response, err error) {
	msgText := ""
	if strings.ToLower(result.Message.Text) == "hi" {
		msgText = "Welcome to Teetbot... \n How can I help you 🙌🏻?"
	} else {
		msgText = "You sent other than hi......."
	}
	postBody, _ := json.Marshal(map[string]string{
		"chat_id": fmt.Sprint(result.Message.Chat.ID),
		"text":    msgText,
	})
	return http.Post(enums.Tele_Url+os.Getenv("TELEGRAM_APITOKEN")+"/sendMessage", "application/json", bytes.NewBuffer(postBody))

}
