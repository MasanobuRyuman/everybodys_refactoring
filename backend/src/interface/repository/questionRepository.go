package repository

import (
	"refactoring-together/src/domain/entity"
	"refactoring-together/src/interface/repository/daoInterface"
)

type questionDao struct {
	questionDao daoInterface.IQuestionDao
}

func NewQuestionRepository(dao daoInterface.IQuestionDao) *questionDao {
	return &questionDao{
		questionDao: dao,
	}
}

func (dao *questionDao) FindByRoomId(id string) (result entity.Question, err error) {
	result, err = dao.questionDao.FindById(id)
	return
}
