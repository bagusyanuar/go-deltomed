package support

import (
	"github.com/bagusyanuar/go-deltomed/model"
	usecaseSupport "github.com/bagusyanuar/go-deltomed/usecase/support"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type implementsComplaintRepository struct {
	Database *gorm.DB
}

// Send implements support.ComplaintRepository
func (repository *implementsComplaintRepository) Send(entity model.Complaint) (data *model.Complaint, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func NewComplaintRepository(database *gorm.DB) usecaseSupport.ComplaintRepository {
	return &implementsComplaintRepository{
		Database: database,
	}
}
