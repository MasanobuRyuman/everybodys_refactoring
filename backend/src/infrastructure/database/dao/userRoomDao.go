package dao

import (
	"everybodys_refactoring/src/domain/entity"
	"everybodys_refactoring/src/infrastructure/utility"
	"fmt"
	"github.com/guregu/dynamo"
	"time"
)

type userRoomTable struct {
	table dynamo.Table
}

func NewUserRoomTable(table dynamo.Table) *userRoomTable {
	return &userRoomTable{
		table: table,
	}
}

func (db userRoomTable) Add(value *entity.UserRoom) (err error) {
	if value.UserId == "" || value.RoomId == "" {
		err = fmt.Errorf("Error in add method of question dao . \n No UserId or RoomId.")
		return
	}
	err = db.table.Put(&entity.Question{Id: utility.GetUuid(), UserId: value.UserId, RoomId: value.RoomId, CreateTime: time.Now(), UpdateTime: time.Now()}).Run()
	return
}

func (db userRoomTable) FindAll() (results []entity.UserRoom, err error) {
	err = db.table.Scan().All(&results)
	return
}

func (db userRoomTable) FindById(id string) (result entity.UserRoom, err error) {
	err = db.table.Get("Id", id).One(&result)
	return
}

func (db userRoomTable) Delete(id string) (err error) {
	err = db.table.Delete("Id", id).Run()
	return
}
