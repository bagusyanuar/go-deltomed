package admin

import (
	"github.com/bagusyanuar/go-deltomed/http/request"
	"github.com/bagusyanuar/go-deltomed/model"
)

type DivisionRepository interface {
	Create(entity model.Division) (data *model.Division, err error)
	Patch(id string, entity model.Division) (data *model.Division, err error)
	Delete(id string) (err error)
	FindAll(param string, limit int, offset int) (data []model.Division, err error)
	FindByID(id string) (data *model.Division, err error)
}

type DivisionService interface {
	Create(request request.CreateDivisionRequest) (data *model.Division, err error)
	Patch(id string, request request.CreateDivisionRequest) (data *model.Division, err error)
	Delete(id string) (err error)
	FindAll(param string, limit int, offset int) (data []model.Division, err error)
	FindByID(id string) (data *model.Division, err error)
}
