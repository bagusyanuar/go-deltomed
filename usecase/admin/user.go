package admin

import (
	"github.com/bagusyanuar/go-deltomed/http/request"
	"github.com/bagusyanuar/go-deltomed/model"
)

type UserRepository interface {
	Create(entity model.User) (data *model.User, err error)
	Patch(id string, entity model.User) (data *model.User, err error)
	Delete(id string) (err error)
	FindAll(param string, limit int, offset int) (data []model.User, err error)
	FindByID(id string) (data *model.User, err error)
}

type UserService interface {
	Create(request request.CreateUserRequest) (data *model.User, err error)
	Patch(id string, request request.CreateUserRequest) (data *model.User, err error)
	Delete(id string) (err error)
	FindAll(param string, limit int, offset int) (data []model.User, err error)
	FindByID(id string) (data *model.User, err error)
}
