package manager

import (
	"github.com/bagusyanuar/go-deltomed/common"
	"github.com/bagusyanuar/go-deltomed/http/request"
	"github.com/bagusyanuar/go-deltomed/http/response"
	"github.com/bagusyanuar/go-deltomed/model"
	usecaseManager "github.com/bagusyanuar/go-deltomed/usecase/manager"
	"github.com/google/uuid"
)

type implementsComplaintService struct {
	ComplaintRepository usecaseManager.ComplaintRepository
	ProfileRepository   usecaseManager.ProfileRepository
}

// SendApproval implements manager.ComplaintService
func (service *implementsComplaintService) SendApproval(id string, accessorID uuid.UUID, request request.SendApprovalRequest) (data *model.Complaint, err error) {
	engineerID, err := uuid.Parse(request.EngineeerID)
	if err != nil {
		return nil, err
	}

	entity := model.Complaint{
		Status:     request.Status,
		EngineerID: &engineerID,
		AccessorID: &accessorID,
	}
	return service.ComplaintRepository.SendApproval(id, entity)
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
				Status:     complaint.Status,
				Estimate:   complaint.Estimate,
				SupportID:  complaint.SupportID,
			},
			Location: service.transformLocation(complaint.Location),
			Division: service.transformDivision(complaint.Division),
			Support:  service.transformSupport(complaint.Support),
		})
	}
	return data, nil
}

func (service *implementsComplaintService) transformLocation(location *model.Location) *response.ComplainWithLocationScheme {
	if location != nil {
		return &response.ComplainWithLocationScheme{
			ID:   location.ID,
			Name: location.Name,
		}
	}
	return nil
}

func (service *implementsComplaintService) transformDivision(division *model.Division) *response.ComplainWithDivisionScheme {
	if division != nil {
		return &response.ComplainWithDivisionScheme{
			ID:   division.ID,
			Name: division.Name,
		}
	}
	return nil
}

func (service *implementsComplaintService) transformSupport(support *model.User) *response.ComplainWithSupportScheme {
	if support != nil {
		return &response.ComplainWithSupportScheme{
			ID:       support.ID,
			Username: support.Username,
		}
	}
	return nil
}

func NewComplaintService(complaintRepository usecaseManager.ComplaintRepository, profileRepository usecaseManager.ProfileRepository) usecaseManager.ComplaintService {
	return &implementsComplaintService{
		ComplaintRepository: complaintRepository,
		ProfileRepository:   profileRepository,
	}
}
