package repository

import (
	"errors"
	"golang-jwttoken/data/request"
	"golang-jwttoken/helper"
	"golang-jwttoken/model"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUsersRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

// Delete implements UserRepository.
func (u *UserRepositoryImpl) Delete(usersId int) {
	var users model.Users
	result := u.Db.Where("id = ?", usersId).Delete(&users)
	helper.ErrorPanic(result.Error)
}

// FindAll implements UserRepository.
func (u *UserRepositoryImpl) FindAll() []model.Users {
	var users []model.Users
	result := u.Db.Find(&users)
	helper.ErrorPanic(result.Error)
	return users
}

// FindById implements UserRepository.
func (u *UserRepositoryImpl) FindById(usersId int) (model.Users, error) {
	var users model.Users
	result := u.Db.Find(&users, usersId)
	if result != nil {
		return users, nil
	} else {
		return users, errors.New("user is not found")
	}
}

// FindByUserName implements UserRepository.
func (u *UserRepositoryImpl) FindByUserName(username string) (model.Users, error) {
	var users model.Users
	result := u.Db.First(&users, "username = ?", username)

	if result.Error != nil {
		return users, errors.New("username is not found")
	}
	return users, nil
}

// Save implements UserRepository.
func (u *UserRepositoryImpl) Save(users model.Users) {
	result := u.Db.Create(&users)
	helper.ErrorPanic(result.Error)
}

// Update implements UserRepository.
func (u *UserRepositoryImpl) Update(users model.Users) {
	var updateUsers = request.UpdateUsersRequest{
		Id:       users.Id,
		Username: users.Username,
		Email:    users.Email,
		Password: users.Password,
	}
	result := u.Db.Model(&users).Updates(updateUsers)
	helper.ErrorPanic(result.Error)
}
