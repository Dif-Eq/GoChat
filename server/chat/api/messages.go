package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gochat/chat-service/adapter"
	"gochat/chat-service/model"

	"github.com/google/uuid"
)

const SUB_NAV = "/messages"

func RegisterMessages() {
	http.HandleFunc(SUB_NAV, func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf(req.Method)
		switch req.Method {
		case "GET":
			getMessages(w, req)
		case "POST":
			createMessage(w, req)
		}
	})
}

func getMessages(w http.ResponseWriter, req *http.Request) {
	messages := adapter.GetMessages()
	j, _ := json.Marshal(messages)
	w.Header().Add("content-type", "application/json")
	w.Write(j)
}

func createMessage(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("content-type", "application/json")
	decoder := json.NewDecoder(req.Body)
	newMessage := &model.Message{}

	err := decoder.Decode(newMessage)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tenantId, err := uuid.Parse(req.Header.Get("tenantId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	userId, err := uuid.Parse(req.Header.Get("userId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	adapter.CreateMessage(tenantId, userId, *newMessage)
	w.Write(nil)
	fmt.Println("Finished processing")
}
