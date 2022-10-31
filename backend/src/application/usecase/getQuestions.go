package usecase

import (
	"refactoring-together/src/application/repositoryInterface"
	"refactoring-together/src/domain/entity"
)

type GetQuestionInteractor struct {
	QuestionRepository repositoryInterface.IQuestionRepository
}

func (interactor *GetQuestionInteractor) GetQuestions(id string) (questions entity.Question, err error) {
	questions, err = interactor.QuestionRepository.FindById(id)
	return
}
