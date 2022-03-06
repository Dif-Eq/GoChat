package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gochat/chat-service/adapter"
)

const SUB_NAV = "/messages"

func RegisterMessages() {
	http.HandleFunc(SUB_NAV, getMessages)
}

func getMessages(w http.ResponseWriter, req *http.Request) {
	fmt.Println("handling request")
	messages := adapter.GetMessages()
	j, _ := json.Marshal(messages)
	w.Write(j)
}
