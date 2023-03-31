package support

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/bagusyanuar/go-deltomed/common"
	"github.com/bagusyanuar/go-deltomed/http/request"
	"github.com/bagusyanuar/go-deltomed/model"
	usecaseSupport "github.com/bagusyanuar/go-deltomed/usecase/support"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type implementsComplaintService struct {
	ComplaintRepository usecaseSupport.ComplaintRepository
}

// Send implements support.ComplaintService
func (service *implementsComplaintService) Send(authorizedID uuid.UUID, request request.SendComplaintRequest) (data *model.Complaint, err error) {

	image := new(string)
	if request.Image != nil {
		// if _, err := os.Stat(common.ImagePath); os.IsNotExist(err) {
		err := os.Mkdir(common.ImagePath, os.ModePerm)
		if err != nil {
			return nil, err
		}
		// }

		ext := filepath.Ext(request.Image.Filename)
		fileName := fmt.Sprintf("%s/%s%s", common.ImagePath, uuid.New().String(), ext)
		image = &fileName
	}

	divisionID, err := uuid.Parse(request.DivisionID)
	if err != nil {
		return nil, err
	}

	locationID, err := uuid.Parse(request.LocationID)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	formattedTime := now.Format("20060102150405")
	ticketID := fmt.Sprintf("TC-%s", formattedTime)
	entity := model.Complaint{
		DivisionID: divisionID,
		LocationID: locationID,
		TicketID:   ticketID,
		Date:       datatypes.Date(time.Now()),
		Complaint:  request.Complaint,
		Image:      *image,
		SupportID:  authorizedID,
	}
	return service.ComplaintRepository.Send(entity)
}

func NewComplaintService(complaintRepository usecaseSupport.ComplaintRepository) usecaseSupport.ComplaintService {
	return &implementsComplaintService{
		ComplaintRepository: complaintRepository,
	}
}
