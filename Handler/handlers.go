package handler

import (
	"encoding/json"
	"fmt"
	models "hello/Models"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func HandleTelegramWebHook(g *gin.Context) {
	update := models.Update{}

	var url = "https://api.telegram.org/bot" + os.Getenv("TELEGRAM_APITOKEN") + "/getUpdates"
	var res, _ = http.Get(url)
	fmt.Println(url)
	fmt.Println("!11!")
	defer res.Body.Close()
	fmt.Println(res.Body)
	var resp = json.NewDecoder(res.Body).Decode(&update)
	fmt.Println("!22!")
	fmt.Println(resp)
	body, _ := io.ReadAll(res.Body)
	err := json.Unmarshal(body, &update)
	if err != nil {
		g.ShouldBindJSON(http.StatusInternalServerError)
		return
	}
	g.JSON(http.StatusOK, update)

}
