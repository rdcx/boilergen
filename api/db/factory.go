package db

import (
	"lightban/api/model"
	"lightban/util"

	"github.com/brianvoe/gofakeit/v6"
)

func (db *DB) CreateRandomUser() *model.User {
	name := util.RandomStringWithPrefix("user-")

	user, err := db.CreateUser(name, gofakeit.Email(), gofakeit.Password(true, true, true, true, true, 32))

	if err != nil {
		panic(err)
	}

	return user
}

func (db *DB) CreateRandomProject() *model.Project {
	name := util.RandomStringWithPrefix("project-")

	proj := &model.Project{
		User: db.CreateRandomUser(),
		Name: name,
	}

	err := db.CreateProject(proj)

	if err != nil {
		panic(err)
	}

	return proj
}
