package dao 

import (
	"everybodys_refactoring/src/domain/entity"
	"github.com/guregu/dynamo"
)

type userTable struct {
	table dynamo.Table 
}

func NewUserTable(table dynamo.Table)*userTable{
  return &userTable{
		table: table,
	} 
}

func(db userTable) FindAll() (results []entity.User, err error){
	err = db.table.Scan().All(&results)
	return
}

func(db userTable) FindId(id string) (result entity.User, err error){
	err = db.table.Get("id",id).One(result)
  return
}