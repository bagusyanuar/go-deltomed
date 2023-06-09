package admin

import (
	"net/http"
	"strconv"

	"github.com/bagusyanuar/go-deltomed/common"
	"github.com/bagusyanuar/go-deltomed/http/middleware"
	"github.com/bagusyanuar/go-deltomed/http/request"
	usecaseAdmin "github.com/bagusyanuar/go-deltomed/usecase/admin"
	"github.com/gin-gonic/gin"
)

type DivisionController struct {
	DivisionService usecaseAdmin.DivisionService
	Middleware       middleware.Middleware
}

func NewDivisionController(divisionService usecaseAdmin.DivisionService) DivisionController {
	return DivisionController{DivisionService: divisionService, Middleware: middleware.Middleware{}}
}

func (controller *DivisionController) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/admin")
	{
		api.GET("/division", controller.Middleware.Auth(), controller.FindAll)
		api.POST("/division", controller.Middleware.Auth(),controller.Create)
		api.GET("/division/:id", controller.Middleware.Auth(),controller.FindByID)
		api.PATCH("/division/:id", controller.Middleware.Auth(),controller.Patch)
		api.DELETE("/division/:id/delete", controller.Middleware.Auth(),controller.Delete)
	}
}

func (controller *DivisionController) FindAll(c *gin.Context) {
	q := c.Query("q")
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	data, err := controller.DivisionService.FindAll(q, limit, offset)
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

func (controller *DivisionController) FindByID(c *gin.Context) {
	id := c.Param("id")
	data, err := controller.DivisionService.FindByID(id)
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

func (controller *DivisionController) Create(c *gin.Context) {
	var request request.CreateDivisionRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.APIResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	data, err := controller.DivisionService.Create(request)
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

func (controller *DivisionController) Patch(c *gin.Context) {
	id := c.Param("id")
	var request request.CreateDivisionRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.APIResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	data, err := controller.DivisionService.Patch(id, request)
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

func (controller *DivisionController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := controller.DivisionService.Delete(id)
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
