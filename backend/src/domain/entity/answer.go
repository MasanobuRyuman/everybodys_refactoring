package entity

import "time" 

type Answer struct {
	id string 
	userId string
  questionId string
  text string
	createTime time.Time
	updateTime time.Time
}