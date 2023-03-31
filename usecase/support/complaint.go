package support

import (
	"github.com/bagusyanuar/go-deltomed/http/request"
	"github.com/bagusyanuar/go-deltomed/model"
	"github.com/google/uuid"
)

type ComplaintRepository interface {
	Send(entity model.Complaint) (data *model.Complaint, err error)
}

type ComplaintService interface {
	Send(authorizedID uuid.UUID, request request.SendComplaintRequest) (data *model.Complaint, err error)
}
