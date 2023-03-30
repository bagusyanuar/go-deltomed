package support

import (
	"github.com/bagusyanuar/go-deltomed/model"
	usecaseSupport "github.com/bagusyanuar/go-deltomed/usecase/support"
)

type implementsComplaintService struct {
	ComplaintRepository usecaseSupport.ComplaintRepository
}

// Send implements support.ComplaintService
func (service *implementsComplaintService) Send() (data *model.Complaint, err error) {
	return
}

func NewComplaintService(complaintRepository usecaseSupport.ComplaintRepository) usecaseSupport.ComplaintService {
	return &implementsComplaintService{
		ComplaintRepository: complaintRepository,
	}
}
