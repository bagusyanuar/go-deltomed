package admin

import (
	"encoding/json"

	"github.com/bagusyanuar/go-deltomed/common"
	"github.com/bagusyanuar/go-deltomed/http/request"
	"github.com/bagusyanuar/go-deltomed/model"
	usecaseAdmin "github.com/bagusyanuar/go-deltomed/usecase/admin"
	"golang.org/x/crypto/bcrypt"
)

type implementsUserService struct {
	UserRepository usecaseAdmin.UserRepository
}

// Create implements admin.UserService
func (service *implementsUserService) Create(request request.CreateUserRequest) (data *model.User, err error) {
	roles, _ := json.Marshal([]string{"admin"})
	var password string
	if request.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 13)
		if err != nil {
			return nil, err
		}
		password = string(hash)
	}
	entity := model.User{
		Email:    request.Email,
		Username: request.Username,
		Password: &password,
		Roles:    roles,
	}
	return service.UserRepository.Create(entity)
}

// Delete implements admin.UserService
func (service *implementsUserService) Delete(id string) (err error) {
	return service.UserRepository.Delete(id)
}

// FindAll implements admin.UserService
func (service *implementsUserService) FindAll(param string, limit int, offset int) (data []model.User, err error) {
	//make default limit = 5
	if limit == 0 {
		limit = common.DefaultLimit
	}
	return service.UserRepository.FindAll(param, limit, offset)
}

// FindByID implements admin.UserService
func (service *implementsUserService) FindByID(id string) (data *model.User, err error) {
	return service.UserRepository.FindByID(id)
}

// Patch implements admin.UserService
func (service *implementsUserService) Patch(id string, request request.CreateUserRequest) (data *model.User, err error) {
	entity := model.User{
		Email:    request.Email,
		Username: request.Username,
	}
	return service.UserRepository.Patch(id, entity)
}

func NewUserService(userRepository usecaseAdmin.UserRepository) usecaseAdmin.UserService {
	return &implementsUserService{
		UserRepository: userRepository,
	}
}
