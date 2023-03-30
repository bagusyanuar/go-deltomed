package support

import (
	"github.com/bagusyanuar/go-deltomed/http/response"
	"github.com/bagusyanuar/go-deltomed/model"
)

type DivisionRepository interface {
	GetData(param string) (data []model.Division, err error)
}

type DivisionService interface {
	GetData(param string) (data []response.Division, err error)
}
