package entity

import "time"

type User struct {
	Id         string `dynamo:"Id,hash"`
	Name       string
	CreateTime time.Time
	UpdateTime time.Time
}
