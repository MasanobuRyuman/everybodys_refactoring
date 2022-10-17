package dao

import (
	"everybodys_refactoring/src/domain/entity"
	"everybodys_refactoring/src/infrastructure/utility"
	"fmt"
	"github.com/guregu/dynamo"
	"time"
)

type roomTable struct {
	table dynamo.Table
}

func NewRoomTable(table dynamo.Table) *roomTable {
	return &roomTable{
		table: table,
	}
}

func (db roomTable) Add(value *entity.Room) (err error) {
	if value.OwnerId == "" {
		err = fmt.Errorf("Error in add method of question dao. \n No ownerId")
		return
	}
	err = db.table.Put(&entity.Room{Id: utility.GetUuid(),OwnerId: value.OwnerId, IsOpen: true, CreateTime: time.Now(), UpdateTime: time.Now()}).Run()
	return
}

func (db roomTable) FindAll() (results []entity.User, err error) {
	err = db.table.Scan().All(&results)
	return
}

func (db roomTable) FindById(id string) (result entity.Room, err error) {
	err = db.table.Get("Id", id).One(&result)
	return
}

func (db roomTable) Update(value *entity.Room) (err error) {
	if value.Id == ""{ 
		err = fmt.Errorf("Error in update method of question dao. \n No id ")
		return
	}
	err = db.table.Update("Id", value.Id).Set("IsOpen", value.IsOpen).Run()
	return
}

func (db roomTable) Delete(id string) (err error) {
	err = db.table.Delete("Id", id).Run()
	return
}