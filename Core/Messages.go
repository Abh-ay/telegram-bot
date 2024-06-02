package core

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	enums "hello/Enums"
	models "hello/Models"
	"net/http"
	"os"
	"strings"
)

type Core struct {
	Quries *models.Queries
}

func (c *Core) SetQueries(query *models.Queries) {
	fmt.Println("Set Quries method")
	fmt.Println(query)
	c.Quries = query
}

func (r *Core) SendMessages(result models.Result) (resp *http.Response, err error) {
	var refQuery models.DtoRefQuery
	msgText := ""
	err = r.Quries.GetQuery.Get(&refQuery, strings.ToLower(result.Message.Text))
	fmt.Println(refQuery.ID)
	fmt.Println(refQuery.ExpectedMessage)
	msgText = refQuery.ExpectedMessage
	fmt.Println(msgText)
	if err != nil {
		if err == sql.ErrNoRows {
			msgText = "Hmm, It's look like you sent wrong input"
		} else {
			panic("Got error while get query from DB")
		}
	}
	if refQuery.ID != 0 {
		fmt.Println(refQuery.ID)
		fmt.Println(refQuery.ExpectedMessage)
		msgText = refQuery.ExpectedMessage
		fmt.Println(msgText)
	}

	postBody, _ := json.Marshal(map[string]string{
		"chat_id": fmt.Sprint(result.Message.Chat.ID),
		"text":    msgText,
	})
	return http.Post(enums.Tele_Url+os.Getenv("TELEGRAM_APITOKEN")+"/sendMessage", "application/json", bytes.NewBuffer(postBody))

}
