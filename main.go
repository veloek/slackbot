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

	conn := slack.Connection{}
	err := conn.Init(os.Args[1])
	if err != nil {
		fmt.Printf("Error while connecting to slack: %s\n", err)
	} else {
		fmt.Println("Connected to slack")

		for {
			msg := conn.GetMessage()
			fmt.Printf("Received: %s\n", msg)
		}
		conn.Close()
	}
}
