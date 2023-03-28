package admin

import (
	"github.com/bagusyanuar/go-deltomed/common"
	"github.com/bagusyanuar/go-deltomed/http/request"
	"github.com/bagusyanuar/go-deltomed/model"
	usecaseAdmin "github.com/bagusyanuar/go-deltomed/usecase/admin"
)

type implementsLocationService struct {
	LocationRepository usecaseAdmin.LocationRepository
}

// Create implements admin.LocationService
func (service *implementsLocationService) Create(request request.CreateLocationRequest) (data *model.Location, err error) {
	entity := model.Location{
		Name: request.Name,
	}
	return service.LocationRepository.Create(entity)
}

// Delete implements admin.LocationService
func (service *implementsLocationService) Delete(id string) (err error) {
	return service.LocationRepository.Delete(id)
}

// FindAll implements admin.LocationService
func (service *implementsLocationService) FindAll(param string, limit int, offset int) (data []model.Location, err error) {
	//make default limit = 5
	if limit == 0 {
		limit = common.DefaultLimit
	}
	return service.LocationRepository.FindAll(param, limit, offset)
}

// FindByID implements admin.LocationService
func (service *implementsLocationService) FindByID(id string) (data *model.Location, err error) {
	return service.LocationRepository.FindByID(id)
}

// Patch implements admin.LocationService
func (service *implementsLocationService) Patch(id string, request request.CreateLocationRequest) (data *model.Location, err error) {
	entity := model.Location{
		Name: request.Name,
	}
	return service.LocationRepository.Patch(id, entity)
}

func NewLocationService(locationRepository usecaseAdmin.LocationRepository) usecaseAdmin.LocationService {
	return &implementsLocationService{
		LocationRepository: locationRepository,
	}
}
