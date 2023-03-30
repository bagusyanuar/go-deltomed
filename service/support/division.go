package support

import (
	"github.com/bagusyanuar/go-deltomed/http/response"
	usecaseSupport "github.com/bagusyanuar/go-deltomed/usecase/support"
)

type implementsDivisionService struct {
	DivisionRepository usecaseSupport.DivisionRepository
}

// GetData implements support.DivisionService
func (service *implementsDivisionService) GetData(param string) (data []response.Division, err error) {
	divisions, err := service.DivisionRepository.GetData(param)
	if err != nil {
		return []response.Division{}, err
	}

	if len(divisions) <= 0 {
		return []response.Division{}, nil
	}

	for _, division := range divisions {
		data = append(data, response.Division{
			ID:   division.ID,
			Name: division.Name,
		})
	}
	return data, nil
}

func NewDivisionService(divisionRepository usecaseSupport.DivisionRepository) usecaseSupport.DivisionService {
	return &implementsDivisionService{
		DivisionRepository: divisionRepository,
	}
}
