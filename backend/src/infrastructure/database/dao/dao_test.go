package dao_test

import (
	"refactoring-together/src/domain/entity"
	"refactoring-together/src/infrastructure/database"
	"refactoring-together/src/infrastructure/database/dao"
	"testing"
)

var db = database.GetDB()

func Test_AnswerDao(t *testing.T) {
	err := db.CreateTable("hoge", entity.Answer{}).Run()
	if err != nil {
		t.Error(err)
	}
	defer db.Table("hoge").DeleteTable().Run()
	answerDao := dao.NewAnswerTable(db.Table("hoge"))

	err = answerDao.Add(&entity.Answer{Id: "hogeId", UserId: "hogeId", QuestionId: "hogeRoomId", RoomId: "hogeRoomId", Text: "testText", Code: "<p>test</p>"})
	if err != nil {
		t.Error(err)
	}

	err = answerDao.Add(&entity.Answer{Id: "hogeId", UserId: "hogeId", QuestionId: "hogeRoomId", RoomId: "hogeRoomId", Text: "testText", Code: "<p>test</p>"})
	if err != nil {
		t.Error(err)
	}

	answers, err := answerDao.FindAll()
	if err != nil {
		t.Error(err)
	}
	if len(answers) != 2 {
		t.Error("Data integrity failed in Answer Dao test")
	}

	err = answerDao.Update(&entity.Answer{Id: answers[0].Id, UserId: answers[0].UserId, QuestionId: answers[0].QuestionId, Text: "editText", Code: answers[0].Code})
	if err != nil {
		t.Error(err)
	}

	answer, err := answerDao.FindById(answers[0].Id)
	if err != nil {
		t.Error(err)
	}
	if answer.UserId != answers[0].UserId || answer.Text != "editText" {
		t.Error("Data integrity failed in Answer Dao test")
	}

	err = answerDao.Delete(answers[0].Id)
	if err != nil {
		t.Error(err)
	}
	answers, err = answerDao.FindAll()
	if err != nil {
		t.Error(err)
	}
	if len(answers) != 1 {
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
	err = questionDao.Add(&entity.Question{UserId: "hogeUserId", RoomId: "hogeRoomId", Text: "hogeText", Code: "<p>test</p>"})
	if err != nil {
		t.Error(err)
	}
	err = questionDao.Add(&entity.Question{UserId: "hogeUserId", RoomId: "hogeRoomId", Text: "hogeText2", Code: "<p>test</p>"})
	if err != nil {
		t.Error(err)
	}
	questions, err := questionDao.FindAll()
	if err != nil {
		t.Error(err)
	}
	if len(questions) != 2 {
		t.Error("Data integrity failed in Question Dao test")
	}
	err = questionDao.Update(&entity.Question{Id: questions[0].Id, UserId: questions[0].UserId, RoomId: questions[0].RoomId, Text: "editText", Code: questions[0].Code})
	question, err := questionDao.FindById(questions[0].Id)
	if err != nil {
		t.Error(err)
	}
	if question.Id != questions[0].Id {
		t.Error("not getting the correct value")
	}
	if question.Text != "editText" {
		t.Error("Question Dao has not been updated")
	}
	err = questionDao.Delete(questions[0].Id)
	if err != nil {
		t.Error(err)
	}
	questions, err = questionDao.FindAll()
	if err != nil {
		t.Error(err)
	}
	if len(questions) != 1 {
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
	err = userDao.Add(&entity.User{Name: "hogeName"})
	if err != nil {
		t.Error(err)
	}
	err = userDao.Add(&entity.User{Name: "hogeName2"})
	if err != nil {
		t.Error(err)
	}
	users, err := userDao.FindAll()
	if err != nil {
		t.Error(err)
	}
	if len(users) != 2 {
		t.Error("Data integrity failed in Answer Dao test")
	}
	err = userDao.Update(&entity.User{Id: users[0].Id, Name: "editName"})
	if err != nil {
		t.Error(err)
	}
	user, err := userDao.FindById(users[0].Id)
	if err != nil {
		t.Error(err)
	}
	if user.Id != users[0].Id {
		t.Error("not getting the correct value")
	}
	if user.Name != "editName" {
		t.Error("Question Dao has not been updated")
	}
	return
}

func Test_RoomDao(t *testing.T) {
	err := db.CreateTable("hoge", entity.User{}).Run()
	if err != nil {
		t.Error(err)
	}
	defer db.Table("hoge").DeleteTable().Run()
	roomDao := dao.NewRoomTable(db.Table("hoge"))
	err = roomDao.Add(&entity.Room{OwnerId: "hogeOwnerId"})
	if err != nil {
		t.Error(err)
	}
	err = roomDao.Add(&entity.Room{OwnerId: "hogeOwnerId2"})
	if err != nil {
		t.Error(err)
	}
	rooms, err := roomDao.FindAll()
	if err != nil {
		t.Error(err)
	}
	if len(rooms) != 2 {
		t.Error("Data integrity failed in Answer Dao test")
	}
	err = roomDao.Update(&entity.Room{Id: rooms[0].Id, IsOpen: false})
	if err != nil {
		t.Error(err)
	}
	room, err := roomDao.FindById(rooms[0].Id)
	if err != nil {
		t.Error(err)
	}
	if room.IsOpen {
		t.Error("Room Dao has not been updated")
	}
	err = roomDao.Delete(rooms[0].Id)
	if err != nil {
		t.Error(err)
	}
	rooms, err = roomDao.FindAll()
	if err != nil {
		t.Error(err)
	}
	if len(rooms) != 1 {
		t.Error("missing delete")
	}
}

func Test_UserRoomDao(t *testing.T) {
	err := db.CreateTable("hoge", entity.User{}).Run()
	if err != nil {
		t.Error(err)
	}
	defer db.Table("hoge").DeleteTable().Run()
	userRoomDao := dao.NewUserRoomTable(db.Table("hoge"))
	err = userRoomDao.Add(&entity.UserRoom{UserId: "hogeUserId", RoomId: "hogeRoomId"})
	if err != nil {
		t.Error(err)
	}
	err = userRoomDao.Add(&entity.UserRoom{UserId: "hogeUserId", RoomId: "hogeRoomId2"})
	if err != nil {
		t.Error(err)
	}
	rooms, err := userRoomDao.FindAll()
	if err != nil {
		t.Error(err)
	}
	if len(rooms) != 2 {
		t.Error("Data integrity failed in UserRoom Dao test")
	}
	room, err := userRoomDao.FindById(rooms[0].Id)
	if err != nil {
		t.Error(err)
	}
	if room.Id != rooms[0].Id {
		t.Error("Data integrity failed in Answer Dao test")
	}
	err = userRoomDao.Delete(rooms[0].Id)
	if err != nil {
		t.Error(err)
	}
	rooms, err = userRoomDao.FindAll()
	if err != nil {
		t.Error(err)
	}
	if len(rooms) != 1 {
		t.Error("Data integrity failed in UserRoom Dao test")
	}
}
