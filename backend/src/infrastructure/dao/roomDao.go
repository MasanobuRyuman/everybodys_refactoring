package dao

import (
	"everybodys_refactoring/src/domain/entity"
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

func (db roomTable) Add(id string, name string) (err error) {
	err = db.table.Put(&entity.User{Id: id, Name: name, CreateTime: time.Now(), UpdateTime: time.Now()}).Run()
	return
}

func (db roomTable) FindAll() (results []entity.User, err error) {
	err = db.table.Scan().All(&results)
	return
}

func (db roomTable) FindById(id string) (result entity.User, err error) {
	err = db.table.Get("Id", id).One(&result)
	return
}

func (db roomTable) Update(id string,name string)(err error){
  err = db.table.Update("Id", id).Set("Name",name).Run()
	return
}
