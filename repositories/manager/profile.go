package manager

import (
	"github.com/bagusyanuar/go-deltomed/model"
	usecaseManager "github.com/bagusyanuar/go-deltomed/usecase/manager"
	"gorm.io/gorm"
)

type implementsProfileRepository struct {
	Database *gorm.DB
}

// GetMyProfile implements manager.ProfileRepository
func (repository *implementsProfileRepository) GetMyProfile(authorizedID string) (data *model.User, err error) {
	if err = repository.Database.Debug().Where("JSON_SEARCH(roles, 'all', 'manager') IS NOT NULL").Where("id = ?", authorizedID).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func NewProfileRepository(database *gorm.DB) usecaseManager.ProfileRepository {
	return &implementsProfileRepository{
		Database: database,
	}
}
