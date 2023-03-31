package manager

import (
	"github.com/bagusyanuar/go-deltomed/http/response"
	usecaseManager "github.com/bagusyanuar/go-deltomed/usecase/manager"
)

type implementsComplaintService struct {
	ComplaintRepository usecaseManager.ComplaintRepository
}

// GetData implements manager.ComplaintService
func (service *implementsComplaintService) GetData(divisionID string, status string) (data []response.APIComplaintManager, err error) {
	// s := new(uint)
	// switch status {
	// case "0":
	// 	tmp := common.StatusPending
	// 	s = &tmp
	// }
	// complaints, e := service.ComplaintRepository.GetData(divisionID, s)

	// for _, complaint := range complaints {
	// 	append(data, response.APIComplaintManager{})
	// }
	return
}

func NewComplaintService(complaintRepository usecaseManager.ComplaintRepository) usecaseManager.ComplaintService {
	return &implementsComplaintService{
		ComplaintRepository: complaintRepository,
	}
}
