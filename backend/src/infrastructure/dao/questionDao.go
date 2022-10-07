package dao

import (
	"everybodys_refactoring/src/domain/entity"
	"github.com/guregu/dynamo"
	"time"
)

type questionTable struct {
	table dynamo.Table
}

func NewQuestionTable(table dynamo.Table) *questionTable {
	return &questionTable{
		table: table,
	}
}

func (db questionTable) Add(id string, text string) (err error) {
	err = db.table.Put(&entity.Question{Id: id, Text: text, CreateTime: time.Now(), UpdateTime: time.Now()}).Run()
	return
}

func (db questionTable) FindAll() (results []entity.Question, err error) {
	err = db.table.Scan().All(&results)
	return
}

func (db questionTable) FindById(id string) (result entity.Question, err error) {
	err = db.table.Get("Id", id).One(&result)
	return
}

func (db questionTable) Update(id string, text string) (err error) {
	err = db.table.Update("Id", id).Set("Text", text).Set("UpdateTime", time.Now()).Run()
	return
}

func (db questionTable) Delete(id string) (err error) {
	err = db.table.Delete("Id", id).Run()
	return
}
