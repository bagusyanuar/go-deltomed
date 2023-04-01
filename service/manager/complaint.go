package manager

import (
	"github.com/bagusyanuar/go-deltomed/common"
	"github.com/bagusyanuar/go-deltomed/http/response"
	usecaseManager "github.com/bagusyanuar/go-deltomed/usecase/manager"
)

type implementsComplaintService struct {
	ComplaintRepository usecaseManager.ComplaintRepository
	ProfileRepository   usecaseManager.ProfileRepository
}

// GetData implements manager.ComplaintService
func (service *implementsComplaintService) GetData(authorizedID string, status string) (data []response.APIComplaintManager, err error) {
	s := common.SetStatusFilter(status)
	profile, err := service.ProfileRepository.GetMyProfile(authorizedID)
	if err != nil {
		return nil, err
	}

	var divisionID string
	if profile.DivisionID != nil {
		divisionID = profile.DivisionID.String()
	}
	complaints, err := service.ComplaintRepository.GetData(divisionID, s)

	if err != nil {
		return []response.APIComplaintManager{}, err
	}

	if len(complaints) <= 0 {
		return []response.APIComplaintManager{}, nil
	}

	for _, complaint := range complaints {
		data = append(data, response.APIComplaintManager{
			Complaint: response.Complaint{
				ID:         complaint.ID,
				DivisionID: complaint.DivisionID,
				LocationID: complaint.LocationID,
				TicketID:   complaint.TicketID,
				Date:       complaint.Date,
				Complaint:  complaint.Complaint,
				Image:      complaint.Image,
				SupportID:  complaint.SupportID,
			},
		})
	}
	return data, nil
}

func NewComplaintService(complaintRepository usecaseManager.ComplaintRepository, profileRepository usecaseManager.ProfileRepository) usecaseManager.ComplaintService {
	return &implementsComplaintService{
		ComplaintRepository: complaintRepository,
		ProfileRepository:   profileRepository,
	}
}
