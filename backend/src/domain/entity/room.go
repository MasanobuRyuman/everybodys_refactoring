package entity

import "time"

type Room struct {
	Id         string `dynamo:"Id,hash"`
	ownerId    string
	IsOpen     bool
	CreateTime time.Time
	UpdateTime time.Time
}
