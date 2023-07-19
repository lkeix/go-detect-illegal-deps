package repository

import "github.com/lkeix/example/model"

type Address interface {
	Update(userID int, address *model.Address) error
}
