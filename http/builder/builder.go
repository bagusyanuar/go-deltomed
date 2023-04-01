package builder

import (
	adminController "github.com/bagusyanuar/go-deltomed/http/controller/admin"
	engineerController "github.com/bagusyanuar/go-deltomed/http/controller/engineer"
	managerController "github.com/bagusyanuar/go-deltomed/http/controller/manager"
	supportController "github.com/bagusyanuar/go-deltomed/http/controller/support"
	usecaseAdminRepositories "github.com/bagusyanuar/go-deltomed/repositories/admin"
	usecaseEngineerRepositories "github.com/bagusyanuar/go-deltomed/repositories/engineer"
	usecaseManagerRepositories "github.com/bagusyanuar/go-deltomed/repositories/manager"
	usecaseSupportRepositories "github.com/bagusyanuar/go-deltomed/repositories/support"
	usecaseAdminService "github.com/bagusyanuar/go-deltomed/service/admin"
	usecaseEngineerService "github.com/bagusyanuar/go-deltomed/service/engineer"
	usecaseManagerService "github.com/bagusyanuar/go-deltomed/service/manager"
	usecaseSupportService "github.com/bagusyanuar/go-deltomed/service/support"
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

	//engineer group
	engineerComplaintRepository := usecaseEngineerRepositories.NewComplaintRepository(db)
	engineerComplaintService := usecaseEngineerService.NewComplaintService(engineerComplaintRepository)
	engineerComplaintController := engineerController.NewComplaintController(engineerComplaintService)
	engineerComplaintController.RegisterRoute(route)

	//support group
	supportAuthRepository := usecaseSupportRepositories.NewAuthRepository(db)
	supportAuthService := usecaseSupportService.NewAuthService(supportAuthRepository)
	supportAuthController := supportController.NewAuthController(supportAuthService)
	supportAuthController.RegisterRoute(route)

	supportDivisionRepository := usecaseSupportRepositories.NewDivisionRepository(db)
	supportDivisionService := usecaseSupportService.NewDivisionService(supportDivisionRepository)
	supportDivisionController := supportController.NewDivisionController(supportDivisionService)
	supportDivisionController.RegisterRoute(route)

	supportComplaintRepository := usecaseSupportRepositories.NewComplaintRepository(db)
	supportComplaintService := usecaseSupportService.NewComplaintService(supportComplaintRepository)
	supportComplaintController := supportController.NewComplaintController(supportComplaintService)
	supportComplaintController.RegisterRoute(route)

	//manager group
	managerProfileRepository := usecaseManagerRepositories.NewProfileRepository(db)

	managerComplaintRepository := usecaseManagerRepositories.NewComplaintRepository(db)
	managerComplaintService := usecaseManagerService.NewComplaintService(managerComplaintRepository, managerProfileRepository)
	managerComplaintController := managerController.NewComplaintController(managerComplaintService)
	managerComplaintController.RegisterRoute(route)
}
