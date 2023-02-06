package server

import (
	"ohlc-data-api/api/config"
	"ohlc-data-api/api/models"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type dserver struct {
	router *gin.Engine
	cont   *dig.Container
}

func NewServer(e *gin.Engine, c *dig.Container) *dserver {
	return &dserver{
		router: e,
		cont:   c,
	}
}

func (ds *dserver) Start() error {
	ds.router.SetTrustedProxies(nil)
	return ds.router.Run(":" + os.Getenv("APP_PORT"))
}

func (ds *dserver) SetupDB() error {

	db, err := config.InitializeDB()
	if err != nil {
		return err
	}

	models := []interface{}{
		&models.Data{},
	}

	err = db.AutoMigrate(models...)
	if err != nil {
		panic(err)
	}

	return nil
}
