package engineer

import (
	"net/http"

	"github.com/bagusyanuar/go-deltomed/common"
	usecaseEngineer "github.com/bagusyanuar/go-deltomed/usecase/engineer"
	"github.com/gin-gonic/gin"
)

type ComplaintController struct {
	ComplaintService usecaseEngineer.ComplainService
}

func NewComplaintController(complaintService usecaseEngineer.ComplainService) ComplaintController {
	return ComplaintController{ComplaintService: complaintService}
}

func (controller *ComplaintController) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/engineer")
	{
		api.GET("/complaint", controller.GetData)
	}
}

func (controller *ComplaintController) GetData(c *gin.Context) {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	data, err := controller.ComplaintService.GetData(startDate, endDate)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, common.APIResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}
