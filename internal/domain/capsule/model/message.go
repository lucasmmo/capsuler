package model

import (
	"time"
)

type Message struct {
	Id        string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	CapsuleId string
	UserId    string
}
