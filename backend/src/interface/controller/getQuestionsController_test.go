package controller_test

import (
	"fmt"
	"os"
	"refactoring-together/src/application/usecase"
	"refactoring-together/src/domain/entity"
	"refactoring-together/src/infrastructure/database"
	"refactoring-together/src/infrastructure/database/dao"
	"refactoring-together/src/interface/controller"
	"testing"

	"github.com/guregu/dynamo"
	"github.com/joho/godotenv"
)

var db *dynamo.DB

func init() {
	var err = godotenv.Load("../../../.env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
	db = database.GetDB(os.Getenv("DATABASE_ENDPOINT"))
}

func Test_getQuestionsController(t *testing.T) {
	err := db.CreateTable("hoge", entity.Question{}).Run()
	if err != nil {
		t.Error(err)
	}
	defer db.Table("hoge").DeleteTable().Run()
	questionDao := dao.NewQuestionTable(db.Table("hoge"))
	err = questionDao.Add(&entity.Question{UserId: "hogeUserId", RoomId: "hogeRoomId", Text: "hogeText", Code: "<p>test</p>"})
	if err != nil {
		t.Error(err)
	}
	questions, err := questionDao.FindAll()
	if err != nil {
		t.Error(err)
	}
	getQuestionsInteractor := new(usecase.GetQuestionInteractor)
	getQuestionsInteractor.QuestionRepository = questionDao
	getQuestionsHandler := controller.GetQuestionsHandler{}
	getQuestionsHandler.GetQuestionsHandler = *getQuestionsInteractor
	question, err := getQuestionsHandler.GetQuestionsController(questions[0].Id)
	if err != nil {
		t.Error(err)
	}
	if question.UserId != "hogeUserId" {
		t.Error("Data integrity failed in Question Dao test")
	}
	return
}
