package manager

import (
	"github.com/bagusyanuar/go-deltomed/model"
	usecaseManager "github.com/bagusyanuar/go-deltomed/usecase/manager"
	"gorm.io/gorm"
)

type implementsAuthRepository struct {
	Database *gorm.DB
}

// SignIn implements support.AuthRepository
func (repository *implementsAuthRepository) SignIn(user model.User) (data *model.User, err error) {
	if err = repository.Database.Debug().
		Where("JSON_SEARCH(roles, 'all', 'manager') IS NOT NULL").
		Where("username = ?", user.Username).
		First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func NewAuthRepository(database *gorm.DB) usecaseManager.AuthRepository {
	return &implementsAuthRepository{
		Database: database,
	}
}
