package daoInterface

import "refactoring-together/src/domain/entity"

type IQuestionDao interface {
	Add(value *entity.Question) (err error)
	FindAll() (results []entity.Question, err error)
	FindById(id string) (result entity.Question, err error)
	Update(value *entity.Question) (err error)
	Delete(id string) (err error)
}
