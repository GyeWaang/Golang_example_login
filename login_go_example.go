package main

import (
	"fmt"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	// Your line bot configuration
	config := &linebot.ClientConfig{
		ChannelSecret: "YOUR_CHANNEL_SECRET",
		ChannelToken:  "YOUR_CHANNEL_TOKEN",
	}

	// Initialize the line bot client
	client, err := linebot.NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set up a handler function for the webhook
	http.HandleFunc("/gye/webhook", func(w http.ResponseWriter, r *http.Request) {
		// Parse the incoming request
		events, err := client.ParseRequest(r)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}

		// Handle each event
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if message.Text == "login" {
						// Perform login logic here
						// ...

						// Send a message to the user
						if _, err := client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Login successful!")).Do(); err != nil {
							fmt.Println(err)
						}
					}
				}
			}
		}
	})

	// Start the server
	http.ListenAndServe(":8080", nil)
}
