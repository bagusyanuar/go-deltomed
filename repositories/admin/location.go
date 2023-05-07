package admin

import (
	"github.com/bagusyanuar/go-deltomed/model"
	usecaseAdmin "github.com/bagusyanuar/go-deltomed/usecase/admin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type implementsLocationRepository struct {
	Database *gorm.DB
}

// Create implements admin.LocationRepository
func (repository *implementsLocationRepository) Create(entity model.Location) (data *model.Location, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete implements admin.LocationRepository
func (repository *implementsLocationRepository) Delete(id string) (err error) {
	if err = repository.Database.Debug().Where("id = ?", id).Delete(&model.Location{}).Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements admin.LocationRepository
func (repository *implementsLocationRepository) FindAll(param string, limit int, offset int) (data []model.Location, err error) {
	if err = repository.Database.Debug().
		Where("name LIKE ?", "%"+param+"%").
		Limit(limit).
		Offset(offset).
		Order("created_at asc").
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FindByID implements admin.LocationRepository
func (repository *implementsLocationRepository) FindByID(id string) (data *model.Location, err error) {
	if err = repository.Database.Debug().
		Where("id = ?", id).
		First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// Patch implements admin.LocationRepository
func (repository *implementsLocationRepository) Patch(id string, entity model.Location) (data *model.Location, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func NewLocationRepository(database *gorm.DB) usecaseAdmin.LocationRepository {
	return &implementsLocationRepository{
		Database: database,
	}
}
