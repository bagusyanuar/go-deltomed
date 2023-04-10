package admin

import (
	"github.com/bagusyanuar/go-deltomed/common"
	"github.com/bagusyanuar/go-deltomed/exception"
	"github.com/bagusyanuar/go-deltomed/http/request"
	"github.com/bagusyanuar/go-deltomed/model"
	usecaseAdmin "github.com/bagusyanuar/go-deltomed/usecase/admin"
	"golang.org/x/crypto/bcrypt"
)

type implementsAuthService struct {
	AuthRepository usecaseAdmin.AuthRepository
}

// SignIn implements admin.AuthService
func (service *implementsAuthService) SignIn(request request.CreateSignInRequest) (accessToken string, err error) {
	entity := model.User{
		Username: request.Username,
	}
	user, err := service.AuthRepository.SignIn(entity)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(request.Password))
	if err != nil {
		return "", exception.ErrorPasswordNotMatch
	}

	jwtSign := common.JWTSignReturn{
		ID:       user.ID,
		Username: user.Username,
		Role:     "manager",
	}
	return common.GenerateAccessToken(&jwtSign)
}

func NewAuthService(authRepository usecaseAdmin.AuthRepository) usecaseAdmin.AuthService {
	return &implementsAuthService{
		AuthRepository: authRepository,
	}
}
