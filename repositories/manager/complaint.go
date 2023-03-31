package manager

import (
	"github.com/bagusyanuar/go-deltomed/model"
	usecaseManager "github.com/bagusyanuar/go-deltomed/usecase/manager"
	"gorm.io/gorm"
)

type implementsComplaintRepository struct {
	Database *gorm.DB
}

// GetData implements manager.ComplaintRepository
func (repository *implementsComplaintRepository) GetData(divisionID string, status *uint) (data []model.Complaint, err error) {
	tx := repository.Database.Debug().Preload("Division").Preload("Location")

	if status != nil {
		tx.Where("status = ?", &status)
	}

	if err = tx.Where("division_id = ?", divisionID).Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func NewComplaintRepository(database *gorm.DB) usecaseManager.ComplaintRepository {
	return &implementsComplaintRepository{
		Database: database,
	}
}
