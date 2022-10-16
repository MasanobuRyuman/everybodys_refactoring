package entity

import "time" 

type Room struct {
	Id string `dynamo:"Id,hash"` 
	CreateTime time.Time
	UpdateTime time.Time
}