package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	_"net/http"
	"os"

	m"github.com/stn117/tmBGOt/models"
	_"github.com/gin-gonic/gin"

)

const (
	BOT_TOKEN = "264642626:AAH55-PhATXCMRRAORcf4Cvs2zPc2ltwi6E"
	MyURL = "https://nameless-journey-85806.herokuapp.com/"
)
var (
	BOT_URL = fmt.Sprintf("https://api.telegram.org/bot%s/", BOT_TOKEN)
)



func update(w http.ResponseWriter, r *http.Request) {

	message := &m.ReceiveMessage{}

	chatID := 0
	msgText := ""

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		fmt.Println(err)
	}



	// if private or group
	if message.Message.Chat.ID != 0 {
		fmt.Println(message.Message.Chat.ID, message.Message.Text)
		chatID = message.Message.Chat.ID
		msgText = message.Message.Text
	} else {
		// if channel
		fmt.Println(message.ChannelPost.Chat.ID, message.ChannelPost.Text)
		chatID = message.ChannelPost.Chat.ID
		msgText = message.ChannelPost.Text
	}

	respMsg := fmt.Sprintf(
		"%s%s/sendMessage?chat_id=%d&text=Received: %s",
		BOT_URL, BOT_TOKEN, chatID, msgText)

	// send echo resp
	_, err = http.Get(respMsg)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	PORT := os.Getenv("PORT")

	http.HandleFunc("/api/v1/update", update)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	fmt.Println("Listenning on port", PORT, ".")
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Fatal(err)
	}
}


//func main() {
//	port := os.Getenv("PORT")
//
//	if port == "" {
//		log.Fatal("$PORT must be set")
//	}
//
//	router := gin.New()
//	router.Use(gin.Logger())
//	router.LoadHTMLGlob("./templates/*.tmpl.html")
//	//router.Static("/static", "static")
//
//	SetupUrls(router)
//	router
//
//	router.Run(":" + port)
//}