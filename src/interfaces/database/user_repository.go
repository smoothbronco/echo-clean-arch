package database

import "github.com/smoothbronco/echo-clean-arch/src/domain"

type UserRepository struct {
	SqlHandler
}

func (db *UserRepository) Store(u domain.User) {
	db.Create(&u)
}

func (db *UserRepository) Select() []domain.User {
	user := []domain.User{}
	db.FindAll(&user)
	return user
}

func (db *UserRepository) SelectById(id int) domain.User {
	user := domain.User{}
	db.FindOne(&user, id)
	return user
}

func (db *UserRepository) Delete(id int) {
	user := []domain.User{}
	db.DeleteById(&user, id)
}
