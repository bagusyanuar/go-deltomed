package admin

import (
	"github.com/bagusyanuar/go-deltomed/model"
	usecaseAdmin "github.com/bagusyanuar/go-deltomed/usecase/admin"
	"gorm.io/gorm"
)

type implementsAuthRepository struct {
	Database *gorm.DB
}

// SignIn implements admin.AuthRepository
func (repository *implementsAuthRepository) SignIn(user model.User) (data *model.User, err error) {
	if err = repository.Database.Debug().
		Where("JSON_SEARCH(roles, 'all', 'administrator') IS NOT NULL").
		Where("username = ?", user.Username).
		First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func NewAuthRepository(database *gorm.DB) usecaseAdmin.AuthRepository {
	return &implementsAuthRepository{
		Database: database,
	}
}
