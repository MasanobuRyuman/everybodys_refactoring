package entity

import "time"

type Question struct {
	Id         string `dynamo:"Id,hash,UUID-index"`
	UserId     string
	RoomId     string
	Text       string
	Code       string
	CreateTime time.Time
	UpdateTime time.Time
}
