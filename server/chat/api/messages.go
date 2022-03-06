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
		switch req.Method {
		case "GET":
			getMessages(w, req)
		case "POST":
			createMessage(w, req)
		}
	})
}

func getMessages(w http.ResponseWriter, req *http.Request) {
	fmt.Print("GET /messages\n")
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
	fmt.Printf("POST /messages: %v\n", newMessage.Contents)

	if err != nil {
		fmt.Printf("error parsing body as JSON: %v\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tenantId, err := uuid.Parse(req.Header.Get("Tenantid"))
	if err != nil {
		fmt.Printf("error parsing TenantId header: %v\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	userId, err := uuid.Parse(req.Header.Get("Userid"))
	if err != nil {
		fmt.Printf("error parsing UserId header: %v\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	adapter.CreateMessage(tenantId, userId, *newMessage)
	w.Write(nil)
	fmt.Println("Finished processing")
}
