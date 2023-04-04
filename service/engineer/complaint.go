package engineer

import (
	"github.com/bagusyanuar/go-deltomed/http/request"
	"github.com/bagusyanuar/go-deltomed/http/response"
	"github.com/bagusyanuar/go-deltomed/model"
	usecaseEngineer "github.com/bagusyanuar/go-deltomed/usecase/engineer"
)

type implementsComplaintService struct {
	ComplaintRepository usecaseEngineer.ComplaintRepository
}

// SubmitComplaint implements engineer.ComplainService
func (service *implementsComplaintService) SubmitComplaint(authorizedID string, id string, request request.SubmitComplaintRequest) (data *model.Complaint, err error) {
	entity := model.Complaint{
		Estimate: request.Estimate,
		Status:   request.Status,
	}
	return service.ComplaintRepository.SubmitComplaint(authorizedID, id, entity)
}

// GetDetailComplaint implements engineer.ComplainService
func (service *implementsComplaintService) GetDetailComplaint(authorizedID string, id string) (data *response.APIComplaintEngineer, err error) {
	complaint, err := service.ComplaintRepository.GetDetailComplaint(authorizedID, id)
	if err != nil {
		return nil, err
	}
	if complaint != nil {
		return &response.APIComplaintEngineer{
			Complaint: response.Complaint{
				ID:         complaint.ID,
				DivisionID: complaint.DivisionID,
				LocationID: complaint.LocationID,
				TicketID:   complaint.TicketID,
				Date:       complaint.Date,
				Complaint:  complaint.Complaint,
				Image:      complaint.Image,
				Status:     complaint.Status,
				SupportID:  complaint.SupportID,
				EngineerID: complaint.EngineerID,
				AccessorID: complaint.AccessorID,
			},
			Division: service.transformDivision(complaint.Division),
			Location: service.transformLocation(complaint.Location),
			Support:  service.transformSupport(complaint.Support),
		}, nil
	}
	return nil, nil
}

// GetData implements engineer.ComplainService
func (service *implementsComplaintService) GetData(authorizedID string, status string) (data []response.APIComplaintEngineer, err error) {
	complaints, err := service.ComplaintRepository.GetData(authorizedID, status)
	if err != nil {
		return []response.APIComplaintEngineer{}, err
	}

	if len(complaints) <= 0 {
		return []response.APIComplaintEngineer{}, err
	}
	for _, complaint := range complaints {
		data = append(data, response.APIComplaintEngineer{
			Complaint: response.Complaint{
				ID:         complaint.ID,
				DivisionID: complaint.DivisionID,
				LocationID: complaint.LocationID,
				TicketID:   complaint.TicketID,
				Date:       complaint.Date,
				Complaint:  complaint.Complaint,
				Image:      complaint.Image,
				Status:     complaint.Status,
				SupportID:  complaint.SupportID,
				EngineerID: complaint.EngineerID,
				AccessorID: complaint.AccessorID,
			},
			Division: service.transformDivision(complaint.Division),
			Location: service.transformLocation(complaint.Location),
			Support:  service.transformSupport(complaint.Support),
		})
	}
	return data, nil
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

func (service *implementsComplaintService) transformLocation(location *model.Location) *response.ComplainWithLocationScheme {
	if location != nil {
		return &response.ComplainWithLocationScheme{
			ID:   location.ID,
			Name: location.Name,
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

func NewComplaintService(complaintRepository usecaseEngineer.ComplaintRepository) usecaseEngineer.ComplainService {
	return &implementsComplaintService{
		ComplaintRepository: complaintRepository,
	}
}
