package admin

import (
	"github.com/bagusyanuar/go-deltomed/model"
	usecaseAdmin "github.com/bagusyanuar/go-deltomed/usecase/admin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type implementsUserRepository struct {
	Database *gorm.DB
}

// Create implements admin.DivisionRepository
func (repository *implementsUserRepository) Create(entity model.User) (data *model.User, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete implements admin.DivisionRepository
func (repository *implementsUserRepository) Delete(id string) (err error) {
	if err = repository.Database.Debug().Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements admin.DivisionRepository
func (repository *implementsUserRepository) FindAll(param string, limit int, offset int) (data []model.User, err error) {
	if err = repository.Database.Debug().
		Where("username LIKE ?", "%"+param+"%").Preload("Division").
		Limit(limit).
		Offset(offset).
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FindByID implements admin.DivisionRepository
func (repository *implementsUserRepository) FindByID(id string) (data *model.User, err error) {
	if err = repository.Database.Debug().
		Where("id = ?", id).
		First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// Patch implements admin.DivisionRepository
func (repository *implementsUserRepository) Patch(id string, entity model.User) (data *model.User, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func NewUserRepository(database *gorm.DB) usecaseAdmin.UserRepository {
	return &implementsUserRepository{
		Database: database,
	}
}
