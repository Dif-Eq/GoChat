package main

import (
	"fmt"
	"gochat/chat-service/api"
	"net/http"
)

func main() {
	api.RegisterMessages()
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
