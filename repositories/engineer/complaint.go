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
func (repository *implementsComplaintRepository) GetData(authorizedID string, startDate string, endDate string, status string) (data []model.Complaint, err error) {
	tx := repository.Database.Debug().
		Preload("Division").
		Preload("Location")

	if status != "" {
		tx.Where("status = ?", status)
	}
	if err = tx.Where("date BETWEEN ? AND ?", startDate, endDate).
		Where("engineer_id = ?", authorizedID).
		Order("created_at DESC").
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func NewComplaintRepository(database *gorm.DB) usecaseEngineer.ComplaintRepository {
	return &implementsComplaintRepository{
		Database: database,
	}
}
