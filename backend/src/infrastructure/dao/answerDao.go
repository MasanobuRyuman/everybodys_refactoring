package dao

import (
	"everybodys_refactoring/src/domain/entity"
	"everybodys_refactoring/src/infrastructure/utility"
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

func (db answerTable) Add(userId string, questionId string, roomId string, text string) (err error) {
	err = db.table.Put(&entity.Answer{Id: utility.GetUuid(), UserId: userId, QuestionId: questionId, RoomId: roomId, Text: text, CreateTime: time.Now(), UpdateTime: time.Now()}).Run()
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

func (db answerTable) Update(value *entity.Answer) (err error) {
	err = db.table.Update("Id", value.Id).Set("Text", value.Text).Set("UpdateTime", time.Now()).Run()
	return
}

func (db answerTable) Delete(id string) (err error) {
	err = db.table.Delete("Id", id).Run()
	return
}
