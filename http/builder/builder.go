package builder

import (
	adminController "github.com/bagusyanuar/go-deltomed/http/controller/admin"
	usecaseAdminRepositories "github.com/bagusyanuar/go-deltomed/repositories/admin"
	usecaseAdminService "github.com/bagusyanuar/go-deltomed/service/admin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BuildGroup(route *gin.Engine, db *gorm.DB) {
	//admin group
	adminDivisionRepository := usecaseAdminRepositories.NewDivisionRepository(db)
	adminDivisionService := usecaseAdminService.NewDivisionService(adminDivisionRepository)
	adminDivisionController := adminController.NewDivisionController(adminDivisionService)
	adminDivisionController.RegisterRoute(route)

	adminUserRepository := usecaseAdminRepositories.NewUserRepository(db)
	adminUserService := usecaseAdminService.NewUserService(adminUserRepository)
	adminUserController := adminController.NewUserController(adminUserService)
	adminUserController.RegisterRoute(route)

	adminLocationRepository := usecaseAdminRepositories.NewLocationRepository(db)
	adminLocationService := usecaseAdminService.NewLocationService(adminLocationRepository)
	adminLocationController := adminController.NewLocationController(adminLocationService)
	adminLocationController.RegisterRoute(route)

}
