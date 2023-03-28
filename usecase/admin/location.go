package admin

import (
	"github.com/bagusyanuar/go-deltomed/http/request"
	"github.com/bagusyanuar/go-deltomed/model"
)

type LocationRepository interface {
	Create(entity model.Location) (data *model.Location, err error)
	Patch(id string, entity model.Location) (data *model.Location, err error)
	Delete(id string) (err error)
	FindAll(param string, limit int, offset int) (data []model.Location, err error)
	FindByID(id string) (data *model.Location, err error)
}

type LocationService interface {
	Create(request request.CreateLocationRequest) (data *model.Location, err error)
	Patch(id string, request request.CreateLocationRequest) (data *model.Location, err error)
	Delete(id string) (err error)
	FindAll(param string, limit int, offset int) (data []model.Location, err error)
	FindByID(id string) (data *model.Location, err error)
}
