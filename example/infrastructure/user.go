package infrastructure

import (
	"database/sql"
	"github.com/lkeix/example/model"
	"github.com/lkeix/example/repository"
)

type user struct {
	db *sql.DB
}

func NewUser(db *sql.DB) repository.User {
	return &user{db: db}
}

func (u *user) FindByID(id int) (*model.User, error) {
	return nil, nil
}

func (u *user) UpdateUser(user *model.User) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("update user set id=$1, name=$2 where id=$1", user.ID, user.Name)
	if err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}
