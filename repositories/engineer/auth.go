package engineer

import (
	"github.com/bagusyanuar/go-deltomed/model"
	usecaseEngineer "github.com/bagusyanuar/go-deltomed/usecase/engineer"
	"gorm.io/gorm"
)

type implementsAuthRepository struct {
	Database *gorm.DB
}

// SignIn implements engineer.AuthRepository
func (repository *implementsAuthRepository) SignIn(user model.User) (data *model.User, err error) {
	if err = repository.Database.Debug().
		Where("JSON_SEARCH(roles, 'all', 'engineer') IS NOT NULL").
		Where("username = ?", user.Username).
		First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func NewAuthRepository(database *gorm.DB) usecaseEngineer.AuthRepository {
	return &implementsAuthRepository{
		Database: database,
	}
}
