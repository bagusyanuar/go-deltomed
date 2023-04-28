package admin

import (
	"github.com/bagusyanuar/go-deltomed/model"
	usecaseAdmin "github.com/bagusyanuar/go-deltomed/usecase/admin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type implementsDivisionRepository struct {
	Database *gorm.DB
}

// Delete implements admin.DivisionRepository
func (repository *implementsDivisionRepository) Delete(id string) (err error) {
	if err = repository.Database.Debug().Where("id = ?", id).Delete(&model.Division{}).Error; err != nil {
		return err
	}
	return nil
}

// Patch implements admin.DivisionRepository
func (repository *implementsDivisionRepository) Patch(id string, entity model.Division) (data *model.Division, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// FindByID implements admin.DivisionRepository
func (repository *implementsDivisionRepository) FindByID(id string) (data *model.Division, err error) {
	if err = repository.Database.Debug().
		Where("id = ?", id).
		First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// FindAll implements admin.DivisionRepository
func (repository *implementsDivisionRepository) FindAll(param string, limit int, offset int) (data []model.Division, err error) {
	if err = repository.Database.Debug().
		Where("name LIKE ?", "%"+param+"%").
		Limit(limit).
		Offset(offset).Order("created_at asc").
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// Create implements admin.DivisionRepository
func (repository *implementsDivisionRepository) Create(entity model.Division) (data *model.Division, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func NewDivisionRepository(database *gorm.DB) usecaseAdmin.DivisionRepository {
	return &implementsDivisionRepository{
		Database: database,
	}
}
