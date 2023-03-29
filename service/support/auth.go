package support

import (
	"github.com/bagusyanuar/go-deltomed/exception"
	"github.com/bagusyanuar/go-deltomed/http/request"
	"github.com/bagusyanuar/go-deltomed/model"
	usecaseSupport "github.com/bagusyanuar/go-deltomed/usecase/support"
	"golang.org/x/crypto/bcrypt"
)

type implementsAuthService struct {
	AuthRepository usecaseSupport.AuthRepository
}

// SignIn implements support.AuthService
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
	return
}

func NewAuthService(authRepository usecaseSupport.AuthRepository) usecaseSupport.AuthService {
	return &implementsAuthService{
		AuthRepository: authRepository,
	}
}
