package engineer

import (
	"github.com/bagusyanuar/go-deltomed/http/response"
	"github.com/bagusyanuar/go-deltomed/model"
)

type ComplaintRepository interface {
	GetData(authorizedID string, startDate string, endDate string, status string) (data []model.Complaint, err error)
}

type ComplainService interface {
	GetData(authorizedID string, status string, startDate string, endDate string) (data []response.APIComplaintEngineer, err error)
}
