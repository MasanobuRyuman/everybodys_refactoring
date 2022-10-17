package dao

import (
	"everybodys_refactoring/src/domain/entity"
	"everybodys_refactoring/src/infrastructure/utility"
	"fmt"
	"github.com/guregu/dynamo"
	"time"
)

type userTable struct {
	table dynamo.Table
}

func NewUserTable(table dynamo.Table) *userTable {
	return &userTable{
		table: table,
	}
}

func (db userTable) Add(value *entity.User) (err error) {
	err = db.table.Put(&entity.User{Id: utility.GetUuid(), Name: value.Name, CreateTime: time.Now(), UpdateTime: time.Now()}).Run()
	return
}

func (db userTable) FindAll() (results []entity.User, err error) {
	err = db.table.Scan().All(&results)
	return
}

func (db userTable) FindById(id string) (result entity.User, err error) {
	err = db.table.Get("Id", id).One(&result)
	return
}

func (db userTable) Update(value *entity.User) (err error) {
	if value.Id == "" {
		err = fmt.Errorf("Error in update method of question dao. \n No id ")
		return
	}
	err = db.table.Update("Id", value.Id).Set("Name", value.Name).Set("UpdateTime", time.Now()).Run()
	return
}
