package repository

import "ims-be/model"

type UserRepository interface {
	LoginRequest(user *model.User) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	GetUserById(id int) (*model.User, error)
	GetListOfUsers(name, role *string, page int, pageSize int) ([]model.User, int, error)
	UpdateUser(userID int, updatedUser *model.User) error
	UpdatePassword(id int, oldPass string, newPass string) error
	UpdateUserStatus(userID int, status bool) error
	SetRoleUser(userID int, role int8)
}
