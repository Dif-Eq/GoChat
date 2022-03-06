package model

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id        uuid.UUID
	Username  string
	Contents  string
	CreatedAt time.Time
}
