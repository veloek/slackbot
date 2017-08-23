package main

import (
	"fmt"
	"github.com/veloek/slackbot/slack"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <slack-bot-token>\n", os.Args[0])
		os.Exit(1)
	}

	conn := slack.Connect(os.Args[1])

	fmt.Println("Connected to slack")
	defer conn.Dispose()

	for msg := range conn.Messages {
		fmt.Printf("Received: %s\n", msg.Text)
		go handleMessage(&msg)
	}
}

func handleMessage(m *slack.Message) {
	m.Respond("Hei på deg!")
}
