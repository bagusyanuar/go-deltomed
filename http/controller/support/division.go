package support

import (
	"net/http"

	"github.com/bagusyanuar/go-deltomed/common"
	"github.com/bagusyanuar/go-deltomed/http/middleware"
	usecaseSupport "github.com/bagusyanuar/go-deltomed/usecase/support"
	"github.com/gin-gonic/gin"
)

type DivisionController struct {
	DivisionService usecaseSupport.DivisionService
	Middleware      middleware.Middleware
}

func NewDivisionController(divisionService usecaseSupport.DivisionService) DivisionController {
	return DivisionController{DivisionService: divisionService, Middleware: middleware.Middleware{}}
}

func (controller *DivisionController) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/support")
	{
		api.GET("/division", controller.Middleware.Auth(), controller.GetData)
	}
}

func (controller *DivisionController) GetData(c *gin.Context) {
	q := c.Query("q")
	data, err := controller.DivisionService.GetData(q)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}
	c.JSON(http.StatusOK, common.APIResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}
