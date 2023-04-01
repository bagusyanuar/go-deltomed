package manager

import (
	"github.com/bagusyanuar/go-deltomed/http/request"
	"github.com/bagusyanuar/go-deltomed/model"
)

type AuthRepository interface {
	SignIn(user model.User) (data *model.User, err error)
}

type AuthService interface {
	SignIn(request request.CreateSignInRequest) (accessToken string, err error)
}