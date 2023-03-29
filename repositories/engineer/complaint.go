package engineer

import (
	"github.com/bagusyanuar/go-deltomed/model"
	usecaseEngineer "github.com/bagusyanuar/go-deltomed/usecase/engineer"
	"gorm.io/gorm"
)

type implementsComplaintRepository struct {
	Database *gorm.DB
}

// GetData implements engineer.ComplaintRepository
func (repository *implementsComplaintRepository) GetData(startDate string, endDate string) (data []model.Complaint, err error) {
	if err = repository.Database.Debug().Preload("Division").Preload("Location").Where("date BETWEEN ? AND ?", startDate, endDate).Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func NewComplaintRepository(database *gorm.DB) usecaseEngineer.ComplaintRepository {
	return &implementsComplaintRepository{
		Database: database,
	}
}
