package engineer

import (
	"net/http"

	"github.com/bagusyanuar/go-deltomed/common"
	"github.com/bagusyanuar/go-deltomed/http/middleware"
	"github.com/bagusyanuar/go-deltomed/http/request"
	usecaseEngineer "github.com/bagusyanuar/go-deltomed/usecase/engineer"
	"github.com/gin-gonic/gin"
)

type ComplaintController struct {
	ComplaintService usecaseEngineer.ComplainService
	Middleware       middleware.Middleware
}

func NewComplaintController(complaintService usecaseEngineer.ComplainService) ComplaintController {
	return ComplaintController{ComplaintService: complaintService, Middleware: middleware.Middleware{}}
}

func (controller *ComplaintController) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/engineer")
	{
		api.GET("/complaint", controller.Middleware.Auth(), controller.GetData)
		api.GET("/complaint/:id", controller.Middleware.Auth(), controller.GetDetailComplaint)
		api.POST("/complaint/:id", controller.Middleware.Auth(), controller.SubmitComplaint)
	}
}

func (controller *ComplaintController) GetData(c *gin.Context) {
	defer common.Catch(c)
	status := c.Query("status")
	authorizedUser := c.MustGet("user").(*common.JWTClaims)
	data, err := controller.ComplaintService.GetData(authorizedUser.Unique.String(), status)
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

func (controller *ComplaintController) GetDetailComplaint(c *gin.Context) {
	defer common.Catch(c)
	id := c.Param("id")
	authorizedUser := c.MustGet("user").(*common.JWTClaims)
	data, err := controller.ComplaintService.GetDetailComplaint(authorizedUser.Unique.String(), id)
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

func (controller *ComplaintController) SubmitComplaint(c *gin.Context) {
	defer common.Catch(c)
	id := c.Param("id")
	authorizedUser := c.MustGet("user").(*common.JWTClaims)
	var request request.SubmitComplaintRequest
	c.BindJSON(&request)
	data, err := controller.ComplaintService.SubmitComplaint(authorizedUser.Unique.String(), id, request)
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
