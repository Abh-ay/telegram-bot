package handler

import (
	"encoding/json"
	"fmt"
	core "hello/Core"
	enums "hello/Enums"
	models "hello/Models"
	utils "hello/Utils"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var CacheStruct = utils.Cache{}
var C core.Core

func GetUpdates(g *gin.Context) {
	update := models.Update{}
	offsetValue, isOffsetPresent := CacheStruct.Get(enums.LatestUpdateId)
	if isOffsetPresent {
		var url = enums.Tele_Url + os.Getenv("TELEGRAM_APITOKEN") + "/getUpdates?offset=" + fmt.Sprint(offsetValue)
		fmt.Println(url)
		FetchResponse(url, g, &update)
	} else {
		var url = enums.Tele_Url + os.Getenv("TELEGRAM_APITOKEN") + "/getUpdates"
		FetchResponse(url, g, &update)
	}
	if utils.IsNil(update.Result) {
		g.JSON(http.StatusOK, "No Updates hase been found till yet")
		return
	}
	count := (len(update.Result) - 1)
	if count < 1 {
		count = 0
		CacheStruct.Set(enums.LatestUpdateId, update.Result[count].UpdateID)
	} else {
		CacheStruct.Set(enums.LatestUpdateId, update.Result[count].UpdateID)
	}
	C.SendMessages(update.Result[0])
}

func FetchResponse(url string, g *gin.Context, update *models.Update) {
	var res, _ = http.Get(url)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	err := json.Unmarshal(body, &update)
	if err != nil {
		g.ShouldBindJSON(http.StatusInternalServerError)
		return
	}
}

func WebHookHandler(g *gin.Context) {
	var result models.Result
	body, err := io.ReadAll(g.Request.Body)
	if err != nil {
		g.ShouldBindJSON(http.StatusMisdirectedRequest)
		return
	}
	if err := json.Unmarshal(body, &result); err != nil {
		g.ShouldBindJSON("Error in unmarshal Json")
		return
	}
	if utils.IsNil(result) {
		g.JSON(http.StatusOK, "No Updates hase been found till yet")
		return
	}
	if _, err := C.SendMessages(result); err != nil {
		g.ShouldBindJSON("Error while sending messsage")
		return
	}
}
