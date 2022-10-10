package dao_test

import (
	"everybodys_refactoring/src/domain/entity"
	"everybodys_refactoring/src/infrastructure"
	"everybodys_refactoring/src/infrastructure/dao"
	"testing"
)

var db = infrastructure.GetDB()

func Test_AnswerDao(t *testing.T) {
	err := db.CreateTable("hoge", entity.Answer{}).Run()
	if err != nil {
		t.Error(err)
	}
	defer db.Table("hoge").DeleteTable().Run()
	answerDao := dao.NewAnswerTable(db.Table("hoge"))
	err = answerDao.Add("hogeId", "hogeId", "hogeId", "testText")
	if err != nil {
		t.Error(err)
	}
	err = answerDao.Add("hogeId2", "hogeId", "hogeId", "testText")
	if err != nil {
		t.Error(err)
	}
	err = answerDao.Update("hogeId2", "editText")
	if err != nil {
		t.Error(err)
	}
	answer, err := answerDao.FindAll()
	if err != nil {
		t.Error(err)
	}
	if answer[0].Id != "hogeId" || answer[0].UserId != "hogeId" || answer[0].QuestionId != "hogeId" || answer[0].Text != "testText" {
		t.Error("Data integrity failed in Answer Dao test")
	}
	answer2, err := answerDao.FindById("hogeId2")
	if err != nil {
		t.Error(err)
	}
	if answer2.Id != "hogeId2" || answer2.UserId != "hogeId" || answer2.QuestionId != "hogeId" || answer2.Text != "editText" {
		t.Error("Data integrity failed in Answer Dao test")
	}
	err = answerDao.Delete("hogeId2")
	if err != nil {
		t.Error(err)
	}
	answer, err = answerDao.FindAll()	
	if err != nil {
		t.Error(err)
	}
	if len(answer) != 1 {
		t.Error("missing delete")
	}
	return
}

func Test_QuestionDao(t *testing.T) {
	err := db.CreateTable("hoge", entity.Question{}).Run()
	if err != nil {
		t.Error(err)
	}
	defer db.Table("hoge").DeleteTable().Run()
	questionDao := dao.NewQuestionTable(db.Table("hoge"))
	err = questionDao.Add("hogeId","testText")
	if err != nil {
		t.Error(err)
	}
	err = questionDao.Add("hogeId2", "hogeId")
	if err != nil {
		t.Error(err)
	}
	err = questionDao.Update("hogeId2", "editText")
	if err != nil {
		t.Error(err)
	}
	question, err := questionDao.FindAll()
	if err != nil {
		t.Error(err)
	}
	if question[0].Id != "hogeId" || question[0].Text != "testText" {
		t.Error("Data integrity failed in Question Dao test")
	}
	question2, err := questionDao.FindById("hogeId2")
	if err != nil {
		t.Error(err)
	}
	if question2.Id != "hogeId2" || question2.Text != "editText" {
		t.Error("Data integrity failed in Question Dao test")
	}
	err = questionDao.Delete("hogeId2")
	if err != nil {
		t.Error(err)
	}
	question, err = questionDao.FindAll()	
	if err != nil {
		t.Error(err)
	}
	if len(question) != 1 {
		t.Error("missing delete")
	}
	return
}

func Test_UserDao(t *testing.T) {
	err := db.CreateTable("hoge", entity.User{}).Run()
	if err != nil {
		t.Error(err)
	}
	defer db.Table("hoge").DeleteTable().Run()
	userDao := dao.NewUserTable(db.Table("hoge"))
	err = userDao.Add("hogeId", "hogeId")
	if err != nil {
		t.Error(err)
	}
	err = userDao.Add("hogeId2", "hogeId")
	if err != nil {
		t.Error(err)
	}
	user, err := userDao.FindAll()
	if err != nil {
		t.Error(err)
	}
	if user[0].Id != "hogeId" || user[0].Name != "hogeId" {
		t.Error("Data integrity failed in Answer Dao test")
	}
	err = userDao.Update("hogeId2","editName")
	if err != nil {
		t.Error(err)
	}
  user2,err := userDao.FindById("hogeId2")
	if user2.Id != "hogeId2" || user2.Name != "editName" {
		t.Error("Data integrity failed in Answer Dao test")
	}
	return
}
