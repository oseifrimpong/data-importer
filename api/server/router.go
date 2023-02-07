package server

import (
	"ohlc-data-api/api/config"
	"ohlc-data-api/api/delivery"
	"ohlc-data-api/api/service"
	"os"

	"github.com/dimiro1/health"
	"github.com/dimiro1/health/db"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// swagger embed files

	_ "ohlc-data-api/docs"
)

func (ds *dserver) MapRoutes() {
	versionGroup := ds.router.Group("api/v1")
	unVer := ds.router.Group("api")

	ds.dataRoutes(versionGroup)
	ds.root(unVer)

}

func (ds *dserver) dataRoutes(api *gin.RouterGroup) {
	u := api.Group("/")
	{
		// var logger *zap.Logger
		// ds.cont.Invoke(func(l *zap.Logger) {
		// 	logger = l
		// })
		var dataSvc service.DataService
		ds.cont.Invoke(func(svc service.DataService) {
			dataSvc = svc
		})

		svc := delivery.NewDataController(dataSvc)
		// u.GET("/data", svc.Get)
		u.POST("/data", svc.Create)
	}
}

// HealthCheck
// @Summary Show the status of server.
// @Description Get the status of server and version
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} dto.APIResponse
// @Failure 400,404 {object} dto.APIResponse
// @Failure 500 {object} dto.APIResponse
// @Failure default {object} dto.APIResponse
// @Router api/healthcheck [get]
func (ds *dserver) root(api *gin.RouterGroup) {
	h := api.Group("/")
	{
		pgDb, _ := config.InitializeDB()

		psql, _ := pgDb.DB()
		postgres := db.NewPostgreSQLChecker(psql)

		handler := health.NewHandler()
		handler.AddChecker("database", postgres)
		handler.AddInfo("api", "Service is alive")
		handler.AddInfo("version", os.Getenv("VERSION"))

		h.GET("healthcheck", gin.WrapH(handler))
		h.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	}
}
