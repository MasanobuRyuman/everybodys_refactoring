package entity

import "time" 

type User struct {
	id string 
	password string
  name string
	createTime time.Time
	updateTime time.Time
}