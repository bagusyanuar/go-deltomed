package admin

import "github.com/bagusyanuar/go-deltomed/model"

type ComplaintRepository interface {
	GetData(divisionID string, status *uint) (data []model.Complaint, err error)
}

type ComplaintService interface {
	GetData(divisionID string, status *uint) (data []model.Complaint, err error)
}
