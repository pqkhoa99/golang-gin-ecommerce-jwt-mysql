package repository

import "golang-jwttoken/model"

type UserRepository interface {
	Save(users model.Users)
	Update(users model.Users)
	Delete(usersId int)
	FindById(usersId int) (model.Users, error)
	FindAll() []model.Users
	FindByUserName(username string) (model.Users, error)
}
