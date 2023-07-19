package infrastructure

import (
	"database/sql"
	"github.com/lkeix/example/model"
	"github.com/lkeix/example/repository"
)

type address struct {
	db *sql.DB
}

func NewAddress(db *sql.DB) repository.Address {
	return &address{db: db}
}

func (a *address) Update(userID int, address *model.Address) error {
	tx, err := a.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("update address set postal=$1 street=$2 where user_id=$3", address.Postal, address.Street, userID)
	if err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}
