package entity

import "time" 

type Question struct {
	Id string `dynamo:"Id,hash"` 
  Text string
	CreateTime time.Time
	UpdateTime time.Time
}