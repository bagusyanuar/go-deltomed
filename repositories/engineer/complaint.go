package engineer

import (
	"github.com/bagusyanuar/go-deltomed/model"
	usecaseEngineer "github.com/bagusyanuar/go-deltomed/usecase/engineer"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type implementsComplaintRepository struct {
	Database *gorm.DB
}

// SubmitComplaint implements engineer.ComplaintRepository
func (repository *implementsComplaintRepository) SubmitComplaint(authorizedID string, id string, entity model.Complaint) (data *model.Complaint, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).
		Where("id = ?", id).
		Where("engineer_id = ?", authorizedID).
		Updates(entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// GetDataDetail implements engineer.ComplaintRepository
func (repository *implementsComplaintRepository) GetDetailComplaint(authorizedID string, id string) (data *model.Complaint, err error) {
	if err = repository.Database.Debug().Where("id = ?", id).Where("engineer_id = ?", authorizedID).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// GetData implements engineer.ComplaintRepository
func (repository *implementsComplaintRepository) GetData(authorizedID string, status string) (data []model.Complaint, err error) {
	tx := repository.Database.Debug().
		Preload("Division").
		Preload("Location")

	if status != "" {
		tx.Where("status = ?", status)
	}
	if err = tx.
		// Where("date BETWEEN ? AND ?", startDate, endDate).
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
