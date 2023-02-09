package main

import (
	"fmt"
	"ohlc-data-api/api/config"
	"ohlc-data-api/api/server"

	"os"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

// @title           OHLC Data API
// @version         1.0
// @description     This is a data API for processing and uploading OHLC data.
// @termsOfService  http://swagger.io/terms/

// @contact.name   CoinHako Support
// @contact.url    https://help.coinhako.com/
// @contact.email  support@coinhako.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /api/v1
// @schemes https http
func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
}

var logger *zap.Logger

func run() error {

	switch gin.Mode() {
	case gin.ReleaseMode:
		logger = config.Logger()
	default:
		err := godotenv.Load()
		if err != nil {
			fmt.Println("error loading .env file")
		}

		logger = config.Logger()
	}

	g := gin.Default()
	g.Use(config.CORSMiddleware())
	g.Use(requestid.New())

	d := config.BuildProject()
	svr := server.NewServer(g, d)

	svr.MapRoutes()
	if err := svr.SetupDB(); err != nil {
		logger.Error("Databases failed to start" + err.Error())
		return err
	}

	defer logger.Sync()
	return svr.Start()
}
