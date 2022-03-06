package model

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Contents  string    `json:"contents"`
	CreatedAt time.Time `json:"createdAt"`
}
