package dao

import (
	"everybodys_refactoring/src/domain/entity"
	"github.com/guregu/dynamo"
	"time"
)

type answerTable struct {
	table dynamo.Table
}

func NewAnswerTable(table dynamo.Table) *answerTable {
	return &answerTable{
		table: table,
	}
}

type User struct {
	UserID string `dynamo:"UserID,hash"`
	Name   string `dynamo:"Name,range"`
	Age    int    `dynamo:"Age"`
	Text   string `dynamo:"Text"`
}

func (db answerTable) Add(id string, userId string, questionId string, text string) (err error) {
	err = db.table.Put(&entity.Answer{Id: id, UserId: userId, QuestionId: questionId, Text: text, CreateTime: time.Now(), UpdateTime: time.Now()}).Run()
	return
}

func (db answerTable) FindAll() (results []entity.Answer, err error) {
	err = db.table.Scan().All(&results)
	return
}

func (db answerTable) FindById(id string) (result entity.Answer, err error) {
	err = db.table.Get("Id", id).One(&result)
	return
}

func (db answerTable) Update(id string, text string) (err error) {
	err = db.table.Update("Id", id).Set("Text", text).Set("UpdateTime", time.Now()).Run()
	return
}

func (db answerTable) Delete(id string) (err error) {
	err = db.table.Delete("Id", id).Run()
	return
}
