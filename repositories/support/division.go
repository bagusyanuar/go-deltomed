package support

import (
	"github.com/bagusyanuar/go-deltomed/model"
	usecaseSupport "github.com/bagusyanuar/go-deltomed/usecase/support"
	"gorm.io/gorm"
)

type implementsDivisionRepository struct {
	Database *gorm.DB
}

// GetData implements support.DivisionRepository
func (repository *implementsDivisionRepository) GetData(param string) (data []model.Division, err error) {
	if err = repository.Database.Debug().
		Where("name LIKE ?", "%"+param+"%").
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func NewDivisionRepository(database *gorm.DB) usecaseSupport.DivisionRepository {
	return &implementsDivisionRepository{
		Database: database,
	}
}
