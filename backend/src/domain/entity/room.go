package entity

import "time"

type Room struct {
	Id         string `dynamo:"Id,hash"`
	OwnerId    string
	IsOpen     bool
	CreateTime time.Time
	UpdateTime time.Time
}
