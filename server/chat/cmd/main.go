package main

import (
	"fmt"
	"gochat/chat-service/model"
)

func main() {
	message := model.Message{}
	message.Contents = "this is a fucking message"

	fmt.Printf(message.Contents)
}
