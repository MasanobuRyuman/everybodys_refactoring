package repository

import (
	"refactoring-together/src/domain/entity"
)

type QuestionDao interface {
	Add(value *entity.Question) (err error)
	FindAll() (results []entity.Question, err error)
	FindById(id string) (result entity.Question, err error)
	Update(value *entity.Question) (err error)
	Delete(id string) (err error)
}

type questionDao struct {
	dao QuestionDao
}

func NewQuestionRepository(dao QuestionDao) *questionDao {
	return &questionDao{
		dao: dao,
	}
}

func (dao *questionDao) FindById(id string) (result entity.Question, err error) {
	result, err = dao.dao.FindById(id)
	return
}
