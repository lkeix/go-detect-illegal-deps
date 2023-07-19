package repository

import "github.com/lkeix/example/model"

type User interface {
	FindByID(id int) (*model.User, error)
	UpdateUser(user *model.User) error
}
