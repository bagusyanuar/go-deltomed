package manager

import (
	"github.com/bagusyanuar/go-deltomed/http/request"
	"github.com/bagusyanuar/go-deltomed/http/response"
	"github.com/bagusyanuar/go-deltomed/model"
	"github.com/google/uuid"
)

type ComplaintRepository interface {
	GetData(divisionID string, status *uint) (data []model.Complaint, err error)
	SendApproval(id string, entity model.Complaint) (data *model.Complaint, err error)
}

type ComplaintService interface {
	GetData(authorizedID string, status string) (data []response.APIComplaintManager, err error)
	SendApproval(id string, accessorID uuid.UUID, request request.SendApprovalRequest) (data *model.Complaint, err error)
}
