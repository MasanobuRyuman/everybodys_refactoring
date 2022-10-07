package entity

import "time" 

type Answer struct {
	Id string `dynamo:"Id,hash"`
	UserId string
  QuestionId string 
  Text string
	CreateTime time.Time
	UpdateTime time.Time
}