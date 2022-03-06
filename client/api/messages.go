package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Contents  string    `json:"contents"`
	CreatedAt time.Time `json:"createdAt"`
}

func GetMessages() []Message {
	req, err := http.NewRequest("GET", "http://localhost:8080/messages", nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	var responseObject []Message
	json.Unmarshal(bodyBytes, &responseObject)

	return responseObject
}

func CreateMessage(message Message) {
	json_data, err := json.Marshal(message)

	fmt.Println(string(json_data))
	if err != nil {
		fmt.Print(err.Error())
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/messages", bytes.NewBuffer(json_data))
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Userid", "7d536c63-66c1-4c62-b353-dd7e129cee1d")
	req.Header.Add("Tenantid", "cd5e3ffb-5064-430f-93de-56c91e806c87")

	reqBytes, _ := httputil.DumpRequestOut(req, true)

	fmt.Print("\nREQUEST START\n")

	fmt.Print(string(reqBytes))

	fmt.Print("\nREQUEST END\n")

	client := &http.Client{}
	response, err := client.Do(req)

	fmt.Print(response)

	if err != nil {
		fmt.Print(err.Error())
	}
}
