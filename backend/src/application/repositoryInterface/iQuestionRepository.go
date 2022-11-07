package repositoryInterface

import "refactoring-together/src/domain/entity"

type IQuestionRepository interface {
	FindByRoomId(id string) (result entity.Question, err error)
}
