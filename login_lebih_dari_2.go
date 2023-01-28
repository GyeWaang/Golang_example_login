package main

import (
	"fmt"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Account struct that contains the access token and secret for each account
type Account struct {
	AccessToken string
	Secret      string
}

var accounts = []Account{
	{AccessToken: "ACCESS_TOKEN_1", Secret: "SECRET_1"},
	{AccessToken: "ACCESS_TOKEN_2", Secret: "SECRET_2"},
}

func main() {
	http.HandleFunc("/callback", callbackHandler)
	http.ListenAndServe(":8080", nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	// Get the account info based on the user id
	account := getAccount(r.Header.Get("X-Line-User-Id"))
	// create a new line bot with the account info
	bot, err := linebot.New(account.Secret, account.AccessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Handle events here
	events, err := bot.ParseRequest(r)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

func getAccount(userID string) Account {
	// Get the account based on the user id
	// You can use a database or other method to store and retrieve the account info
	// In this example, the account info is hardcoded
	return accounts[0]
}
