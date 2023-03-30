package support

import "github.com/bagusyanuar/go-deltomed/model"

type ComplaintRepository interface {
	Send(entity model.Complaint) (data *model.Complaint, err error)
}

type ComplaintService interface {
	Send() (data *model.Complaint, err error)
}
