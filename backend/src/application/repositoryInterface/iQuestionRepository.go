package repositoryInterface

import "refactoring-together/src/domain/entity"

type IQuestionRepository interface {
	FindById(id string) (result entity.Question, err error)
}
