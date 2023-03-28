package admin

import (
	"github.com/bagusyanuar/go-deltomed/common"
	"github.com/bagusyanuar/go-deltomed/http/request"
	"github.com/bagusyanuar/go-deltomed/model"
	usecaseAdmin "github.com/bagusyanuar/go-deltomed/usecase/admin"
)

type implementsDivisionService struct {
	DivisionRepository usecaseAdmin.DivisionRepository
}

// Delete implements admin.DivisionService
func (service *implementsDivisionService) Delete(id string) (err error) {
	return service.DivisionRepository.Delete(id)
}

// Patch implements admin.DivisionService
func (service *implementsDivisionService) Patch(id string, request request.CreateDivisionRequest) (data *model.Division, err error) {
	entity := model.Division{
		Name: request.Name,
	}
	return service.DivisionRepository.Patch(id, entity)
}

// FindByID implements admin.DivisionService
func (service *implementsDivisionService) FindByID(id string) (data *model.Division, err error) {
	return service.DivisionRepository.FindByID(id)
}

// FindAll implements admin.DivisionService
func (service *implementsDivisionService) FindAll(param string, limit int, offset int) (data []model.Division, err error) {
	//make default limit = 5
	if limit == 0 {
		limit = common.DefaultLimit
	}
	return service.DivisionRepository.FindAll(param, limit, offset)
}

// Create implements admin.DivisionService
func (service *implementsDivisionService) Create(request request.CreateDivisionRequest) (data *model.Division, err error) {
	entity := model.Division{
		Name: request.Name,
	}
	return service.DivisionRepository.Create(entity)
}

func NewDivisionService(divisionRepository usecaseAdmin.DivisionRepository) usecaseAdmin.DivisionService {
	return &implementsDivisionService{
		DivisionRepository: divisionRepository,
	}
}
