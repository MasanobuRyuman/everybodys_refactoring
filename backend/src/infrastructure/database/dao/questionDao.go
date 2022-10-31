package dao

import (
	"fmt"
	"refactoring-together/src/domain/entity"
	"refactoring-together/src/infrastructure/utility"
	"time"

	"github.com/guregu/dynamo"
)

type questionTable struct {
	table dynamo.Table
}

func NewQuestionTable(table dynamo.Table) *questionTable {
	return &questionTable{
		table: table,
	}
}

func (db questionTable) Add(value *entity.Question) (err error) {
	if value.UserId == "" || value.RoomId == "" {
		err = fmt.Errorf("Error in add method of question dao . \n No UserId or RoomId.")
		return
	}
	err = db.table.Put(&entity.Question{Id: utility.GetUuid(), UserId: value.UserId, RoomId: value.RoomId, Text: value.Text, Code: value.Code, CreateTime: time.Now(), UpdateTime: time.Now()}).Run()
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

func (db questionTable) Update(value *entity.Question) (err error) {
	if value.Id == "" || value.UserId == "" || value.RoomId == "" {
		err = fmt.Errorf("Error in add method of question dao. \n No UserId or RoomId.")
	}
	err = db.table.Update("Id", value.Id).Set("UserId", value.UserId).Set("RoomId", value.RoomId).Set("Text", value.Text).Set("UpdateTime", time.Now()).Run()
	return
}

func (db questionTable) Delete(id string) (err error) {
	err = db.table.Delete("Id", id).Run()
	return
}
