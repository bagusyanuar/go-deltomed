package support

import (
	"github.com/bagusyanuar/go-deltomed/http/response"
	"github.com/bagusyanuar/go-deltomed/model"
)

type LocationRepository interface {
	GetData(param string) (data []model.Location, err error)
}

type LocationService interface {
	GetData(param string) (data []response.Location, err error)
}
