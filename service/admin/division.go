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

// FindByID implements admin.DivisionService
func (service *implementsDivisionService) FindByID(id string) (data *model.Division, err error) {
	return service.DivisionRepository.FindByID(id)
}

// FindAll implements admin.DivisionService
func (service *implementsDivisionService) FindAll(param string, limit int, offset int) (data []model.Division, err error) {
	//make default offset = 5
	if offset == 0 {
		offset = common.DefaultOffset
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