package manager

import (
	"log"
	"net/http"

	"github.com/bagusyanuar/go-deltomed/common"
	"github.com/bagusyanuar/go-deltomed/http/middleware"
	"github.com/bagusyanuar/go-deltomed/http/request"
	usecaseManager "github.com/bagusyanuar/go-deltomed/usecase/manager"
	"github.com/gin-gonic/gin"
)

type ComplaintController struct {
	ComplaintService usecaseManager.ComplaintService
	Middleware       middleware.Middleware
}

func NewComplaintController(complaintService usecaseManager.ComplaintService) ComplaintController {
	return ComplaintController{ComplaintService: complaintService, Middleware: middleware.Middleware{}}
}

func (controller *ComplaintController) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/manager")
	{
		api.GET("/complaint", controller.Middleware.Auth(), controller.GetData)
		api.POST("/complaint/:id", controller.Middleware.Auth(), controller.SendApproval)
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

func (controller *ComplaintController) GetData(c *gin.Context) {
	status := c.Query("status")
	defer catch(c)
	authorizedUser := c.MustGet("user").(*common.JWTClaims)
	authorizedID := authorizedUser.Unique.String()

	data, err := controller.ComplaintService.GetData(authorizedID, status)
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

func (controller *ComplaintController) SendApproval(c *gin.Context) {
	defer catch(c)
	id := c.Param("id")
	var request request.SendApprovalRequest
	c.BindJSON(&request)

	authorizedUser := c.MustGet("user").(*common.JWTClaims)
	_, err := controller.ComplaintService.SendApproval(id, authorizedUser.Unique, request)
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
		Data:    nil,
	})
}
