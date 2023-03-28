package admin

import "github.com/bagusyanuar/go-deltomed/model"

type UserRepository interface {
	Create(entity model.User) (data *model.User, err error)
}
