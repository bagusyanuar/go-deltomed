package manager

import (
	"github.com/bagusyanuar/go-deltomed/http/response"
	"github.com/bagusyanuar/go-deltomed/model"
)

type ComplaintRepository interface {
	GetData(divisionID string, status *uint) (data []model.Complaint, err error)
}

type ComplaintService interface {
	GetData(divisionID string, status string) (data []response.APIComplaintManager, err error)
}
