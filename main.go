package main

import (
	"fmt"
	goTel "github.com/yoruba-codigy/goTelegram"
	"log"
	"net/http"
	"os"
)

var db = initDatabase()
var userList map[int]*userPendingData
var reportList map[int]*reportPendingData
var bot, err = goTel.NewBot(os.Getenv("token"))

func main() {
	userList = make(map[int]*userPendingData)
	reportList = make(map[int]*reportPendingData)

	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Printf("Bot Name: %s\nBot Username: %s\n", bot.Me.Firstname, bot.Me.Username)

	bot.SetHandler(handler)

	log.Println("Starting Server")
	err = http.ListenAndServe(":"+os.Getenv("PORT"), http.HandlerFunc(bot.UpdateHandler))

	if err != nil {
		log.Println("Failed")
		log.Fatalln(err)
		return
	}
}
