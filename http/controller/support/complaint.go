package support

import (
	"log"
	"net/http"

	"github.com/bagusyanuar/go-deltomed/common"
	"github.com/bagusyanuar/go-deltomed/http/middleware"
	"github.com/bagusyanuar/go-deltomed/http/request"
	usecaseSupport "github.com/bagusyanuar/go-deltomed/usecase/support"
	"github.com/gin-gonic/gin"
)

type ComplaintController struct {
	ComplaintService usecaseSupport.ComplaintService
	Middleware       middleware.Middleware
}

func NewComplaintController(complaintService usecaseSupport.ComplaintService) ComplaintController {
	return ComplaintController{ComplaintService: complaintService, Middleware: middleware.Middleware{}}
}

func (controller *ComplaintController) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/support")
	{
		api.POST("/complaint", controller.Middleware.Auth(), controller.Send)
	}
}

func catch(c *gin.Context) {
	if r := recover(); r != nil {
		log.Println(r)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
			Data:    nil,
		})
		return
	}
}
func (controller *ComplaintController) Send(c *gin.Context) {
	defer catch(c)
	authorizedUser := c.MustGet("user").(*common.JWTClaims)
	var request request.SendComplaintRequest
	c.Bind(&request)
	data, err := controller.ComplaintService.Send(authorizedUser.Unique, request)
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
