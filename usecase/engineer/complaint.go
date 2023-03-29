package engineer

import (
	"github.com/bagusyanuar/go-deltomed/http/response"
	"github.com/bagusyanuar/go-deltomed/model"
)

type ComplaintRepository interface {
	GetData(startDate string, endDate string) (data []model.Complaint, err error)
}

type ComplainService interface {
	GetData(startDate string, endDate string) (data []response.APIComplaintEngineer, err error)
}
