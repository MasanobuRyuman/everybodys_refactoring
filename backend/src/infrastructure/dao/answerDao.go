package dao 

import (
	"everybodys_refactoring/src/domain/entity"
	"github.com/guregu/dynamo"
)

type answerTable struct {
	table dynamo.Table 
}

func NewAnswerTable(table dynamo.Table)*answerTable{
  return &answerTable{
		table: table,
	} 
}

func(db answerTable) FindAll() (results []entity.User, err error){
	err = db.table.Scan().All(&results)
	return
}

func(db answerTable) FindId(id string) (result entity.User, err error){
	err = db.table.Get("id",id).One(result)
  return
}