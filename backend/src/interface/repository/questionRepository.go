package repository

import (
	"refactoring-together/src/domain/entity"
)

type IQuestionDao interface {
	Add(value *entity.Question) (err error)
	FindAll() (results []entity.Question, err error)
	FindById(id string) (result entity.Question, err error)
	Update(value *entity.Question) (err error)
	Delete(id string) (err error)
}

type questionDao struct {
	questionDao IQuestionDao
}

func NewQuestionRepository(dao IQuestionDao) *questionDao {
	return &questionDao{
		questionDao: dao,
	}
}

func (dao *questionDao) FindById(id string) (result entity.Question, err error) {
	result, err = dao.questionDao.FindById(id)
	return
}
