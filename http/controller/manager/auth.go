package manager

import (
	"net/http"

	"github.com/bagusyanuar/go-deltomed/common"
	"github.com/bagusyanuar/go-deltomed/exception"
	"github.com/bagusyanuar/go-deltomed/http/request"
	usecaseManager "github.com/bagusyanuar/go-deltomed/usecase/manager"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	AuthService usecaseManager.AuthService
}

func NewAuthController(authService usecaseManager.AuthService) AuthController {
	return AuthController{AuthService: authService}
}

func (controller *AuthController) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/manager")
	{
		api.POST("/sign-in", controller.SignIn)
	}
}

func (controller *AuthController) SignIn(c *gin.Context) {
	var request request.CreateSignInRequest
	c.BindJSON(&request)
	accessToken, err := controller.AuthService.SignIn(request)
	if err != nil {
		switch err {
		case exception.ErrorPasswordNotMatch:
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.APIResponse{
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
				Data:    nil,
			})
			return
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, common.APIResponse{
				Code:    http.StatusNotFound,
				Message: err.Error(),
				Data:    nil,
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, common.APIResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
				Data:    nil,
			})
			return
		}
	}
	c.JSON(http.StatusOK, common.APIResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    accessToken,
	})
}