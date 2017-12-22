package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

)

const (
	BOT_TOKEN = "264642626:AAH55-PhATXCMRRAORcf4Cvs2zPc2ltwi6E"
	MyURL = "https://nameless-journey-85806.herokuapp.com/"
)
var (
	BOT_URL = fmt.Sprintf("https://api.telegram.org/bot%s/", BOT_TOKEN)
)

func main() {
	PORT := os.Getenv("PORT")

	if PORT == "" {
		log.Fatal("$PORT must be set")
	}

	resp, err := http.Get(BOT_URL + "setWebhook?url=" + MyURL + "tmbgot")
	print(resp.StatusCode)

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	if err != nil {
		log.Fatal(fmt.Sprintf("Can't set hook: %s. Quit.", bodyString))
	}
	resp.Body.Close()
	print(bodyBytes)

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("./templates/*.tmpl.html")
	//router.Static("/static", "static")

	SetupUrls(router)
	router.Run(":" + PORT)
}
