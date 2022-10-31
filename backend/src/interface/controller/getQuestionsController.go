package controller

import (
	"refactoring-together/src/application/usecase"
	"refactoring-together/src/domain/entity"
)

type GetQuestionsHandler struct {
	GetQuestionsHandler usecase.GetQuestionInteractor
}

func (handler GetQuestionsHandler) GetQuestionsController(id string) (questions entity.Question, err error) {
	questions, err = handler.GetQuestionsHandler.GetQuestions(id)
	return
}
