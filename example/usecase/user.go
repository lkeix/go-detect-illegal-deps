package usecase

import (
	"database/sql"
	"github.com/lkeix/example/infrastructure"
	"github.com/lkeix/example/model"
	"github.com/lkeix/example/repository"
)

type User interface {
	UpdateUser(user *model.User, address *model.Address) error
}

type user struct {
	repository.User
	repository.Address
}

func NewUser(db *sql.DB) User {
	return &user{
		User:    infrastructure.NewUser(db),
		Address: infrastructure.NewAddress(db),
	}
}

func (u *user) UpdateUser(user *model.User, address *model.Address) error {
	if err := u.User.UpdateUser(user); err != nil {
		return err
	}

	if err := u.Address.Update(user.ID, address); err != nil {
		return err
	}

	return nil
}
