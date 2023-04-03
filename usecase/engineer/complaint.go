package engineer

import (
	"github.com/bagusyanuar/go-deltomed/http/request"
	"github.com/bagusyanuar/go-deltomed/http/response"
	"github.com/bagusyanuar/go-deltomed/model"
)

type ComplaintRepository interface {
	GetData(authorizedID string, status string) (data []model.Complaint, err error)
	GetDetailComplaint(authorizedID string, id string) (data *model.Complaint, err error)
	SubmitComplaint(authorizedID string, id string, entity model.Complaint) (data *model.Complaint, err error)
}

type ComplainService interface {
	GetData(authorizedID string, status string) (data []response.APIComplaintEngineer, err error)
	GetDetailComplaint(authorizedID string, id string) (data *response.APIComplaintEngineer, err error)
	SubmitComplaint(authorizedID string, id string, request request.SubmitComplaintRequest) (data *model.Complaint, err error)
}
