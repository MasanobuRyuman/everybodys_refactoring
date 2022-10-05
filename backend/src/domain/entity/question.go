package entity

import "time" 

type Question struct {
	id string 
	userId string
  text string
	createTime time.Time
	updateTime time.Time
}