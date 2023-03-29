package support

import (
	"github.com/bagusyanuar/go-deltomed/model"
	usecaseSupport "github.com/bagusyanuar/go-deltomed/usecase/support"
	"gorm.io/gorm"
)

type implementsAuthRepository struct {
	Database *gorm.DB
}

// SignIn implements support.AuthRepository
func (repository *implementsAuthRepository) SignIn(user model.User) (data *model.User, err error) {
	if err = repository.Database.Debug().Where("username = ?", user.Username).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func NewAuthRepository(database *gorm.DB) usecaseSupport.AuthRepository {
	return &implementsAuthRepository{
		Database: database,
	}
}
