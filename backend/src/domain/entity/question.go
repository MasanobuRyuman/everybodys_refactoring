package entity

import "time"

type Question struct {
	Id         string `dynamo:"Id,hash,UUID-index"`
	RoomId     string
	UserId     string
	Text       string
	CreateTime time.Time
	UpdateTime time.Time
}
