package engineer

import (
	"github.com/bagusyanuar/go-deltomed/http/response"
	"github.com/bagusyanuar/go-deltomed/model"
	usecaseEngineer "github.com/bagusyanuar/go-deltomed/usecase/engineer"
)

type implementsComplaintService struct {
	ComplaintRepository usecaseEngineer.ComplaintRepository
}

// GetData implements engineer.ComplainService
func (service *implementsComplaintService) GetData(authorizedID string, status string, startDate string, endDate string) (data []response.APIComplaintEngineer, err error) {
	complaints, err := service.ComplaintRepository.GetData(authorizedID, startDate, endDate, status)
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
				SupportID:  complaint.SupportID,
				EngineerID: complaint.EngineerID,
				AccessorID: complaint.AccessorID,
			},
			Division: service.transformDivision(complaint.Division),
			Location: service.transformLocation(complaint.Location),
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

func NewComplaintService(complaintRepository usecaseEngineer.ComplaintRepository) usecaseEngineer.ComplainService {
	return &implementsComplaintService{
		ComplaintRepository: complaintRepository,
	}
}
