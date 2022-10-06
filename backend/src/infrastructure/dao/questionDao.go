package dao 

import (
	"everybodys_refactoring/src/domain/entity"
	"github.com/guregu/dynamo"
)

type questionTable struct {
	table dynamo.Table 
}

func NewQuestionTable(table dynamo.Table)*questionTable{
  return &questionTable{
		table: table,
	} 
}

func(db questionTable) FindAll() (results []entity.User, err error){
	err = db.table.Scan().All(&results)
	return
}

func(db questionTable) FindId(id string) (result entity.User, err error){
	err = db.table.Get("id",id).One(result)
  return
}