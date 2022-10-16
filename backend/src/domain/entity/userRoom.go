package entity

import "time" 

type UserRoom struct {
	Id string `dynamo:"Id,hash"` 
	UserId string
	RoomId string
	CreateTime time.Time
	UpdateTime time.Time
}