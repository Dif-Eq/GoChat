package main

import (
	"fmt"
	"gochat/chat-service/adapter"
)

func main() {
	messages := adapter.GetMessages()

	for _, message := range messages {
		fmt.Printf(message.Contents)
	}
}
