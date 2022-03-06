package main

import (
	"fmt"
	"gochat/chat-service/adapter"
)

func main() {
	messages := adapter.GetMessages()

	for _, message := range messages {
		fmt.Printf("%v (%v): %v", message.Username, message.CreatedAt, message.Contents)
	}
}
